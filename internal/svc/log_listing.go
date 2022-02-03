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
		Owner:        common.BytesToAddress(evt.Topics[1].Bytes()),
		Contract:     common.BytesToAddress(evt.Topics[2].Bytes()),
		TokenId:      hexutil.Big(*new(big.Int).SetBytes(evt.Data[:32])),
		Marketplace:  evt.Address,
		Quantity:     hexutil.Big(*new(big.Int).SetBytes(evt.Data[32:64])),
		PayToken:     common.BytesToAddress(evt.Data[64:96]),
		UnitPrice:    hexutil.Big(*new(big.Int).SetBytes(evt.Data[96:128])),
		Created:      types.Time(time.Unix(int64(blk.Time), 0)),
		StartTime:    types.Time(time.Unix(new(big.Int).SetBytes(evt.Data[128:]).Int64(), 0)),
		LastUpdate:   nil,
		Closed:       nil,
		OrdinalIndex: types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
		IsActive:     true,
	}

	tokenPrice, err := repo.GetUnifiedPrice(lst.PayToken, lst.UnitPrice)
	if err != nil {
		log.Errorf("could not convert price; %s", err.Error())
	}
	lst.UnifiedPrice = tokenPrice.Usd

	// store the listing into database
	if err := repo.StoreListing(&lst); err != nil {
		log.Errorf("could not store listing; %s", err.Error())
	}

	// mark the token as listed
	if err := repo.TokenMarkListed(
		lst.Contract,
		lst.TokenId,
		tokenPrice,
		(*time.Time)(&lst.Created),
	); err != nil {
		log.Errorf("could not mark token as listed; %s", err.Error())
	}

	// log activity
	activity := types.Activity{
		Transaction:  evt.TxHash,
		OrdinalIndex: types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
		Time:         lst.Created,
		ActType:      types.EvtListingCreated,
		Contract:     lst.Contract,
		TokenId:      lst.TokenId,
		Quantity:     &lst.Quantity,
		Marketplace:  &evt.Address,
		From:         lst.Owner,
		PayToken:     &lst.PayToken,
		UnitPrice:    &lst.UnitPrice,
		UnifiedPrice: tokenPrice.Usd,
		StartTime:    &lst.StartTime,
	}
	if err := repo.StoreActivity(&activity); err != nil {
		log.Errorf("could not store listing activity; %s", err.Error())
		return
	}

	notifyEventToOwner(types.NotifyListingCreated, lst.Contract, lst.TokenId, lst.Owner, nil, lst.Created)
	notifyEventToFollowers(types.NotifyFollowerListingAdded, lst.Contract, lst.TokenId, lst.Owner, lst.Created)

	log.Infof("added new listing of %s/%s owner %s", lst.Contract.String(), lst.TokenId.String(), lst.Owner.String())
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
	lst, err := repo.GetListing(&contract, tokenID, &owner, &evt.Address)
	if lst == nil {
		log.Errorf("update listing not found; %s", err)
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

	tokenPrice, err := repo.GetUnifiedPrice(lst.PayToken, lst.UnitPrice)
	if err != nil {
		log.Errorf("could not convert price; %s", err.Error())
	}
	lst.UnifiedPrice = tokenPrice.Usd
	lst.LastUpdate = (*types.Time)(&up)

	// store the listing into database
	if err := repo.StoreListing(lst); err != nil {
		log.Errorf("could not store listing; %s", err.Error())
	}

	// mark the token as listed
	if err := repo.TokenMarkListed(
		lst.Contract,
		lst.TokenId,
		tokenPrice,
		(*time.Time)(&lst.Created),
	); err != nil {
		log.Errorf("could not mark token as listed; %s", err.Error())
	}

	// log activity
	activity := types.Activity{
		Transaction:  evt.TxHash,
		OrdinalIndex: types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
		Time:         *lst.LastUpdate,
		ActType:      types.EvtListingUpdated,
		Contract:     lst.Contract,
		TokenId:      lst.TokenId,
		Marketplace:  &evt.Address,
		From:         lst.Owner,
		PayToken:     &lst.PayToken,
		UnitPrice:    &lst.UnitPrice,
		UnifiedPrice: tokenPrice.Usd,
	}
	if err := repo.StoreActivity(&activity); err != nil {
		log.Errorf("could not store listing activity; %s", err.Error())
		return
	}

	log.Infof("updated listing of %s/%s owner %s", lst.Contract.String(), lst.TokenId.String(), lst.Owner.String())
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
	lst, err := repo.GetListing(&contract, tokenID, &owner, &evt.Address)
	if lst == nil {
		log.Errorf("listing not found; %s", err)
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
	if err := repo.StoreListing(lst); err != nil {
		log.Errorf("could not store listing; %s", err.Error())
	}

	// mark the token as listed
	if err := repo.TokenMarkUnlisted(lst.Contract, hexutil.Big(*tokenID)); err != nil {
		log.Errorf("could not mark token as unlisted; %s", err.Error())
	}

	// log activity
	activity := types.Activity{
		Transaction:  evt.TxHash,
		OrdinalIndex: types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
		Time:         *lst.Closed,
		ActType:      types.EvtListingCancelled,
		Contract:     lst.Contract,
		TokenId:      lst.TokenId,
		Marketplace:  &evt.Address,
		From:         lst.Owner,
	}
	if err := repo.StoreActivity(&activity); err != nil {
		log.Errorf("could not store listing activity; %s", err.Error())
		return
	}

	notifyEventToOwner(types.NotifyListingCanceled, lst.Contract, lst.TokenId, lst.Owner, nil, *lst.Closed)
	log.Infof("canceled and closed listing of %s/%s owner %s", lst.Contract.String(), lst.TokenId.String(), lst.Owner.String())
}

// marketItemSold processes NFT listing being finished with sale event; an offer is resulted by the same event.
// Marketplace::ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 unitPrice, uint256 pricePerItem)
func marketItemSold(evt *eth.Log, lo *logObserver) {
	// sanity check: 1 + 3 topics; 4 x uint256 + 1 x address = 5 x 32 bytes = 160 bytes
	if len(evt.Data) != 160 || len(evt.Topics) != 4 {
		log.Errorf("not Marketplace::ItemCanceled() event #%d/#%d; expected 160 bytes of data, %d given; expected 4 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	owner := common.BytesToAddress(evt.Topics[1].Bytes())
	buyer := common.BytesToAddress(evt.Topics[2].Bytes())
	contract := common.BytesToAddress(evt.Topics[3].Bytes())
	tokenID := new(big.Int).SetBytes(evt.Data[:32])

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}

	// try to get the listing
	lst, err := repo.GetListing(&contract, tokenID, &owner, &evt.Address)
	if lst != nil {
		marketCloseListingWithSale(evt, lst, blk, lo, &buyer)
		return
	}

	// try to get an offer
	offer, err := repo.GetOffer(&contract, tokenID, &buyer, &evt.Address)
	if offer != nil {
		marketCloseOfferWithSale(evt, offer, blk, lo, &owner)
		return
	}

	notifyEventToOwner(types.NotifyNFTSold, contract, (hexutil.Big)(*tokenID), owner, &buyer, types.Time(time.Unix(int64(blk.Time), 0)))
	notifyEventToOwner(types.NotifyNFTPurchased, contract, (hexutil.Big)(*tokenID), buyer, &owner, types.Time(time.Unix(int64(blk.Time), 0)))

	log.Errorf("could not process sale of %s/%s by %s to %s", contract.String(), (*hexutil.Big)(tokenID).String(), owner.String(), buyer.String())
}

// marketCloseListingWithSale processes a listing wrap up by a sale.
func marketCloseListingWithSale(evt *eth.Log, lst *types.Listing, blk *eth.Header, _ *logObserver, buyer *common.Address) {
	up := time.Unix(int64(blk.Time), 0)
	lst.Closed = (*types.Time)(&up)
	lst.PayToken = common.BytesToAddress(evt.Data[64:96])
	lst.UnitPrice = hexutil.Big(*new(big.Int).SetBytes(evt.Data[128:]))

	tokenPrice, err := repo.GetUnifiedPrice(lst.PayToken, lst.UnitPrice)
	if err != nil {
		log.Errorf("could not convert price; %s", err.Error())
	}
	lst.UnifiedPrice = tokenPrice.Usd

	// store the listing into database
	if err := repo.StoreListing(lst); err != nil {
		log.Errorf("could not store listing; %s", err.Error())
	}

	// mark the token as sold
	if err := repo.TokenMarkSold(
		lst.Contract,
		lst.TokenId,
		&tokenPrice,
		&up,
	); err != nil {
		log.Errorf("could not mark token as sold; %s", err.Error())
	}

	// log activity
	activity := types.Activity{
		Transaction:  evt.TxHash,
		OrdinalIndex: types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
		Time:         *lst.Closed,
		ActType:      types.EvtListingSold,
		Contract:     lst.Contract,
		TokenId:      lst.TokenId,
		Marketplace:  &evt.Address,
		From:         lst.Owner,
		To:           buyer,
		PayToken:     &lst.PayToken,
		UnitPrice:    &lst.UnitPrice,
		UnifiedPrice: tokenPrice.Usd,
	}
	if err := repo.StoreActivity(&activity); err != nil {
		log.Errorf("could not store listing activity; %s", err.Error())
		return
	}

	// send email notifications
	notifyEventToOwner(types.NotifyNFTSold, lst.Contract, lst.TokenId, lst.Owner, buyer, types.Time(up))
	notifyEventToOwner(types.NotifyNFTPurchased, lst.Contract, lst.TokenId, *buyer, &lst.Owner, types.Time(up))

	log.Infof("closed sold listing of %s/%s owner %s", lst.Contract.String(), lst.TokenId.String(), lst.Owner.String())
}
