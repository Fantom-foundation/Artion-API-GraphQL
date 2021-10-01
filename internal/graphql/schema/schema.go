// Package schema implements GraphQL schema definition and parser.
package schema

import (
	"artion-api-graphql/internal/graphql/resolvers"
	_ "embed"
	"github.com/graph-gophers/graphql-go"
)

//go:embed gen/schema.graphql
var schema string

// Schema parses API schema for GraphQL handlers chain.
func Schema() (*graphql.Schema, error) {
	// we don't want to write a method for each type field if it could be matched directly
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}

	// parse the schema
	schema, err := graphql.ParseSchema(schema, resolvers.Resolver(), opts...)
	if err != nil {
		return nil, err
	}

	return schema, nil
}
