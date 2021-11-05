package sorting

import "artion-api-graphql/internal/types"

// TokenSorting is Sorting implementation for Token collections.
type TokenSorting int8
const (
	TokenSortingNone TokenSorting = iota
	TokenSortingCreated         // Recently Created / Oldest
	TokenSortingLastListTime    // Recently Listed
	TokenSortingLastTradeTime   // Recently Sold
	TokenSortingAuctionUntil    // Ending Soon
	TokenSortingPrice           // Most Expensive / Cheapest
	TokenSortingLastTradeAmount // Highest Last Sale
	// Mostly Viewed
)

func (ts TokenSorting) SortedFieldBson() string {
	switch ts {
	case TokenSortingNone: return ""
	case TokenSortingCreated: return "created"
	case TokenSortingLastListTime: return "last_list"
	case TokenSortingLastTradeTime: return "last_trade"
	case TokenSortingAuctionUntil: return "auction_until"
	case TokenSortingPrice: return "price"
	case TokenSortingLastTradeAmount: return "amo_trade"
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
	if ts == TokenSortingLastListTime {
		params["last_list"] = token.LastListing
	}
	if ts == TokenSortingLastTradeTime {
		params["last_trade"] = token.LastTrade
	}
	if ts == TokenSortingAuctionUntil {
		params["auction_until"] = token.HasAuctionUntil
	}
	if ts == TokenSortingPrice {
		params["price"] = token.Price
	}
	if ts == TokenSortingLastTradeAmount {
		params["amo_trade"] = token.AmountLastTrade
	}
	return CursorFromParams(params)
}
