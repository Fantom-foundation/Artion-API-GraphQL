package sorting

import "artion-api-graphql/internal/types"

type FollowSorting int8

const (
	FollowSortingNone FollowSorting = iota
)

func (ts FollowSorting) SortedFieldBson() string {
	return ""
}

func (ts FollowSorting) OrdinalFieldBson() string {
	return "_id"
}

func (ts FollowSorting) GetCursor(follow *types.Follow) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["_id"] = follow.Id
	return CursorFromParams(params)
}
