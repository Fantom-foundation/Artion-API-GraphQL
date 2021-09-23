package resolvers

import "errors"

// input type for GraphQL cursor pagination
// specified by https://relay.dev/graphql/connections.htm
type PaginationInput struct {
	First   *int32
	After   *Cursor
	Last    *int32
	Before  *Cursor
}

func (input *PaginationInput) ToCursorCount() (cursor string, count int, err error) {
	if input == nil || (input.First == nil && input.Last == nil) {
		cursor = ""
		count = 10
		return
	}

	if input.First != nil && *input.First > 0 && input.Last == nil && input.Before == nil {
		if input.After != nil {
			cursor = string(*input.After)
		}
		count = int(*input.First)
		return
	}

	if input.Last != nil && *input.Last > 0 && input.First == nil && input.After == nil {
		if input.Before != nil {
			cursor = string(*input.Before)
		}
		count = -int(*input.Last)
		return
	}

	err = errors.New("invalid pagination parameters")
	return
}
