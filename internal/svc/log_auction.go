// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

// auctionCreated processes an event for newly created auction on an ERC-721 token.
// Auction::AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address payToken)
func auctionCreated(evt *eth.Log, lo *logObserver) {
	// sanity check: 1 + 2 topics; 1 x uint256 = 32 bytes
	if len(evt.Data) != 32 || len(evt.Topics) != 3 {
		log.Errorf("not Auction::AuctionCreated() event #%d/#%d; expected 32 bytes of data, %d given; expected 3 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}

	// make the listing
	auction := types.Auction{
		Contract:      common.BytesToAddress(evt.Topics[1].Bytes()),
		TokenId:       hexutil.Big(*new(big.Int).SetBytes(evt.Topics[2].Bytes())),
		Owner:         common.Address{},
		Quantity:      hexutil.Big(*new(big.Int).SetInt64(1)),
		PayToken:      common.BytesToAddress(evt.Data[:]),
		MinimalBid:    hexutil.Big{},
		ReservePrice:  hexutil.Big{},
		Created:       types.Time(time.Unix(int64(blk.Time), 0)),
		StartTime:     types.Time{},
		EndTime:       types.Time{},
		Closed:        nil,
		LastBid:       nil,
		LastBidPlaced: nil,
		LastBidder:    nil,
		Winner:        nil,
		Resolved:      nil,
		OrdinalIndex:  types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
	}

	// extend the auction with details pulled from the contract
	if err := repo.ExtendAuctionDetailAt(&auction, new(big.Int).SetUint64(evt.BlockNumber)); err != nil {
		log.Errorf("failed to load extended auction details; %s", err.Error())
	}

	// clear previous bids for the token
	if err := repo.ClearAuctionBids(&auction.Contract, (*big.Int)(&auction.TokenId)); err != nil {
		log.Errorf("could not clear auction bids; %s", err.Error())
	}

	// store the listing into database
	if err := repo.StoreAuction(&auction); err != nil {
		log.Errorf("could not store auction; %s", err.Error())
	}

	// mark the token as being auctioned
	if err := repo.TokenMarkAuctioned(
		&auction.Contract,
		(*big.Int)(&auction.TokenId),
		repo.GetUnitPriceAt(lo.marketplace, &auction.PayToken, new(big.Int).SetUint64(evt.BlockNumber), (*big.Int)(&auction.ReservePrice)),
		(*time.Time)(&auction.Created),
	); err != nil {
		log.Errorf("could not mark token as having auction; %s", err.Error())
	}

	log.Infof("added new auction of %s/%s started by %s", auction.Contract.String(), auction.TokenId.String(), auction.Owner.String())
}

// auctionStartTimeUpdated processes auction start time update event log.
// Auction::UpdateAuctionStartTime(address indexed nftAddress, uint256 indexed tokenId, uint256 startTime)
func auctionStartTimeUpdated(evt *eth.Log, lo *logObserver) {
	auctionTimeBoundaryUpdated(evt, lo, func(au *types.Auction, tx types.Time) {
		au.StartTime = tx
		log.Infof("auction %s/%s start time updated to %s", au.Contract.String(), au.TokenId.String(), time.Time(tx).Format(time.RFC1123))
	})
}

// auctionEndTimeUpdated processes auction end time update event log.
// Auction::UpdateAuctionEndTime(address indexed nftAddress, uint256 indexed tokenId, uint256 endTime)
func auctionEndTimeUpdated(evt *eth.Log, lo *logObserver) {
	auctionTimeBoundaryUpdated(evt, lo, func(au *types.Auction, tx types.Time) {
		au.EndTime = tx
		log.Infof("auction %s/%s end time updated to %s", au.Contract.String(), au.TokenId.String(), time.Time(tx).Format(time.RFC1123))
	})
}

// auctionTimeBoundaryUpdated processes given time boundary change event log.
func auctionTimeBoundaryUpdated(evt *eth.Log, lo *logObserver, update func(*types.Auction, types.Time)) {
	// sanity check: 1 + 2 topics; 1 x uint256 = 32 bytes
	if len(evt.Data) != 32 || len(evt.Topics) != 3 {
		log.Errorf("not Auction::UpdateAuction-X-Time() event #%d/#%d; expected 32 bytes of data, %d given; expected 3 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	contract := common.BytesToAddress(evt.Topics[1].Bytes())
	tokenID := new(big.Int).SetBytes(evt.Topics[2].Bytes())

	// pull the auction involved
	auction, err := repo.GetAuction(&contract, tokenID)
	if err != nil {
		log.Errorf("updated auction %s/%s not found; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err.Error())
		return
	}

	// make the change using provided callback
	update(auction, types.Time(time.Unix(new(big.Int).SetBytes(evt.Data[:]).Int64(), 0)))

	// store the listing into database
	if err := repo.StoreAuction(auction); err != nil {
		log.Errorf("could not store auction; %s", err.Error())
	}

	// mark the token as being re-auctioned
	if err := repo.TokenMarkAuctioned(
		&auction.Contract,
		(*big.Int)(&auction.TokenId),
		repo.GetUnitPriceAt(lo.marketplace, &auction.PayToken, new(big.Int).SetUint64(evt.BlockNumber), (*big.Int)(&auction.ReservePrice)),
		(*time.Time)(&auction.Created),
	); err != nil {
		log.Errorf("could not mark token as having auction; %s", err.Error())
	}
}

// auctionReserveUpdated processes auction reserve price updated.
// Auction::UpdateAuctionReservePrice(address indexed nftAddress, uint256 indexed tokenId, address payToken, uint256 reservePrice)
func auctionReserveUpdated(evt *eth.Log, lo *logObserver) {
	// sanity check: 1 + 2 topics; 1 x uint256 + 1 x address = 64 bytes
	if len(evt.Data) != 64 || len(evt.Topics) != 3 {
		log.Errorf("not Auction::UpdateAuctionReservePrice() event #%d/#%d; expected 64 bytes of data, %d given; expected 3 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	contract := common.BytesToAddress(evt.Topics[1].Bytes())
	tokenID := new(big.Int).SetBytes(evt.Topics[2].Bytes())

	// pull the auction involved
	auction, err := repo.GetAuction(&contract, tokenID)
	if err != nil {
		log.Errorf("updated auction %s/%s not found; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err.Error())
		return
	}

	auction.PayToken = common.BytesToAddress(evt.Data[:32])
	auction.ReservePrice = hexutil.Big(*new(big.Int).SetBytes(evt.Data[32:]))

	// store the listing into database
	if err := repo.StoreAuction(auction); err != nil {
		log.Errorf("could not store auction; %s", err.Error())
	}

	// mark the token as being re-auctioned
	if err := repo.TokenMarkAuctioned(
		&auction.Contract,
		(*big.Int)(&auction.TokenId),
		repo.GetUnitPriceAt(lo.marketplace, &auction.PayToken, new(big.Int).SetUint64(evt.BlockNumber), (*big.Int)(&auction.ReservePrice)),
		(*time.Time)(&auction.Created),
	); err != nil {
		log.Errorf("could not mark token as having auction; %s", err.Error())
	}
}

// auctionCanceled processes auction being canceled event log.
// Auction::AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId)
func auctionCanceled(evt *eth.Log, _ *logObserver) {
	// sanity check: 1 + 2 topics; 0 bytes data
	if len(evt.Data) != 0 || len(evt.Topics) != 3 {
		log.Errorf("not Auction::AuctionCancelled() event #%d/#%d; expected no data, %d bytes given; expected 3 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}

	contract := common.BytesToAddress(evt.Topics[1].Bytes())
	tokenID := new(big.Int).SetBytes(evt.Topics[2].Bytes())

	// pull the auction involved
	auction, err := repo.GetAuction(&contract, tokenID)
	if err != nil {
		log.Errorf("canceled auction %s/%s not found; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err.Error())
		return
	}

	ts := types.Time(time.Unix(int64(blk.Time), 0))
	auction.Closed = &ts

	// store the listing into database
	if err := repo.StoreAuction(auction); err != nil {
		log.Errorf("could not store auction; %s", err.Error())
	}

	// mark the token as being re-auctioned
	if err := repo.TokenMarkUnAuctioned(&auction.Contract, (*big.Int)(&auction.TokenId)); err != nil {
		log.Errorf("could not mark token as not having auction; %s", err.Error())
	}
}

// auctionResolved processes the auction resolved event log.
// Auction::AuctionResulted(address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 unitPrice, uint256 winningBid)
func auctionResolved(evt *eth.Log, lo *logObserver) {
	// sanity check: 1 + 3 topics; 2 x uint256 + 1 address = 96 bytes data
	if len(evt.Data) != 96 || len(evt.Topics) != 4 {
		log.Errorf("not Auction::AuctionResulted() event #%d/#%d; expected 96 bytes of data, %d given; expected 4 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}

	contract := common.BytesToAddress(evt.Topics[1].Bytes())
	tokenID := new(big.Int).SetBytes(evt.Topics[2].Bytes())
	winner := common.BytesToAddress(evt.Topics[3].Bytes())
	winAmount := new(big.Int).SetBytes(evt.Data[32:64])
	payToken := common.BytesToAddress(evt.Data[:32])

	// pull the auction involved
	auction, err := repo.GetAuction(&contract, tokenID)
	if err != nil {
		log.Errorf("resolved auction %s/%s not found; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err.Error())
		return
	}

	ts := types.Time(time.Unix(int64(blk.Time), 0))
	auction.Resolved = &ts
	auction.Closed = &ts
	auction.Winner = &winner
	auction.WinningBid = (*hexutil.Big)(winAmount)

	// store the listing into database
	if err := repo.StoreAuction(auction); err != nil {
		log.Errorf("could not store auction; %s", err.Error())
	}

	// mark the token as sold
	if err := repo.TokenMarkSold(
		&auction.Contract,
		(*big.Int)(&auction.TokenId),
		repo.GetUnitPriceAt(lo.marketplace, &payToken, new(big.Int).SetUint64(evt.BlockNumber), new(big.Int).SetBytes(evt.Data[64:])),
		(*time.Time)(&ts),
	); err != nil {
		log.Errorf("could not mark token as sold; %s", err.Error())
	}
}
