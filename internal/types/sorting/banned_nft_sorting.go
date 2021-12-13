package sorting

import "artion-api-graphql/internal/types"

type BannedNftSorting int8

const (
	BannedNftSortingNone BannedNftSorting = iota
	BannedNftSortingUpdated
)

func (ts BannedNftSorting) SortedFieldBson() string {
	switch ts {
	case BannedNftSortingNone: return ""
	case BannedNftSortingUpdated: return "updatedAt"
	}
	return ""
}

func (ts BannedNftSorting) OrdinalFieldBson() string {
	return "_id"
}

func (ts BannedNftSorting) GetCursor(nft *types.BannedNft) (types.Cursor, error) {
	params := make(map[string]interface{})
	params["_id"] = nft.Id
	if ts == BannedNftSortingUpdated {
		params["updatedAt"] = nft.Updated
	}
	return CursorFromParams(params)
}
