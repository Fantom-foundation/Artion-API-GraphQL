// Package handlers holds HTTP/WS handlers chain along with separate middleware implementations.
package handlers

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"github.com/rs/cors"
	"net/http"
)

// Cors constructs and return the wrapping HTTP handler adding CORS header.
func Cors(cfg *config.Config, log logger.Logger, handler http.Handler) http.Handler {
	// create new CORS handler and attach the GraphQL resolver
	corsHandler := cors.New(corsOptions(cfg))
	corsHandler.Log = log

	// return the constructed API handler chain
	return corsHandler.Handler(handler)
}

// corsOptions constructs new set of options for the CORS handler based on provided configuration.
func corsOptions(cfg *config.Config) cors.Options {
	return cors.Options{
		AllowedOrigins: cfg.Server.CorsOrigin,
		AllowedMethods: []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
		ExposedHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:         3000,
	}
}
