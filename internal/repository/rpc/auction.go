// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// IAuctionContract defines single interface for all auction contract versions.
type IAuctionContract interface {
	ExtendAuctionDetails(au *types.Auction, block *big.Int) error
	GetMinBid(contract *common.Address, tokenId *big.Int) (*big.Int, error)
}

// ExtendAuctionDetailAt adds contract stored details to the provided auction record.
func (o *Opera) ExtendAuctionDetailAt(au *types.Auction, block *big.Int) error {
	ac, exists := o.auctionContracts[au.AuctionHall]
	if !exists {
		return fmt.Errorf("unable to get detail for auction %s/%s - unknown auction contract %s",
			au.Contract.String(), au.TokenId.String(), au.AuctionHall.String())
	}
	return ac.ExtendAuctionDetails(au, block)
}

// AuctionGetMinBid provides a minimal bid amount required to participate in the auction.
func (o *Opera) AuctionGetMinBid(au *types.Auction) (*big.Int, error) {
	ac, exists := o.auctionContracts[au.AuctionHall]
	if !exists {
		return nil, fmt.Errorf("unable to get min bid for auction %s/%s - unknown auction contract %s",
			au.Contract.String(), au.TokenId.String(), au.AuctionHall.String())
	}
	minBid, err := ac.GetMinBid(&au.Contract, (*big.Int)(&au.TokenId))
	if err != nil {
		log.Errorf("unable to get auction %s/%s min bid; %s",
			au.Contract.String(), au.TokenId.String(), err.Error())
	}
	return minBid, err
}

// GetAuctionProps provides properties of auctions running on given auction contract
func (o *Opera) GetAuctionProps(auctionHall common.Address) (*types.AuctionProps, error) {
	props, exists := o.auctionContractsProps[auctionHall]
	if exists {
		return &props, nil
	}
	return nil, fmt.Errorf("trying to obtain props of unknown auction contract %s", auctionHall.String())
}
