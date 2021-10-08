package sorting

import "artion-api-graphql/internal/types"

type CollectionSorting int8

const (
	CollectionSortingNone CollectionSorting = iota
	CollectionSortingName
)

func (ts CollectionSorting) SortedFieldBson() string {
	switch ts {
	case CollectionSortingName: return "name"
	}
	return ""
}

func (ts CollectionSorting) OrdinalFieldBson() string {
	return "_id"
}

func (ts CollectionSorting) GetCursor(collection *types.Collection) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["_id"] = collection.Address
	if ts == CollectionSortingName {
		params["name"] = collection.Name
	}
	return types.CursorFromParams(params)
}
