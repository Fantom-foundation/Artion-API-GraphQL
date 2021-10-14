package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

func (p *Proxy) GetShippingAddress(user common.Address) (*types.ShippingAddress, error) {
	return p.shared.GetShippingAddress(user)
}

func (p *Proxy) UpsertShippingAddress(address *types.ShippingAddress) error {
	return p.shared.UpsertShippingAddress(address)
}
