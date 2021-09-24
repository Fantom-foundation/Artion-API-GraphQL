package resolvers

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type TokenEvent types.TokenEvent

type TokenEventEdge struct {
	Node *TokenEvent
}

func (edge TokenEventEdge) Cursor() (Cursor, error) {
	return CursorFromObjectId(edge.Node.Id), nil
}

type TokenEventConnection struct {
	Edges      []TokenEventEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
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
