// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/core/types"
	"github.com/status-im/keycard-go/hexutils"
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

	// make the auction
	auction := types.Auction{
		Owner:         common.Address{},
		Contract:      common.BytesToAddress(evt.Topics[1].Bytes()),
		TokenId:       hexutil.Big(*new(big.Int).SetBytes(evt.Topics[2].Bytes())),
		AuctionHall:   evt.Address,
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
		IsActive:      true,
	}

	// extend the auction with details pulled from the contract
	if err := repo.ExtendAuctionDetailAt(&auction, new(big.Int).SetUint64(evt.BlockNumber)); err != nil {
		log.Errorf("failed to load extended auction details; %s", err.Error())
	}

	// if auction owner is not known, find them by the transaction sender
	if 0 == bytes.Compare(auction.Owner.Bytes(), zeroAddress.Bytes()) || 0 == bytes.Compare(auction.PayToken.Bytes(), zeroAddress.Bytes()) {
		log.Warningf("parsing trx at #%d / #%d: %s", evt.BlockNumber, evt.TxIndex, evt.TxHash.String())
		extendAuctionFromTransaction(&auction, evt.TxHash)
	}

	// zero pay token means native FTM (an option on older auctions)
	if 0 == bytes.Compare(auction.PayToken.Bytes(), zeroAddress.Bytes()) {
		copy(auction.PayToken[:], cfg.Contracts.WrappedFTM[:])
		log.Infof("using wFTM %s instead of native pay token", auction.PayToken.String())
	}

	// clear previous bids for the token
	if err := repo.ClearAuctionBids(&auction.Contract, (*big.Int)(&auction.TokenId)); err != nil {
		log.Errorf("could not clear auction bids; %s", err.Error())
	}

	// store the listing into database
	if err := repo.StoreAuction(&auction); err != nil {
		log.Errorf("could not store auction; %s", err.Error())
	}

	price := repo.GetUnifiedPriceAt(&auction.PayToken, new(big.Int).SetUint64(evt.BlockNumber), (*big.Int)(&auction.ReservePrice))

	// mark the token as being auctioned
	if err := repo.TokenMarkAuctioned(
		&auction.Contract,
		(*big.Int)(&auction.TokenId),
		price,
		(*time.Time)(&auction.Created),
	); err != nil {
		log.Errorf("could not mark token as having auction; %s", err.Error())
	}

	// log activity
	activity := types.Activity{
		Transaction:  evt.TxHash,
		OrdinalIndex: types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
		Time:         auction.Created,
		ActType:      types.EvtAuctionCreated,
		Contract:     auction.Contract,
		TokenId:      auction.TokenId,
		AuctionHall:  &evt.Address,
		Quantity:     &auction.Quantity,
		From:         auction.Owner,
		PayToken:     &auction.PayToken,
		UnitPrice:    &auction.ReservePrice,
		UnifiedPrice: price.Usd,
		StartTime:    &auction.StartTime,
		EndTime:      &auction.EndTime,
	}
	if err := repo.StoreActivity(&activity); err != nil {
		log.Errorf("could not store auction activity; %s", err.Error())
		return
	}

	log.Infof("added new auction of %s/%s started by %s at %s", auction.Contract.String(), auction.TokenId.String(), auction.Owner.String(), evt.TxHash)
}

// extendAuctionFromTransaction tries to use transaction data to populate missing auction details.
func extendAuctionFromTransaction(auction *types.Auction, tx common.Hash) {
	// collect transaction data
	sender, _, data := repo.MustTransactionData(tx)

	// we expect 4 bytes for func call + 6 params of 32 bytes = 196 bytes
	if len(data) != 196 {
		log.Criticalf("invalid call at %s; expected 196 bytes, loaded %d bytes", tx.String(), len(data))
		return
	}

	// is this the right function call?
	// createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, uint256 _endTimestamp)
	if 0 != bytes.Compare(data[:4], hexutils.HexToBytes("14ec4106")) {
		log.Criticalf("invalid function call at %s", tx.String())
		return
	}

	auction.Owner = sender
	auction.PayToken = common.BytesToAddress(data[68:100])
	auction.ReservePrice = (hexutil.Big)(*new(big.Int).SetBytes(data[100:132]))
	auction.StartTime = types.Time(time.Unix(new(big.Int).SetBytes(data[132:164]).Int64(), 0))
	auction.EndTime = types.Time(time.Unix(new(big.Int).SetBytes(data[164:]).Int64(), 0))
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
	auction, err := repo.GetAuction(&contract, tokenID, &evt.Address)
	if auction == nil {
		log.Errorf("updated auction %s/%s not found; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err)
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
		repo.GetUnifiedPriceAt(&auction.PayToken, new(big.Int).SetUint64(evt.BlockNumber), (*big.Int)(&auction.ReservePrice)),
		(*time.Time)(&auction.Created),
	); err != nil {
		log.Errorf("could not mark token as having auction; %s", err.Error())
	}

	// log activity
	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}
	activity := types.Activity{
		Transaction:  evt.TxHash,
		OrdinalIndex: types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
		Time:         types.Time(time.Unix(int64(blk.Time), 0)),
		ActType:      types.EvtAuctionUpdated,
		Contract:     auction.Contract,
		TokenId:      auction.TokenId,
		AuctionHall:  &evt.Address,
		Quantity:     &auction.Quantity,
		From:         auction.Owner,
		StartTime:    &auction.StartTime,
		EndTime:      &auction.EndTime,
	}
	if err := repo.StoreActivity(&activity); err != nil {
		log.Errorf("could not store auction activity; %s", err.Error())
		return
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
	auction, err := repo.GetAuction(&contract, tokenID, &evt.Address)
	if auction == nil {
		log.Errorf("updated auction %s/%s not found; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err)
		return
	}

	auction.PayToken = common.BytesToAddress(evt.Data[:32])
	auction.ReservePrice = hexutil.Big(*new(big.Int).SetBytes(evt.Data[32:]))

	// store the listing into database
	if err := repo.StoreAuction(auction); err != nil {
		log.Errorf("could not store auction; %s", err.Error())
	}

	price := repo.GetUnifiedPriceAt(&auction.PayToken, new(big.Int).SetUint64(evt.BlockNumber), (*big.Int)(&auction.ReservePrice))

	// mark the token as being re-auctioned
	if err := repo.TokenMarkAuctioned(
		&auction.Contract,
		(*big.Int)(&auction.TokenId),
		price,
		(*time.Time)(&auction.Created),
	); err != nil {
		log.Errorf("could not mark token as having auction; %s", err.Error())
	}

	// log activity
	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}
	activity := types.Activity{
		Transaction:  evt.TxHash,
		OrdinalIndex: types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
		Time:         types.Time(time.Unix(int64(blk.Time), 0)),
		ActType:      types.EvtAuctionUpdated,
		Contract:     auction.Contract,
		TokenId:      auction.TokenId,
		AuctionHall:  &evt.Address,
		Quantity:     &auction.Quantity,
		From:         auction.Owner,
		UnitPrice:    &auction.ReservePrice,
		UnifiedPrice: price.Usd,
		PayToken:     &auction.PayToken,
	}
	if err := repo.StoreActivity(&activity); err != nil {
		log.Errorf("could not store auction activity; %s", err.Error())
		return
	}

	// notify subscribers
	event := types.Event{Type: "AUCTION_RESERVE_UPDATED", Auction: auction}
	subscriptionManager := GetSubscriptionsManager()
	subscriptionManager.PublishAuctionEvent(event)
	subscriptionManager.PublishUserEvent(auction.Owner, event)
	if auction.LastBidder != nil {
		subscriptionManager.PublishUserEvent(*auction.LastBidder, event)
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
	auction, err := repo.GetAuction(&contract, tokenID, &evt.Address)
	if auction == nil {
		log.Errorf("canceled auction %s/%s not found; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err)
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

	// log activity
	activity := types.Activity{
		Transaction:  evt.TxHash,
		OrdinalIndex: types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
		Time:         ts,
		ActType:      types.EvtAuctionCancelled,
		Contract:     auction.Contract,
		TokenId:      auction.TokenId,
		AuctionHall:  &evt.Address,
		Quantity:     &auction.Quantity,
		From:         auction.Owner,
	}
	if err := repo.StoreActivity(&activity); err != nil {
		log.Errorf("could not store auction activity; %s", err.Error())
		return
	}

	// notify subscribers
	event := types.Event{Type: "AUCTION_CANCELLED", Auction: auction}
	subscriptionManager := GetSubscriptionsManager()
	subscriptionManager.PublishAuctionEvent(event)
	subscriptionManager.PublishUserEvent(auction.Owner, event)
	if auction.LastBidder != nil {
		subscriptionManager.PublishUserEvent(*auction.LastBidder, event)
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

	contract := common.BytesToAddress(evt.Topics[1].Bytes())
	tokenID := new(big.Int).SetBytes(evt.Topics[2].Bytes())
	winner := common.BytesToAddress(evt.Topics[3].Bytes())
	tokenPrice := new(big.Int).SetBytes(evt.Data[32:64])
	winAmount := new(big.Int).SetBytes(evt.Data[64:])
	payToken := common.BytesToAddress(evt.Data[:32])

	// finish the auction
	finishAuction(
		&contract,
		tokenID,
		nil,
		&winner,
		winAmount,
		tokenPrice,
		&payToken,
		evt,
		lo,
	)
}

// auctionResolved processes the auction resolved event log.
// Auction::AuctionResulted(address oldOwner, address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 unitPrice, uint256 winningBid)
func auctionResolvedV2(evt *eth.Log, lo *logObserver) {
	// sanity check: 1 + 3 topics; 2 x uint256 + 2 address = 128 bytes data
	if len(evt.Data) != 128 || len(evt.Topics) != 4 {
		log.Errorf("not Auction::AuctionResultedV2() event #%d/#%d; expected 128 bytes of data, %d given; expected 4 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	contract := common.BytesToAddress(evt.Topics[1].Bytes())
	tokenID := new(big.Int).SetBytes(evt.Topics[2].Bytes())
	winner := common.BytesToAddress(evt.Topics[3].Bytes())
	owner := common.BytesToAddress(evt.Data[:32])
	tokenPrice := new(big.Int).SetBytes(evt.Data[64:96])
	winAmount := new(big.Int).SetBytes(evt.Data[96:])
	payToken := common.BytesToAddress(evt.Data[32:64])

	// finish the auction
	finishAuction(
		&contract,
		tokenID,
		&owner,
		&winner,
		winAmount,
		tokenPrice,
		&payToken,
		evt,
		lo,
	)
}

// finishAuction finalises auction on the given NFT token.
func finishAuction(contract *common.Address, tokenID *big.Int, owner *common.Address, winner *common.Address, amount *big.Int, tokenPrice *big.Int, payToken *common.Address, evt *eth.Log, _ *logObserver) {
	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}

	// pull the auction involved
	auction, err := repo.GetAuction(contract, tokenID, &evt.Address)
	if auction == nil {
		log.Errorf("resolved auction %s/%s not found; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err)
		return
	}

	// owner of the auction
	if owner == nil {
		owner = &auction.Owner
	}

	when := types.Time(time.Unix(int64(blk.Time), 0))
	auction.Resolved = &when
	auction.Closed = &when
	auction.Winner = winner
	auction.WinningBid = (*hexutil.Big)(amount)

	// store the listing into database
	if err := repo.StoreAuction(auction); err != nil {
		log.Errorf("could not store auction; %s", err.Error())
	}

	// what's the unified price the NFT was sold for
	price := repo.CalculateUnifiedPrice(payToken, amount, tokenPrice)

	// mark the token as sold
	if err := repo.TokenMarkSold(
		&auction.Contract,
		(*big.Int)(&auction.TokenId),
		&price,
		(*time.Time)(&when),
	); err != nil {
		log.Errorf("could not mark token as sold; %s", err.Error())
	}

	// log activity
	activity := types.Activity{
		Transaction:  evt.TxHash,
		OrdinalIndex: types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
		Time:         when,
		ActType:      types.EvtAuctionResolved,
		Contract:     auction.Contract,
		TokenId:      auction.TokenId,
		AuctionHall:  &evt.Address,
		Quantity:     &auction.Quantity,
		From:         *owner,
		To:           winner,
		UnitPrice:    auction.WinningBid,
		UnifiedPrice: price.Usd,
		PayToken:     payToken,
	}
	if err := repo.StoreActivity(&activity); err != nil {
		log.Errorf("could not store auction activity; %s", err.Error())
		return
	}

	// notify subscribers
	event := types.Event{Type: "AUCTION_RESOLVED", Auction: auction}
	subscriptionManager := GetSubscriptionsManager()
	subscriptionManager.PublishAuctionEvent(event)
	subscriptionManager.PublishUserEvent(auction.Owner, event)
	subscriptionManager.PublishUserEvent(*auction.Winner, event)
}
