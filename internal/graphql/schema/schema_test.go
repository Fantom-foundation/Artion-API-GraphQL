// Package schema implements GraphQL schema definition and parser.
package schema

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/graphql/resolvers"
	"artion-api-graphql/internal/logger"
	"github.com/onsi/gomega"
	"testing"
)

// schemaContentTestCase represents a single test case for GraphQL API schema bundle content.
type schemaContentTestCase struct {
	re  string
	msg string
}

// TestSchemaContent tests if the authorized private GraphQL schema contains all the expected elements
// we need to support the expected API interface.
func TestSchemaContent(t *testing.T) {
	// what test cases are we going to check?
	cases := []schemaContentTestCase{
		{
			re:  "(?m)^schema\\s+{(\\s+(query|mutation|subscription)\\s*:\\s*(Query|Mutation|Subscription))+\\s*}",
			msg: "root/master schema definition must exist",
		},
		{
			re:  "(?m)^\\s*type\\s+Query\\s+{$",
			msg: "query entry point type definition must exist",
		},
		{
			re:  "(?m)^\\s*type\\s+Mutation\\s+{$",
			msg: "mutation entry point type definition must exist",
		},
		{
			re:  "(?m)^\\s*scalar\\s+BigInt$",
			msg: "shared scalars definition must exist",
		},
	}

	// modules definition to exist based on modules
	// we check only one type for each module to confirm the module does exist
	modules := map[string]string{
		"Collection": "Collection related types definition must exist",
		"Token": "Token related types definition must exist",
	}

	// create gΩ instance
	g := gomega.NewGomegaWithT(t)

	// use list of cases to check the base schema
	for _, tc := range cases {
		g.Expect(schema).To(gomega.MatchRegexp(tc.re), tc.msg)
	}

	// check modules by types definition
	// we expect the schema to contain: type <type name> {
	for mt, msg := range modules {
		g.Expect(schema).To(gomega.MatchRegexp("(?m)^type\\s+%s\\s+{$", mt), msg)
	}
}

func TestParseSchema(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	cfg, err := config.Load()
	g.Expect(err).To(gomega.BeNil())
	resolvers.SetConfig(cfg)
	resolvers.SetLogger(logger.New(cfg))
	_, err = Schema()
	g.Expect(err).To(gomega.BeNil())
}
