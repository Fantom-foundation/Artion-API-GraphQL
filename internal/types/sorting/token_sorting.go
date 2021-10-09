package sorting

import "artion-api-graphql/internal/types"

// TokenSorting is Sorting implementation for Token collections.
type TokenSorting int8
const (
	TokenSortingNone TokenSorting = iota
	TokenSortingCreated
	TokenSortingLastList
	TokenSortingLastTrade
)

func (ts TokenSorting) SortedFieldBson() string {
	switch ts {
	case TokenSortingNone: return ""
	case TokenSortingCreated: return "created"
	case TokenSortingLastList: return "last_list"
	case TokenSortingLastTrade: return "last_trade"
	}
	return ""
}

func (ts TokenSorting) OrdinalFieldBson() string {
	return "index"
}

func (ts TokenSorting) GetCursor(token *types.Token) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["index"] = token.OrdinalIndex
	if ts == TokenSortingCreated {
		params["created"] = token.Created
	}
	if ts == TokenSortingLastList {
		params["last_list"] = token.LastList
	}
	if ts == TokenSortingLastTrade {
		params["last_trade"] = token.LastTrade
	}
	return CursorFromParams(params)
}
