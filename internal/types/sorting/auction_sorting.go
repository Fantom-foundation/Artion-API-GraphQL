package sorting

import "artion-api-graphql/internal/types"

type AuctionSorting int8

const (
	AuctionSortingNone AuctionSorting = iota
)

func (ts AuctionSorting) SortedFieldBson() string {
	return ""
}

func (ts AuctionSorting) OrdinalFieldBson() string {
	return "_id"
}

func (ts AuctionSorting) GetCursor(offer *types.Auction) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["_id"] = offer.ID()
	return CursorFromParams(params)
}
