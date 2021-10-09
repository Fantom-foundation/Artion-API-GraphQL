package sorting

import "artion-api-graphql/internal/types"

type TokenEventSorting int8

const (
	TokenEventSortingNone TokenEventSorting = iota
)

func (ts TokenEventSorting) SortedFieldBson() string {
	return ""
}

func (ts TokenEventSorting) OrdinalFieldBson() string {
	return "_id"
}

func (ts TokenEventSorting) GetCursor(tokenEvent *types.TokenEvent) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["_id"] = tokenEvent.Id
	return CursorFromParams(params)
}
