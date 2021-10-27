package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// GetShippingAddress provides a shipping address stored for the given user address.
func (p *Proxy) GetShippingAddress(user common.Address) (*types.ShippingAddress, error) {
	return p.shared.GetShippingAddress(user)
}

// StoreShippingAddress stores the given shipping address into the persistent storage.
func (p *Proxy) StoreShippingAddress(address *types.ShippingAddress) error {
	return p.shared.UpsertShippingAddress(address)
}
