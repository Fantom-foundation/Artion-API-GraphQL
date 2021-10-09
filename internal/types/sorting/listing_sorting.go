package sorting

import "artion-api-graphql/internal/types"

type ListingSorting int8

const (
	ListingSortingNone ListingSorting = iota
)

func (ts ListingSorting) SortedFieldBson() string {
	return ""
}

func (ts ListingSorting) OrdinalFieldBson() string {
	return "_id"
}

func (ts ListingSorting) GetCursor(listing *types.Listing) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["_id"] = listing.Id
	return CursorFromParams(params)
}
