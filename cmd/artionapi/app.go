// Package main implements the API server entry point.
package main

import (
	"artion-api-graphql/cmd/artionapi/build"
	"artion-api-graphql/internal/auth"
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/graphql/resolvers"
	"artion-api-graphql/internal/handlers"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/svc"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// apiServer implements the API server application
type apiServer struct {
	cfg          *config.Config
	log          logger.Logger
	srv          *http.Server
	isVersionReq bool
}

// init initializes the API server
func (app *apiServer) init() {
	// make sure to capture version request and rescan depth
	flag.BoolVar(&app.isVersionReq, "v", false, "get the application version")

	// get the configuration including parsing the calling flags
	var err error
	app.cfg, err = config.Load()
	if nil != err {
		log.Fatal(err)
		return
	}

	// configure logger based on the configuration
	app.log = logger.New(app.cfg)

	// make sure to pass logger and config to internals
	handlers.SetConfig(app.cfg)
	repository.SetConfig(app.cfg)
	repository.SetLogger(app.log)
	resolvers.SetConfig(app.cfg)
	resolvers.SetLogger(app.log)
	svc.SetConfig(app.cfg)
	svc.SetLogger(app.log)
	auth.SetConfig(app.cfg)

	// make the HTTP server
	app.makeHttpServer()
}

// run executes the API server function.
func (app *apiServer) run() {
	// print the app version and exit if this is the only thing requested
	build.PrintVersion(app.cfg)
	if app.isVersionReq {
		return
	}

	// start the services
	svc.Mgr()

	// make sure to capture terminate signals
	app.observeSignals()

	// start responding to requests
	app.log.Infof("welcome to Artion GraphQL server")
	app.log.Infof("listening for requests on %s", app.cfg.Server.BindAddress)

	// listen the interface
	err := app.srv.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			app.log.Notice("HTTP server closed")
		} else {
			app.log.Error(err)
		}
	}

	// terminate the app
	app.terminate()
}

// makeHttpServer creates and configures the HTTP server to be used to serve incoming requests
func (app *apiServer) makeHttpServer() {
	// create request MUXer
	srvMux := new(http.ServeMux)

	// create HTTP server to handle our requests
	app.srv = &http.Server{
		Addr:              app.cfg.Server.BindAddress,
		ReadTimeout:       time.Second * time.Duration(app.cfg.Server.ReadTimeout),
		WriteTimeout:      time.Second * time.Duration(app.cfg.Server.WriteTimeout),
		IdleTimeout:       time.Second * time.Duration(app.cfg.Server.IdleTimeout),
		ReadHeaderTimeout: time.Second * time.Duration(app.cfg.Server.HeaderTimeout),
		Handler:           handlers.Cors(app.cfg, app.log, srvMux),
	}

	// setup handlers
	app.setupHandlers(srvMux)
}

// setupHandlers initializes an array of handlers for our HTTP API end-points.
func (app *apiServer) setupHandlers(mux *http.ServeMux) {
	// setup GraphQL API handler
	h := //http.TimeoutHandler(
		handlers.AuthHandler(handlers.Api(app.log), app.log)
	//time.Second*time.Duration(app.cfg.Server.ResolverTimeout), "Service timeout.",
	//)
	mux.Handle("/api", h)
	mux.Handle("/graphql", h)

	// handle GraphiQL interface
	mux.Handle("/graphi", handlers.GraphiHandler(app.log))

	// handle images
	mux.Handle("/images/token/", handlers.ImageHandler(app.log, handlers.TokenImageResolver))
	mux.Handle("/images/avatar/", handlers.ImageHandler(app.log, handlers.UserAvatarResolver))
	mux.Handle("/images/collection/", handlers.ImageHandler(app.log, handlers.CollectionImageResolver))

	mux.Handle("/metadata/token/", handlers.TokenMetadataHandler(app.log))

	// handle image upload
	mux.Handle("/upload-image/user-avatar", handlers.AuthHandler(handlers.UploadImageHandler(app.log, handlers.StoreUserAvatar), app.log))
	mux.Handle("/upload-image/user-banner", handlers.AuthHandler(handlers.UploadImageHandler(app.log, handlers.StoreUserBanner), app.log))
	mux.Handle("/upload-image/token", handlers.AuthHandler(handlers.UploadImageHandler(app.log, handlers.StoreToken), app.log))
	mux.Handle("/upload-image/collection", handlers.AuthHandler(handlers.UploadImageHandler(app.log, handlers.StoreCollection), app.log))
}

// observeSignals setups terminate signals observation.
func (app *apiServer) observeSignals() {
	// log what we do
	app.log.Info("os signals captured")

	// make the signal consumer
	ts := make(chan os.Signal, 1)
	signal.Notify(ts, syscall.SIGINT, syscall.SIGTERM)

	// start monitoring
	go func() {
		<-ts
		if err := app.srv.Close(); err != nil {
			app.log.Errorf("could not terminate HTTP listener")
			os.Exit(0)
		}
	}()
}

// terminate modules of the API server.
func (app *apiServer) terminate() {
	// close resolvers
	app.log.Notice("closing resolver")
	resolvers.Resolver().Close()

	// terminate services
	app.log.Notice("closing services")
	svc.Close()

	// terminate connections to DB, blockchain, etc.
	app.log.Notice("closing repository")
	repository.Close()
}
