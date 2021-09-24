// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import "artion-api-graphql/internal/types"

// PageInfo represents general resolvable information about the current page of a list of elements.
type PageInfo struct {
	StartCursor     *types.Cursor
	EndCursor       *types.Cursor
	HasNextPage     bool
	HasPreviousPage bool
}
