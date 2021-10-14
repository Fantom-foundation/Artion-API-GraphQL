// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// ExtendAuctionDetailAt adds contract stored details to the provided auction record.
func (o *Opera) ExtendAuctionDetailAt(au *types.Auction, block *big.Int) error {
	// get auction details
	res, err := o.auctionContract.Auctions(&bind.CallOpts{
		BlockNumber: block,
		Context:     context.Background(),
	}, au.Contract, (*big.Int)(&au.TokenId))
	if err != nil {
		log.Errorf("auction %s/%s not available; %s",
			au.Contract.String(), au.TokenId.String(), au.Owner.String(), err.Error())
		return err
	}

	// make sure we have what we came for
	if nil == res.ReservePrice || nil == res.StartTime || nil == res.EndTime {
		return fmt.Errorf("missing mandatory field on auction %s/%s", au.Contract.String(), au.TokenId.String())
	}

	// transfer values
	au.Owner = res.Owner
	au.ReservePrice = (hexutil.Big)(*res.ReservePrice)
	au.StartTime = types.Time(time.Unix(res.StartTime.Int64(), 0))
	au.EndTime = types.Time(time.Unix(res.EndTime.Int64(), 0))

	return nil
}

// AuctionHighestBidAmount collects the value of the current highest bid from the auction.
func (o *Opera) AuctionHighestBidAmount(contract *common.Address, tokenID *big.Int) (*big.Int, error) {
	res, err := o.auctionContract.HighestBids(nil, *contract, tokenID)
	if err != nil {
		log.Errorf("can not get the highest bid of %s/%s; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err.Error())
		return nil, err
	}
	return res.Bid, nil
}

// AuctionMinimalBidAmount collects the value of the current minimal bid from the auction.
func (o *Opera) AuctionMinimalBidAmount(contract *common.Address, tokenID *big.Int) *big.Int {
	res, err := o.auctionContract.Auctions(nil, *contract, tokenID)
	if err != nil {
		log.Errorf("can not get auction %s/%s; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err.Error())
		return new(big.Int)
	}
	return res.MinBid
}

// AuctionMinBidIncrement collects the amount of minimal bid increment for auctions.
func (o *Opera) AuctionMinBidIncrement() *big.Int {
	val, err := o.auctionContract.MinBidIncrement(nil)
	if err != nil {
		log.Errorf("failed to get min bid increment; %s", err.Error())
		return new(big.Int)
	}

	if 0 == val.Cmp(new(big.Int)) {
		return new(big.Int).Add(val, new(big.Int).SetInt64(1))
	}
	return val
}
