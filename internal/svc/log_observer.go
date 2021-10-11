// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"time"
)

const (
	// obsBlocksNotificationTickInterval represents the interval
	// in which observed blocks are notified to repository.
	obsBlocksNotificationTickInterval = 15 * time.Second
)

// EventHandler represents a function used to process event log record.
type EventHandler func(*eth.Log, *logObserver)

// logObserver represents the service responsible for processing event logs of interest.
type logObserver struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the observer
	sigStop chan bool

	// outEvent represents an output channel being fed
	// with recognized block events for processing
	inEvents chan eth.Log

	// outNftTokens represents an output channel receiving new NFT tokens
	// for processing and metadata update
	outNftTokens chan *types.Token

	// currentBlock contains the number of the currently processed block
	currentBlock uint64

	// lastProcessedBlock contains the number of the last processed block
	lastProcessedBlock uint64

	// topics represents a map of topics to their respective event handlers.
	topics map[common.Hash]EventHandler

	// contracts represents a list of observed contracts.
	contracts []common.Address

	// nftTypes represents a map of types of observed NFT contracts.
	nftTypes map[common.Address]string

	// marketplace is the address of the Marketplace contract.
	marketplace *common.Address
}

// newLogObserver creates a new instance of the event logs observer service.
func newLogObserver(mgr *Manager) *logObserver {
	return &logObserver{
		mgr:          mgr,
		sigStop:      make(chan bool, 1),
		outNftTokens: make(chan *types.Token, nftMetadataUpdaterQueueCapacity),
		topics: map[common.Hash]EventHandler{
			/* Factory::event ContractCreated(address creator, address nft) */
			common.HexToHash("0x2d49c67975aadd2d389580b368cfff5b49965b0bd5da33c144922ce01e7a4d7b"): newNFTContract,

			/* erc721::event Minted(uint256 tokenId, address beneficiary, string tokenUri, address minter) */
			common.HexToHash("0x997115af5924f5e38964c6d65c804d4cb85129b65e62eb20a8ca6329dbe57e18"): erc721TokenMinted,

			/* erc721::event Transfer(address indexed from, address indexed to, uint256 indexed tokenId) */
			common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"): erc721TokenTransfer,

			/* erc1155::event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _amount) */
			common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"): erc1155TokenTransfer,

			/* erc1155::event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _amounts) */
			// common.HexToHash("0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"): erc1155BatchTransfer,

			/* erc1155::event URI(string _uri, uint256 indexed _id) */
			// common.HexToHash("0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b"): erc1155UriChanged,

			/* Marketplace::event ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime) */
			common.HexToHash("0xa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa2629"): marketNFTListed,

			/* Marketplace::event ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice) */
			common.HexToHash("0x60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd"): marketNFTUpdated,

			/* Marketplace::event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId) */
			common.HexToHash("0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158"): marketNFTUnlisted,

			/* Marketplace::event ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 unitPrice, uint256 pricePerItem) */
			common.HexToHash("0x949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d"): marketItemSold,

			/* Marketplace::event OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline) */
			common.HexToHash("0x89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f6"): marketOfferCreated,

			/* Marketplace::event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId) */
			common.HexToHash("0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7"): marketOfferCanceled,
		},
	}
}

// name provides the name fo the log observer service.
func (lo *logObserver) name() string {
	return "events observer"
}

// init configures the log observer and subscribes it with the manager.
func (lo *logObserver) init() {
	lo.inEvents = lo.mgr.blkObserver.outEvents
	lo.contracts = repo.ObservedContractsAddressList()
	lo.nftTypes = repo.NFTContractsTypeMap()
	lo.marketplace = repo.ObservedContractAddressByType("market")

	// make sure we have what we need
	if lo.marketplace == nil {
		log.Panicf("marketplace contract not found")
	}

	lo.mgr.add(lo)
}

// close signals the log observer to terminate.
func (lo *logObserver) close() {
	lo.sigStop <- true
}

// run collects incoming event logs from the channel and processes them using
//
func (lo *logObserver) run() {
	// start the notification ticker
	tick := time.NewTicker(obsBlocksNotificationTickInterval)

	defer func() {
		tick.Stop()

		close(lo.outNftTokens)
		lo.mgr.closed(lo)
	}()

	for {
		select {
		case <-lo.sigStop:
			return
		case <-tick.C:
			lo.notify()
		case evt, ok := <-lo.inEvents:
			if !ok {
				return
			}
			lo.process(&evt)
			lo.processed(&evt)
		}
	}
}

// process an incoming event
func (lo *logObserver) process(evt *eth.Log) {
	// is this an event from an observed contract?
	if !lo.isObservedContract(evt) {
		log.Debugf("event #%d / %d on foreign contract %s skipped", evt.BlockNumber, evt.Index, evt.Address.String())
		return
	}

	// get the handler
	handler, ok := lo.topics[evt.Topics[0]]
	if !ok {
		log.Criticalf("event log handler not found for event #%d / %d, topic %s", evt.BlockNumber, evt.Index, evt.Topics[0].String())
		return
	}

	// do the handler job
	log.Infof("processing event #%d/#%d at %s -> topic %s", evt.BlockNumber, evt.Index, evt.Address.String(), evt.Topics[0].String())
	handler(evt, lo)
}

// processed updates the information about the current and processed block number.
func (lo *logObserver) processed(evt *eth.Log) {
	// the last block is done
	if lo.currentBlock < evt.BlockNumber {
		lo.lastProcessedBlock = lo.currentBlock
		lo.currentBlock = evt.BlockNumber
	}
}

// notify the repository about the latest observed block, if any.
func (lo *logObserver) notify() {
	if lo.lastProcessedBlock == 0 {
		return
	}
	repo.NotifyLastObservedBlock(lo.lastProcessedBlock)
	log.Infof("last processed block is #%d", lo.lastProcessedBlock)
}

// isObservedContract checks if the given event log should be investigated and processed.
func (lo *logObserver) isObservedContract(evt *eth.Log) bool {
	for _, adr := range lo.contracts {
		if 0 == bytes.Compare(adr.Bytes(), evt.Address.Bytes()) {
			return true
		}
	}
	return false
}

// addObservedContract is used to extend the list of observed contracts
// with a newly created NFT contract address; subsequent NFT events should be observed on it.
func (lo *logObserver) addObservedContract(oc *types.ObservedContract) {
	// check if the contract is actually a new one
	for _, adr := range lo.contracts {
		if 0 == bytes.Compare(adr.Bytes(), oc.Address.Bytes()) {
			return
		}
	}

	// add the contract to the list
	lo.contracts = append(lo.contracts, oc.Address)

	// an NFT contract? add it to the types map as well
	if oc.Type == types.ContractTypeERC721 || oc.Type == types.ContractTypeERC1155 {
		lo.nftTypes[oc.Address] = oc.Type
	}

	log.Infof("new contract %s is now observed", oc.Address.String())
}

// topicsList provides a list of observed topics for blocks event filtering
func (lo *logObserver) topicsList() [][]common.Hash {
	list := make([][]common.Hash, 1)
	list[0] = make([]common.Hash, len(lo.topics))

	var i int
	for h := range lo.topics {
		list[0][i] = h
		i++
	}

	return list
}

// contractTypeByAddress provides type of contract based on known address.
// We use pre-loaded NFT types map to perform this.
func (lo *logObserver) contractTypeByAddress(fa *common.Address) (string, error) {
	tp, ok := lo.nftTypes[*fa]
	if !ok {
		return "", fmt.Errorf("address %s unknown", fa.String())
	}
	return tp, nil
}
