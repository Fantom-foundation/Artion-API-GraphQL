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

// marketNFTListed handles log event for NFT token to receive buy offer on the Marketplace.
// Marketplace::OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline)
func marketOfferCreated(evt *eth.Log, lo *logObserver) {
	// sanity check: 1 + 2 topics; 4 x uint256 + 1 address = 5 x 32 bytes of data = 160 bytes
	if len(evt.Data) != 160 || len(evt.Topics) != 3 {
		log.Errorf("not Marketplace::OfferCreated() event #%d/#%d; expected 160 bytes of data, %d given; expected 3 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}

	// create the offer record
	offer := types.Offer{
		Contract:     common.BytesToAddress(evt.Topics[2].Bytes()),
		TokenId:      hexutil.Big(*new(big.Int).SetBytes(evt.Data[:32])),
		ProposedBy:   common.BytesToAddress(evt.Topics[1].Bytes()),
		Quantity:     hexutil.Big(*new(big.Int).SetBytes(evt.Data[32:64])),
		PayToken:     common.BytesToAddress(evt.Data[64:96]),
		UnitPrice:    hexutil.Big(*new(big.Int).SetBytes(evt.Data[96:128])),
		Created:      types.Time(time.Unix(int64(blk.Time), 0)),
		Deadline:     types.Time(time.Unix(new(big.Int).SetBytes(evt.Data[128:]).Int64(), 0)),
		Closed:       nil,
		OrdinalIndex: types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
	}

	// store the listing into database
	if err := repo.StoreOffer(&offer); err != nil {
		log.Errorf("could not store offer; %s", err.Error())
	}

	// mark the token as listed
	if err := repo.TokenMarkOffered(
		&offer.Contract,
		(*big.Int)(&offer.TokenId),
		repo.GetUnitPriceAt(lo.marketplace, &offer.PayToken, new(big.Int).SetUint64(evt.BlockNumber), (*big.Int)(&offer.UnitPrice)),
		(*time.Time)(&offer.Created),
	); err != nil {
		log.Errorf("could not mark token as having offer; %s", err.Error())
	}

	log.Infof("added new offer of %s/%s proposed by %s", offer.Contract.String(), offer.TokenId.String(), offer.ProposedBy.String())

	// notify subscribers (will be skipped for tokens owned by 100 or more owners)
	event := types.Event{ Type: "OFFER_CREATED", Offer: &offer }
	subscriptionManager := GetSubscriptionsManager()
	owners, err := repo.ListOwnerships(&offer.Contract, &offer.TokenId, nil, "", 100, false)
	if err != nil {
		log.Errorf("failed to list owner to get subscribers for token; %s", err)
	} else if len(owners.Collection) < 100 {
		for _, ownership := range owners.Collection {
			subscriptionManager.PublishUserEvent(ownership.Owner, event)
		}
	}
}

// marketOfferCanceled handles log event for NFT token to loose buy offer on the Marketplace.
// Marketplace::OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId)
func marketOfferCanceled(evt *eth.Log, _ *logObserver) {
	// sanity check: 1 + 2 topics; 1 x uint256 = 32 bytes
	if len(evt.Data) != 32 || len(evt.Topics) != 3 {
		log.Errorf("not Marketplace::OfferCanceled() event #%d/#%d; expected 32 bytes of data, %d given; expected 3 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	proposer := common.BytesToAddress(evt.Topics[1].Bytes())
	contract := common.BytesToAddress(evt.Topics[2].Bytes())
	tokenID := new(big.Int).SetBytes(evt.Data[:])

	// try to get the offer being canceled
	offer, err := repo.GetOffer(&contract, tokenID, &proposer)
	if err != nil {
		log.Errorf("offer not found; %s", err.Error())
		return
	}

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("could not get header #%d, %s", evt.BlockNumber, err.Error())
		return
	}
	up := time.Unix(int64(blk.Time), 0)
	offer.Closed = (*types.Time)(&up)

	// store the offer back into database
	if err := repo.StoreOffer(offer); err != nil {
		log.Errorf("could not update listing; %s", err.Error())
	}

	// mark the token as listed
	if err := repo.TokenMarkUnOffered(&offer.Contract, tokenID); err != nil {
		log.Errorf("could not mark token as not having offer; %s", err.Error())
	}

	log.Infof("canceled offer on %s/%s proposed by %s", offer.Contract.String(), offer.TokenId.String(), offer.ProposedBy.String())
}

// marketCloseOfferWithSale processes a listing wrap up by a sale.
func marketCloseOfferWithSale(evt *eth.Log, offer *types.Offer, blk *eth.Header, lo *logObserver) {
	up := time.Unix(int64(blk.Time), 0)
	offer.Closed = (*types.Time)(&up)
	offer.PayToken = common.BytesToAddress(evt.Data[64:96])
	offer.UnitPrice = hexutil.Big(*new(big.Int).SetBytes(evt.Data[128:]))

	// store the listing into database
	if err := repo.StoreOffer(offer); err != nil {
		log.Errorf("could not store offer; %s", err.Error())
	}

	// mark the token as sold
	if err := repo.TokenMarkSold(
		&offer.Contract,
		(*big.Int)(&offer.TokenId),
		repo.GetUnitPriceAt(lo.marketplace, &offer.PayToken, new(big.Int).SetUint64(evt.BlockNumber), (*big.Int)(&offer.UnitPrice)),
		&up,
	); err != nil {
		log.Errorf("could not mark token as sold; %s", err.Error())
	}

	log.Infof("closed buy offer of %s/%s proposed by %s", offer.Contract.String(), offer.TokenId.String(), offer.ProposedBy.String())
}
