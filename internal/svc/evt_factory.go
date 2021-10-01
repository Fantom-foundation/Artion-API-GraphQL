// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"time"
)

// newNFTContract handles log event for new factory deployed ERC721/ERC1155 contract.
// event ContractCreated(address creator, address nft)
func newNFTContract(evt *eth.Log) {
	// sanity check: no additional topics; 2 x Address = 2 x 32 bytes
	if len(evt.Data) != 64 || len(evt.Topics) != 1 {
		log.Errorf("invalid event %s / %d; expected 64 bytes of data, %d given; expected 1 topic, %d given",
			evt.TxHash.String(), evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	// make NFT address
	ca := common.Address{}
	ca.SetBytes(evt.Data[32:])
	nft := types.NFTCollection{
		Address:  ca,
		IsActive: true,
	}
	log.Debugf("found NFT contract %s", ca.String())

	// load NFT details
	if err := extendNFTCollectionDetails(&nft, evt); err != nil {
		log.Criticalf("failed to load NFT collection %s; %s", nft.Address.String(), err.Error())
		return
	}

	// add the collection to persistent storage
	if err := repository.R().AddNFTCollection(&nft); err != nil {
		log.Criticalf("can not store NFT collection %s; %s", nft.Address.String(), err.Error())
		return
	}

	// add observed contract based on the collection
	addObservedContract(&nft, evt)
	log.Infof("new NFT collection %s found at %s", nft.Name, nft.Address.String())
}

// extendNFTCollectionDetails collects details of an NFT contract.
func extendNFTCollectionDetails(nft *types.NFTCollection, evt *eth.Log) (err error) {
	// NFT contract type is derived from the factory contract type
	nft.Type, err = Mgr().logObserver.contractTypeByFactory(&evt.Address)
	if err != nil {
		log.Errorf("contract %s type not known; %s", evt.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s is %s", nft.Address.String(), nft.Type)

	nft.Name, err = repo.CollectionName(&nft.Address)
	if err != nil {
		log.Errorf("%s %s name not known; %s", nft.Type, evt.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s name: %s", nft.Address.String(), nft.Name)

	nft.Symbol, err = repo.CollectionSymbol(&nft.Address)
	if err != nil {
		log.Errorf("%s %s symbol not known; %s", nft.Type, evt.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s symbol: %s", nft.Address.String(), nft.Symbol)

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		return err
	}
	nft.Created = time.Unix(int64(blk.Time), 0)
	return nil
}

// addObservedContract adds new observed contract into repository and log observer.
func addObservedContract(nft *types.NFTCollection, evt *eth.Log) {
	ca := common.Address{}
	ca.SetBytes(evt.Data[:32])

	oc := types.ObservedContract{
		Address:     nft.Address,
		Type:        nft.Type,
		Created:     nft.Created,
		Creator:     ca,
		BlockNumber: evt.BlockNumber,
		DeployedBy:  evt.TxHash,
	}

	// store observed contract into the repository
	repo.AddObservedContract(&oc)

	// let the log observer know there is a new contract it needs to monitor
	Mgr().logObserver.addObservedContract(&oc)
}
