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
func auctionCreated(evt *eth.Log, _ *logObserver) {
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
		ReservePrice:  hexutil.Big{},
		Created:       types.Time(time.Unix(int64(blk.Time), 0)),
		StartTime:     types.Time{},
		EndTime:       types.Time{},
		Closed:        nil,
		LastBidPlaced: nil,
		OrdinalIndex:  types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
	}

	// extend the auction with details pulled from the contract

	// store the listing into database
	if err := repo.StoreAuction(&auction); err != nil {
		log.Errorf("could not store auction; %s", err.Error())
	}
}
