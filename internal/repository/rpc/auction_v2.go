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

// AuctionContractV2 is IAuctionContract implementation for contracts in 2.0.0 version
type AuctionContractV2 struct {
	auctionV2Contract *contracts.FantomAuctionV2
}

func (ac *AuctionContractV2) ExtendAuctionDetails(au *types.Auction, _ *big.Int) error {
	err2 := ac.extendAuctionDetailsV2(au)
	if err2 != nil {
		// log error only if both versions fails
		log.Errorf("unable to extend auction %s/%s details; V2: %s",
			au.Contract.String(), au.TokenId.String(), au.Owner.String(), err2)
	}
	return nil
}

func (ac *AuctionContractV2) extendAuctionDetailsV2(au *types.Auction) error {
	// get auction details
	res, err := ac.auctionV2Contract.Auctions(&bind.CallOpts{
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

// GetMinBid extract minimal required bid a used must place to participate on auction.
func (ac *AuctionContractV2) GetMinBid(contract *common.Address, tokenId *big.Int) (*big.Int, error) {
	// get the highest bid
	highest, err := ac.auctionV2Contract.HighestBids(nil, *contract, tokenId)
	if err != nil {
		return nil, err
	}

	// for zero highest bid, we use min. bid instead
	if 0 == new(big.Int).Cmp(highest.Bid) {
		auction, err := ac.auctionV2Contract.Auctions(nil, *contract, tokenId)
		if err != nil {
			return nil, err
		}

		// use MinBid instead (could be zero, or reserve price, based on auction config)
		highest.Bid = auction.MinBid
	}

	minIncrement, err := ac.auctionV2Contract.MinBidIncrement(nil)
	if err != nil {
		return nil, err
	}
	if 0 == minIncrement.Cmp(new(big.Int)) {
		minIncrement = new(big.Int).SetInt64(1)
	}
	return new(big.Int).Add(highest.Bid, minIncrement), nil
}
