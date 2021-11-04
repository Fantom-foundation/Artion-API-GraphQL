// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"time"
)

// reScanBlockQueueCapacity represents the capacity of re-scan queue.
const reScanBlockQueueCapacity = 100

// nftCollectionValidator represents a service responsible for processing new NFT collections
// by validating them, adding them into the service observers and making sure new collections
// are properly scanned and registered.
type nftCollectionValidator struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the observer
	sigStop chan bool

	// inCollections represents an input channel being fed with addresses
	// of newly registered collection contracts.
	inCollections chan common.Address

	// outRescanQueue is the channel being fed with blocks to be re-visited by the scanner
	outRescanQueue chan uint64
}

// newNFTCollectionValidator creates a new instance of the NFT validator.
func newNFTCollectionValidator(mgr *Manager) *nftCollectionValidator {
	return &nftCollectionValidator{
		mgr:            mgr,
		sigStop:        make(chan bool, 1),
		outRescanQueue: make(chan uint64, reScanBlockQueueCapacity),
	}
}

// name provides the name fo the log observer service.
func (cv *nftCollectionValidator) name() string {
	return "collection validator"
}

// init configures the NFT collection validator and subscribes it with the manager.
func (cv *nftCollectionValidator) init() {
	// link channels
	cv.inCollections = repo.NewCollectionQueue()

	// add and run
	cv.mgr.add(cv)
}

// close signals the NFT collection validator to terminate.
func (cv *nftCollectionValidator) close() {
	cv.sigStop <- true
}

// run collects incoming NFT collections and processes them through.
func (cv *nftCollectionValidator) run() {
	defer func() {
		cv.mgr.closed(cv)
	}()

	for {
		select {
		case <-cv.sigStop:
			return
		case adr, ok := <-cv.inCollections:
			if !ok {
				return
			}
			cv.process(&adr)
		}
	}
}

// process new collection to be added
func (cv *nftCollectionValidator) process(adr *common.Address) {
	// try to get the collection structure for the address
	col, blk, err := cv.collection(adr)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	// add the collection to persistent storage
	if err = repo.StoreCollection(col); err != nil {
		log.Criticalf("can not store NFT collection %s; %s", col.Address.String(), err.Error())
		return
	}

	// observe the contract from now on
	evt := new(eth.Log)
	evt.BlockNumber = blk
	evt.TxHash = common.Hash{}
	addObservedContract(col, evt)

	select {
	case cv.outRescanQueue <- blk:
		log.Infof("added collection %s found at %s on block %d", col.Name, col.Address.String(), blk)
	default:
		log.Infof("can not re-scan collection %s found at %s from block %d", col.Name, col.Address.String(), blk)
	}
}

// collection checks the address to be a valid ERC-721 or ERC-1155 contract
// and if so, it builds a Collection structure to be added into the API to follow.
func (cv *nftCollectionValidator) collection(adr *common.Address) (col *types.Collection, blkNumber uint64, err error) {
	log.Infof("validating contract at %s", adr.String())

	col = &types.Collection{
		Address: *adr,
	}

	// detect ERC-721 contract
	if repo.IsErc721Contract(adr) {
		col.Type = types.ContractTypeERC721
		blkNumber, err = repo.Erc721StartingBlockNumber(adr)
	}

	// detect ERC-1155 contract
	if col.Type == "" && repo.IsErc1155Contract(adr) {
		col.Type = types.ContractTypeERC1155
		blkNumber, err = repo.Erc1155StartingBlockNumber(adr)
	}

	// unknown contract type?
	if err != nil || col.Type == "" {
		return nil, 0, fmt.Errorf("%s is not supported NFT contract", adr.String())
	}

	return cv.collectionDetails(col, blkNumber)
}

// collectionDetails loads details of the provided NFT collection contract.
func (cv *nftCollectionValidator) collectionDetails(col *types.Collection, blk uint64) (*types.Collection, uint64, error) {
	// extend metadata
	if err := extendCollectionMetadata(col); err != nil {
		return nil, 0, err
	}

	// find the block
	head, err := repo.GetHeader(blk)
	if err != nil {
		return nil, 0, err
	}

	col.Created = types.Time(time.Unix(int64(head.Time), 0))
	return col, blk, nil
}
