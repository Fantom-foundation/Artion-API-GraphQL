// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

// PageInfo represents general resolvable information about the current page of a list of elements.
type PageInfo struct {
	StartCursor     *Cursor
	EndCursor       *Cursor
	HasNextPage     bool
	HasPreviousPage bool
}
