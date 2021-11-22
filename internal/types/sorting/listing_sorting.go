package sorting

import "artion-api-graphql/internal/types"

type ListingSorting int8

const (
	ListingSortingNone ListingSorting = iota
	ListingSortingCreated
)

func (ts ListingSorting) SortedFieldBson() string {
	switch ts {
	case ListingSortingNone: return ""
	case ListingSortingCreated: return "created"
	}
	return ""
}

func (ts ListingSorting) OrdinalFieldBson() string {
	return "_id"
}

func (ts ListingSorting) GetCursor(listing *types.Listing) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["_id"] = listing.ID()
	if ts == ListingSortingCreated {
		params["created"] = listing.Created
	}
	return CursorFromParams(params)
}
