// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

// SetAuctionBidder sets a new bidder (or no bidder) into the specified auction.
func (p *Proxy) SetAuctionBidder(contract *common.Address, tokenID *big.Int, bidder *common.Address, placed *types.Time) error {
	return p.db.SetAuctionBidder(contract, tokenID, bidder, placed)
}

func (p *Proxy) ListAuctions(contract *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (out *types.AuctionList, err error) {
	return p.db.ListAuctions(contract, tokenId, owner, cursor, count, backward)
}
