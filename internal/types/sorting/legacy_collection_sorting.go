package sorting

import "artion-api-graphql/internal/types"

type LegacyCollectionSorting int8

const (
	LegacyCollectionSortingNone LegacyCollectionSorting = iota
	LegacyCollectionSortingName
)

func (ts LegacyCollectionSorting) SortedFieldBson() string {
	switch ts {
	case LegacyCollectionSortingName: return "collectionName"
	}
	return ""
}

func (ts LegacyCollectionSorting) OrdinalFieldBson() string {
	return "_id"
}

func (ts LegacyCollectionSorting) GetCursor(LegacyCollection *types.LegacyCollection) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["_id"] = LegacyCollection.Address
	if ts == LegacyCollectionSortingName {
		params["collectionName"] = LegacyCollection.Name
	}
	return CursorFromParams(params)
}
