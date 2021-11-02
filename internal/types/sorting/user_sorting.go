package sorting

import "artion-api-graphql/internal/types"

type UserSorting int8

const (
	UserSortingNone UserSorting = iota
)

func (ts UserSorting) SortedFieldBson() string {
	return ""
}

func (ts UserSorting) OrdinalFieldBson() string {
	return "_id"
}

func (ts UserSorting) GetCursor(user *types.User) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["_id"] = user.Id
	return CursorFromParams(params)
}
