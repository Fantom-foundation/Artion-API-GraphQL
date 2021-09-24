package resolvers

import (
	"artion-api-graphql/internal/types"
	"errors"
)

// PaginationInput represents input type for GraphQL cursor pagination
// specified by https://relay.dev/graphql/connections.htm
type PaginationInput struct {
	First   *int32
	After   *types.Cursor
	Last    *int32
	Before  *types.Cursor
}

func (input *PaginationInput) ToCursorCount() (cursor types.Cursor, count int, backward bool, err error) {
	if input == nil || (input.First == nil && input.Last == nil) {
		cursor = ""
		count = 10
		backward = false
		return
	}

	if input.First != nil && *input.First > 0 && input.Last == nil && input.Before == nil {
		if input.After != nil {
			cursor = *input.After
		}
		count = int(*input.First)
		backward = false
		return
	}

	if input.Last != nil && *input.Last > 0 && input.First == nil && input.After == nil {
		if input.Before != nil {
			cursor = *input.Before
		}
		count = int(*input.Last)
		backward = true
		return
	}

	err = errors.New("invalid pagination parameters")
	return
}
