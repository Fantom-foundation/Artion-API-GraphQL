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

// marketNFTListed handles log event for NFT token to get listed for sale on the Marketplace.
// Marketplace::ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func marketNFTListed(evt *eth.Log, _ *logObserver) {
	// sanity check: 1 + 2 topics; 4 x uint256 + 1 address = 5 x 32 bytes of data = 160 bytes
	if len(evt.Data) != 160 || len(evt.Topics) != 3 {
		log.Errorf("not Marketplace::ItemListed() event #%d/#%d; expected 160 bytes of data, %d given; expected 3 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}

	// make the listing
	lst := types.Listing{
		Owner:      common.BytesToAddress(evt.Topics[1].Bytes()),
		Contract:   common.BytesToAddress(evt.Topics[2].Bytes()),
		TokenId:    hexutil.Big(*new(big.Int).SetBytes(evt.Data[:32])),
		Quantity:   hexutil.Big(*new(big.Int).SetBytes(evt.Data[32:64])),
		PayToken:   common.BytesToAddress(evt.Data[64:96]),
		UnitPrice:  hexutil.Big(*new(big.Int).SetBytes(evt.Data[96:128])),
		Created:    types.Time(time.Unix(int64(blk.Time), 0)),
		StartTime:  types.Time(time.Unix(new(big.Int).SetBytes(evt.Data[128:]).Int64(), 0)),
		LastUpdate: nil,
		Closed:     nil,
	}
	lst.OrdinalIndex = types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index))

	// store the listing into database
	if err := repo.ListingStore(&lst); err != nil {
		log.Errorf("could not store listing; %s", err.Error())
	}

	// mark the token as listed
	if err := repo.TokenMarkListed(&lst.Contract, (*big.Int)(&lst.TokenId), &lst.UnitPrice, (*time.Time)(&lst.Created)); err != nil {
		log.Errorf("could not mark token as listed; %s", err.Error())
	}
}

// marketNFTUpdated handles an update call on already listed NFT token.
// Marketplace::ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice)
func marketNFTUpdated(evt *eth.Log, _ *logObserver) {
	// sanity check: 1 + 2 topics; 2 x uint256 + 1 address = 3 x 32 bytes of data = 96 bytes
	if len(evt.Data) != 96 || len(evt.Topics) != 3 {
		log.Errorf("not Marketplace::ItemUpdated() event #%d/#%d; expected 96 bytes of data, %d given; expected 3 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	owner := common.BytesToAddress(evt.Topics[1].Bytes())
	contract := common.BytesToAddress(evt.Topics[2].Bytes())
	tokenID := new(big.Int).SetBytes(evt.Data[:32])

	// try to get the listing
	lst, err := repo.ListingGet(&contract, tokenID, &owner)
	if err != nil {
		log.Errorf("update listing not found; %s", err.Error())
		return
	}

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}
	up := time.Unix(int64(blk.Time), 0)

	// do the update
	lst.PayToken = common.BytesToAddress(evt.Data[32:64])
	lst.UnitPrice = hexutil.Big(*new(big.Int).SetBytes(evt.Data[64:]))
	lst.LastUpdate = (*types.Time)(&up)

	// store the listing into database
	if err := repo.ListingStore(lst); err != nil {
		log.Errorf("could not store listing; %s", err.Error())
	}

	// mark the token as listed
	if err := repo.TokenMarkListed(&lst.Contract, (*big.Int)(&lst.TokenId), &lst.UnitPrice, (*time.Time)(&lst.Created)); err != nil {
		log.Errorf("could not mark token as listed; %s", err.Error())
	}
}

// marketNFTUnlisted processes canceled NFT listing event.
// Marketplace::ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId)
func marketNFTUnlisted(evt *eth.Log, _ *logObserver) {
	// sanity check: 1 + 2 topics; 1 x uint256 = 32 bytes
	if len(evt.Data) != 32 || len(evt.Topics) != 3 {
		log.Errorf("not Marketplace::ItemCanceled() event #%d/#%d; expected 32 bytes of data, %d given; expected 3 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	owner := common.BytesToAddress(evt.Topics[1].Bytes())
	contract := common.BytesToAddress(evt.Topics[2].Bytes())
	tokenID := new(big.Int).SetBytes(evt.Data[:])

	// try to get the listing
	lst, err := repo.ListingGet(&contract, tokenID, &owner)
	if err != nil {
		log.Errorf("listing not found; %s", err.Error())
		return
	}

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}
	up := time.Unix(int64(blk.Time), 0)
	lst.Closed = (*types.Time)(&up)

	// store the listing into database
	if err := repo.ListingStore(lst); err != nil {
		log.Errorf("could not store listing; %s", err.Error())
	}

	// mark the token as listed
	if err := repo.TokenMarkUnlisted(&lst.Contract, tokenID); err != nil {
		log.Errorf("could not mark token as listed; %s", err.Error())
	}
}

// marketNFTSold processes NFT listing being finished with sale event.
// Marketplace::ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 unitPrice, uint256 pricePerItem)
func marketNFTSold(evt *eth.Log, _ *logObserver) {
	// sanity check: 1 + 3 topics; 4 x uint256 + 1 x address = 5 x 32 bytes = 160 bytes
	if len(evt.Data) != 160 || len(evt.Topics) != 4 {
		log.Errorf("not Marketplace::ItemCanceled() event #%d/#%d; expected 160 bytes of data, %d given; expected 4 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	owner := common.BytesToAddress(evt.Topics[1].Bytes())
	contract := common.BytesToAddress(evt.Topics[3].Bytes())
	tokenID := new(big.Int).SetBytes(evt.Data[:32])

	// try to get the listing
	lst, err := repo.ListingGet(&contract, tokenID, &owner)
	if err != nil {
		log.Errorf("listing not found; %s", err.Error())
		return
	}

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}
	up := time.Unix(int64(blk.Time), 0)

	lst.Closed = (*types.Time)(&up)
	lst.PayToken = common.BytesToAddress(evt.Data[64:96])
	lst.UnitPrice = hexutil.Big(*new(big.Int).SetBytes(evt.Data[128:]))

	// store the listing into database
	if err := repo.ListingStore(lst); err != nil {
		log.Errorf("could not store listing; %s", err.Error())
	}

	// mark the token as sold
	if err := repo.TokenMarkSold(&lst.Contract, tokenID, &lst.UnitPrice, &up); err != nil {
		log.Errorf("could not mark token as listed; %s", err.Error())
	}
}
