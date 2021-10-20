// Package handlers holds HTTP/WS handlers chain along with separate middleware implementations.
package handlers

import (
	"artion-api-graphql/internal/graphql/schema"
	"artion-api-graphql/internal/logger"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
	"net/http"
)

// Api constructs and return the API HTTP handlers chain for serving GraphQL API calls.
func Api(log logger.Logger) http.Handler {
	// create new parsed GraphQL schema
	sch, err := schema.Schema()
	if err != nil {
		log.Panicf("can not get the GraphQL schema; %s", err)
		return nil
	}

	// return the constructed API handler chain
	return graphqlws.NewHandlerFunc(sch, &relay.Handler{Schema: sch})
}
