package sorting

import "artion-api-graphql/internal/types"

type OwnershipSorting int8

const (
	OwnershipSortingNone OwnershipSorting = iota
)

func (ts OwnershipSorting) SortedFieldBson() string {
	return ""
}

func (ts OwnershipSorting) OrdinalFieldBson() string {
	return "_id"
}

func (ts OwnershipSorting) GetCursor(ownership *types.Ownership) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["_id"] = ownership.ID()
	return CursorFromParams(params)
}
