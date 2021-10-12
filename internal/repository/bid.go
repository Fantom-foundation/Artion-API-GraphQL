// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// StoreAuctionBid stores given auction bid into the database.
func (p *Proxy) StoreAuctionBid(bid *types.AuctionBid) error {
	return p.db.StoreAuctionBid(bid)
}

// DeleteAuctionBid removes a bid of specific bidder from specific auction.
func (p *Proxy) DeleteAuctionBid(contract *common.Address, tokenID *big.Int, bidder *common.Address) error {
	return p.db.DeleteAuctionBid(contract, tokenID, bidder)
}

// ClearAuctionBids removes all the bids stored for the given auction.
func (p *Proxy) ClearAuctionBids(contract *common.Address, tokenID *big.Int) error {
	return p.db.ClearAuctionBids(contract, tokenID)
}
