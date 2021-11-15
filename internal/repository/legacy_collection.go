package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

func (p *Proxy) GetLegacyCollection(address common.Address) (*types.LegacyCollection, error) {
	return p.shared.GetLegacyCollection(address)
}

func (p *Proxy) InsertLegacyCollection(collection *types.LegacyCollection) error {
	return p.shared.InsertLegacyCollection(collection)
}

func (p *Proxy) ListLegacyCollections(search *string, mintableBy *common.Address, cursor types.Cursor, count int, backward bool) (out *types.LegacyCollectionList, err error) {
	return p.shared.ListLegacyCollections(search, mintableBy, cursor, count, backward)
}
