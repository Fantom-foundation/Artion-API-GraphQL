package handlers

import (
	"artion-api-graphql/internal/logger"
	_ "embed"
	"html/template"
	"net/http"
)

// graphiqlTemplate represents the template for the GraphiQL HTML output.
//go:embed graphi.html
var graphiqlTemplate string

// GraphiHandler builds a HTTP handler function for GraphiQL playground.
func GraphiHandler(log logger.Logger) http.Handler {
	// build the handler function
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse the template, we don't expect it to fail
		t, err := template.New("graphiql").Parse(graphiqlTemplate)
		if err != nil {
			// log and send 500 response to client
			log.Critical("can not parse GraphiQL template; %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// execute the template
		err = t.Execute(w, r.Host)
		if err != nil {
			// log and send 500 response to client
			log.Critical("can not server GraphiQL playground; %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
