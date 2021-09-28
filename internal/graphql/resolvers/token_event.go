package resolvers

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type TokenEvent types.TokenEvent

type TokenEventEdge struct {
	Node *TokenEvent
}

func (edge TokenEventEdge) Cursor() (types.Cursor, error) {
	return types.CursorFromId(edge.Node.Id), nil
}

type TokenEventConnection struct {
	Edges      []TokenEventEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

func NewTokenEventConnection(list *types.TokenEventList) (con *TokenEventConnection, err error) {
	con = new(TokenEventConnection)
	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]TokenEventEdge, len(list.Collection))
	for i := 0; i < len(list.Collection); i++ {
		con.Edges[i].Node = (*TokenEvent)(list.Collection[i])
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

func (event TokenEvent) Type() (string, error) {
	return TokenEventTypeToString(event.EventType), nil
}

func TokenEventTypeToString(t types.TokenEventType) string {
	switch t {
	case types.EvtTpItemListed:
		return "ITEM_LISTED"
	case types.EvtTpItemUpdated:
		return "ITEM_UPDATED"
	case types.EvtTpItemCanceled:
		return "ITEM_CANCELED"
	case types.EvtTpItemSold:
		return "ITEM_SOLD"
	case types.EvtTpOfferCreated:
		return "OFFER_CREATED"
	case types.EvtTpOfferCanceled:
		return "OFFER_CANCELED"
	}
	return "UNKNOWN"
}

func TokenEventTypeFromString(t string) types.TokenEventType {
	switch t {
	case "ITEM_LISTED":
		return types.EvtTpItemListed
	case "ITEM_UPDATED":
		return types.EvtTpItemUpdated
	case "ITEM_CANCELED":
		return types.EvtTpItemCanceled
	case "ITEM_SOLD":
		return types.EvtTpItemSold
	case "OFFER_CREATED":
		return types.EvtTpOfferCreated
	case "OFFER_CANCELED":
		return types.EvtTpOfferCanceled
	}
	return types.EvtTpUnknown
}
