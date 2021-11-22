package sorting

import "artion-api-graphql/internal/types"

type OfferSorting int8

const (
	OfferSortingNone OfferSorting = iota
	OfferSortingCreated
)

func (ts OfferSorting) SortedFieldBson() string {
	switch ts {
	case OfferSortingNone: return ""
	case OfferSortingCreated: return "created"
	}
	return ""
}

func (ts OfferSorting) OrdinalFieldBson() string {
	return "_id"
}

func (ts OfferSorting) GetCursor(offer *types.Offer) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["_id"] = offer.ID()
	if ts == OfferSortingCreated {
		params["created"] = offer.Created
	}
	return CursorFromParams(params)
}
