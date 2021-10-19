package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

func (p *Proxy) GetLegacyCollection(address common.Address) (*types.LegacyCollection, error) {
	return p.shared.GetLegacyCollection(address)
}
