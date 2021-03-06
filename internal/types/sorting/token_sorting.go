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
	TokenSortingListPrice       // Most Expensive / Cheapest (by listings only)
	TokenSortingOfferPrice      // Most Expensive / Cheapest (by offers only)
	TokenSortingLastTradeAmount // Highest Last Sale
	TokenSortingCachedViews     // Mostly Viewed
	TokenSortingCachedLikes     // Mostly Liked
)

func (ts TokenSorting) SortedFieldBson() string {
	switch ts {
	case TokenSortingNone: return ""
	case TokenSortingCreated: return "created"
	case TokenSortingLastListTime: return "last_list"
	case TokenSortingLastTradeTime: return "last_trade"
	case TokenSortingAuctionUntil: return "auction_until"
	case TokenSortingPrice: return "price"
	case TokenSortingListPrice: return "min_list.usd"
	case TokenSortingOfferPrice: return "max_offer.usd"
	case TokenSortingLastTradeAmount: return "amo_trade.usd"
	case TokenSortingCachedViews: return "views"
	case TokenSortingCachedLikes: return "likes"
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
		params["price"] = token.AmountPrice
	}
	if ts == TokenSortingListPrice {
		params["min_list.usd"] = token.MinListPrice.Usd
	}
	if ts == TokenSortingOfferPrice {
		params["max_offer.usd"] = token.MaxOfferPrice.Usd
	}
	if ts == TokenSortingLastTradeAmount {
		params["amo_trade.usd"] = token.AmountLastTrade.Usd
	}
	if ts == TokenSortingCachedViews {
		params["views"] = token.CachedViews
	}
	if ts == TokenSortingCachedLikes {
		params["likes"] = token.CachedLikes
	}
	return CursorFromParams(params)
}
