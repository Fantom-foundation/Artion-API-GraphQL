package resolvers

import (
	"artion-api-graphql/internal/types/sorting"
)

// tokenSortingFromString converts string from GraphQL to types.TokenSorting
func tokenSortingFromString(s *string) (sorting.TokenSorting, error) {
	if s == nil {
		return sorting.TokenSortingNone, nil
	}
	switch *s {
	case "CREATED":
		return sorting.TokenSortingCreated, nil
	case "LAST_LISTING":
		return sorting.TokenSortingLastListTime, nil
	case "LAST_TRADE":
		return sorting.TokenSortingLastTradeTime, nil
	case "AUCTION_UNTIL":
		return sorting.TokenSortingAuctionUntil, nil
	case "PRICE":
		return sorting.TokenSortingPrice, nil
	case "LAST_TRADE_AMOUNT":
		return sorting.TokenSortingLastTradeAmount, nil
	case "VIEWS":
		return sorting.TokenSortingCachedViews, nil
	case "LIKES":
		return sorting.TokenSortingCachedLikes, nil
	}
	panic("Unknown TokenSorting")
}

// isSortingDirectionDesc converts GraphQL SortingDirection string to isDescending bool
func isSortingDirectionDesc(direction *string) bool {
	if direction != nil && *direction == "DESC" {
		return true
	}
	return false
}
