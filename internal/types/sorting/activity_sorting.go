package sorting

import "artion-api-graphql/internal/types"

type ActivitySorting int8

const (
	ActivitySortingNone ActivitySorting = iota
)

func (ts ActivitySorting) SortedFieldBson() string {
	return ""
}

func (ts ActivitySorting) OrdinalFieldBson() string {
	return "index"
}

func (ts ActivitySorting) GetCursor(activity *types.Activity) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["index"] = activity.OrdinalIndex
	return CursorFromParams(params)
}
