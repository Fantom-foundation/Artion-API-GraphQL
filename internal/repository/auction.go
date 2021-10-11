// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// GetAuction provides the auction stored in the database, if available.
func (p *Proxy) GetAuction(contract *common.Address, tokenID *big.Int) (*types.Auction, error) {
	return p.db.GetAuction(contract, tokenID)
}

// StoreAuction adds the provided auction into the database.
func (p *Proxy) StoreAuction(au *types.Auction) error {
	return p.db.StoreAuction(au)
}

// ExtendAuctionDetailAt adds contract stored details to the provided auction record.
func (p *Proxy) ExtendAuctionDetailAt(au *types.Auction, block *big.Int) error {
	return p.rpc.ExtendAuctionDetailAt(au, block)
}
