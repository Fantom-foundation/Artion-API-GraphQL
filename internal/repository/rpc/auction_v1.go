package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"
	"artion-api-graphql/internal/types"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// auctionDefaultDurationShift defines the default duration of an auction, if not defined otherwise.
const auctionDefaultDurationShift = (24 * 365 * 10) * time.Hour

// AuctionContractV1 is IAuctionContract implementation for contracts in 1.0.0 version
type AuctionContractV1 struct {
	auctionV1aContract *contracts.FantomAuction   // newer version V1a
	auctionV1Contract  *contracts.FantomAuctionV1 // older version V1
}

func (ac *AuctionContractV1) ExtendAuctionDetails(au *types.Auction, _ *big.Int) error {
	// try to decode V1a first
	err0 := ac.extendAuctionDetailsV1a(au)
	if err0 != nil {
		// try to decode V1 if V1a failed
		err1 := ac.extendAuctionDetailsV1(au)
		if err1 != nil {
			// log error only if both versions fails
			log.Errorf("unable to extend auction %s/%s details; V0: %s; V1: %s",
				au.Contract.String(), au.TokenId.String(), au.Owner.String(), err0, err1)
		}
	}
	return nil
}

func (ac *AuctionContractV1) extendAuctionDetailsV1a(au *types.Auction) error {
	// get auction details
	res, err := ac.auctionV1aContract.Auctions(&bind.CallOpts{
		BlockNumber: nil,
		Context:     context.Background(),
	}, au.Contract, (*big.Int)(&au.TokenId))
	if err != nil {
		return err
	}

	// copy the owner
	au.Owner = res.Owner

	// make sure we have what we came for
	if nil == res.ReservePrice || nil == res.StartTime || nil == res.EndTime {
		log.Warningf("missing mandatory field on auction %s/%s", au.Contract.String(), au.TokenId.String())
		return nil
	}

	// transfer values
	au.ReservePrice = (hexutil.Big)(*res.ReservePrice)
	au.MinimalBid = (hexutil.Big)(*res.MinBid)
	au.StartTime = types.Time(time.Unix(res.StartTime.Int64(), 0))
	au.EndTime = types.Time(time.Unix(res.EndTime.Int64(), 0))

	return nil
}

func (ac *AuctionContractV1) extendAuctionDetailsV1(au *types.Auction) error {
	// get auction details
	res, err := ac.auctionV1Contract.Auctions(&bind.CallOpts{
		BlockNumber: nil,
		Context:     context.Background(),
	}, au.Contract, (*big.Int)(&au.TokenId))
	if err != nil {
		return err
	}

	// copy the owner
	au.Owner = res.Owner

	// make sure we have what we came for
	if nil == res.ReservePrice || nil == res.StartTime || nil == res.EndTime {
		log.Warningf("missing mandatory field on auction %s/%s", au.Contract.String(), au.TokenId.String())
		return nil
	}

	// transfer values
	au.ReservePrice = (hexutil.Big)(*res.ReservePrice)
	au.MinimalBid = (hexutil.Big)(*res.ReservePrice)

	// do we have a start time? use creation time, if not
	if 0 < res.StartTime.Int64() {
		au.StartTime = types.Time(time.Unix(res.StartTime.Int64(), 0))
	} else {
		au.StartTime = au.Created
	}

	// do we have an end time? use creation + constant if not
	if 0 < res.EndTime.Int64() {
		au.EndTime = types.Time(time.Unix(res.EndTime.Int64(), 0))
	} else {
		log.Infof("Impossible auction EndTime %d for %s/%s", res.EndTime.Int64(), au.Contract.String(), au.TokenId.String())
		au.EndTime = types.Time(time.Time(au.StartTime).Add(auctionDefaultDurationShift))
	}
	return nil
}

// GetMinBid extract minimal required bid a used must place to participate on auction.
func (ac *AuctionContractV1) GetMinBid(contract *common.Address, tokenId *big.Int) (*big.Int, error) {
	// get the highest bid
	highest, err := ac.auctionV1aContract.HighestBids(nil, *contract, tokenId)
	if err != nil {
		return nil, err
	}

	// for zero highest bid, we use min. bid instead
	if 0 == new(big.Int).Cmp(highest.Bid) {
		auction, err := ac.auctionV1aContract.Auctions(nil, *contract, tokenId)
		if err != nil {
			return nil, err
		}

		// use MinBid instead (could be zero, or reserve price, based on auction config)
		highest.Bid = auction.MinBid
	}

	minIncrement, err := ac.auctionV1aContract.MinBidIncrement(nil)
	if err != nil {
		return nil, err
	}
	if 0 == minIncrement.Cmp(new(big.Int)) {
		minIncrement = new(big.Int).SetInt64(1)
	}
	return new(big.Int).Add(highest.Bid, minIncrement), nil
}
