package resolvers

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type Activity types.Activity

type ActivityEdge struct {
	Node *Activity
}

func (edge ActivityEdge) Cursor() (types.Cursor, error) {
	return sorting.ActivitySortingNone.GetCursor((*types.Activity)(edge.Node))
}

type ActivityConnection struct {
	Edges      []ActivityEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

func NewActivityConnection(list *types.ActivityList) (con *ActivityConnection, err error) {
	con = new(ActivityConnection)
	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]ActivityEdge, len(list.Collection))
	for i := 0; i < len(list.Collection); i++ {
		con.Edges[i].Node = (*Activity)(list.Collection[i])
	}
	con.PageInfo.HasNextPage = list.HasNext
	con.PageInfo.HasPreviousPage = list.HasPrev
	if len(list.Collection) > 0 {
		startCur, err := con.Edges[0].Cursor()
		if err != nil {
			return nil, err
		}
		endCur, err := con.Edges[len(con.Edges)-1].Cursor()
		if err != nil {
			return nil, err
		}
		con.PageInfo.StartCursor = &startCur
		con.PageInfo.EndCursor = &endCur
	}
	return con, err
}

func (activity Activity) Type() (string, error) {
	return ActivityTypeToString(activity.ActType), nil
}

func (activity Activity) Token() (*Token, error) {
	return NewToken(&activity.Contract, &activity.TokenId)
}

func ActivityTypeToString(t types.ActivityType) string {
	switch t {
	case types.EvtListingCreated:
		return "LISTING_CREATED"
	case types.EvtListingUpdated:
		return "LISTING_UPDATED"
	case types.EvtListingCancelled:
		return "LISTING_CANCELLED"
	case types.EvtListingSold:
		return "LISTING_SOLD"
	case types.EvtOfferCreated:
		return "OFFER_CREATED"
	case types.EvtOfferCancelled:
		return "OFFER_CANCELLED"
	case types.EvtOfferSold:
		return "OFFER_SOLD"
	case types.EvtAuctionCreated:
		return "AUCTION_CREATED"
	case types.EvtAuctionBid:
		return "AUCTION_BID"
	case types.EvtAuctionBidWithdrawn:
		return "AUCTION_BID_WITHDRAW"
	case types.EvtAuctionCancelled:
		return "AUCTION_CANCELLED"
	case types.EvtAuctionResolved:
		return "AUCTION_RESOLVED"
	case types.EvtAuctionUpdated:
		return "AUCTION_UPDATED"
	}
	return "UNKNOWN"
}
