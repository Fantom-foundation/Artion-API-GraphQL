// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// StoreOwnership stores the given NFT ownership record in persistent storage.
func (p *Proxy) StoreOwnership(to *types.Ownership) error {
	return p.db.StoreOwnership(to)
}

// StoreBurn stores the given NFT burn record in persistent storage.
func (p *Proxy) StoreBurn(bu *types.NFTBurn) error {
	return p.db.StoreBurn(bu)
}

// ListOwnerships lists token ownerships records.
func (p *Proxy) ListOwnerships(contract *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (out *types.OwnershipList, err error) {
	return p.db.ListOwnerships(contract, tokenId, owner, cursor, count, backward)
}
