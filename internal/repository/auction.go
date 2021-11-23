// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// GetAuction provides the auction stored in the database, if available.
func (p *Proxy) GetAuction(contract *common.Address, tokenID *big.Int, auctionHall *common.Address) (*types.Auction, error) {
	return p.db.GetAuction(contract, tokenID, auctionHall)
}

func (p *Proxy) GetLastAuction(contract *common.Address, tokenID *big.Int) (*types.Auction, error) {
	return p.db.GetLastAuction(contract, tokenID)
}

// StoreAuction adds the provided auction into the database.
func (p *Proxy) StoreAuction(au *types.Auction) error {
	return p.db.StoreAuction(au)
}

// SetAuctionActive sets IsActive state of auction when token ownership changes.
func (p *Proxy) SetAuctionActive(contract *common.Address, tokenID *big.Int, owner *common.Address, isActive bool) error {
	return p.db.SetAuctionActive(contract, tokenID, owner, isActive)
}

// ExtendAuctionDetailAt adds contract stored details to the provided auction record.
func (p *Proxy) ExtendAuctionDetailAt(au *types.Auction, block *big.Int) error {
	return p.rpc.ExtendAuctionDetailAt(au, block)
}

// AuctionGetMinBid provides a minimal bid amount required to participate in the auction.
func (p *Proxy) AuctionGetMinBid(contract *common.Address, tokenID *big.Int) (*big.Int, error) {
	// get the highest bid
	hb, err := p.rpc.AuctionHighestBidAmount(contract, tokenID)
	if err != nil {
		return nil, err
	}

	// for zero highest bid, we use min. bid instead
	if 0 == new(big.Int).Cmp(hb) {
		return p.rpc.AuctionMinimalBidAmount(contract, tokenID), nil
	}
	return new(big.Int).Add(hb, p.rpc.AuctionMinBidIncrement()), nil
}
