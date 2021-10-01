// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

const (
	// obsBlocksNotificationTickInterval represents the interval
	// in which observed blocks are notified to repository.
	obsBlocksNotificationTickInterval = 10 * time.Second
)

// EventHandler represents a function used to process event log record.
type EventHandler func(*eth.Log)

// logObserver represents the service responsible for processing event logs of interest.
type logObserver struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the observer
	sigStop chan bool

	// outEvent represents an output channel being fed
	// with recognized block events for processing
	inEvents chan *eth.Log

	// currentBlock contains the number of the currently processed block
	currentBlock *big.Int

	// lastProcessedBlock contains the number of the last processed block
	lastProcessedBlock *big.Int

	// topics represents a map of topics to their respective event handlers.
	topics map[common.Hash]EventHandler

	// contracts represents a list of observed contracts.
	contracts []common.Address

	// nftTypes represents a map of types of observed NFT contracts.
	nftTypes map[common.Address]string
}

// newLogObserver creates a new instance of the event logs observer service.
func newLogObserver(mgr *Manager) *logObserver {
	return &logObserver{
		mgr:     mgr,
		sigStop: make(chan bool, 1),
		topics: map[common.Hash]EventHandler{
			/* event ContractCreated(address creator, address nft) */
			common.HexToHash("0x2d49c67975aadd2d389580b368cfff5b49965b0bd5da33c144922ce01e7a4d7b"): newNFTContract,
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
		lo.mgr.closed(lo)
	}()

	for {
		select {
		case <-lo.sigStop:
			return
		case <-tick.C:
			lo.notify()
		case evt := <-lo.inEvents:
			lo.process(evt)
			lo.processed(evt)
		}
	}
}

// process an incoming event
func (lo *logObserver) process(evt *eth.Log) {
	// is this an event from an observed contract?
	if !lo.isObservedContract(evt) {
		log.Infof("event #%d / %d on foreign contract %s skipped", evt.BlockNumber, evt.Index, evt.Address.String())
		return
	}

	// get the handler
	handler, ok := lo.topics[evt.Topics[0]]
	if !ok {
		log.Criticalf("event log handler not found for event #%d / %d, topic %s", evt.BlockNumber, evt.Index, evt.Topics[0].String())
		return
	}

	// do the handler job
	handler(evt)
}

// processed updates the information about the current and processed block number.
func (lo *logObserver) processed(evt *eth.Log) {
	if lo.currentBlock == nil {
		lo.currentBlock = new(big.Int).SetUint64(evt.BlockNumber)
		return
	}

	// the last block is done
	if lo.currentBlock.Uint64() < evt.BlockNumber {
		lo.lastProcessedBlock = lo.currentBlock
		lo.currentBlock = new(big.Int).SetUint64(evt.BlockNumber)
	}
}

// notify the repository about the latest observed block, if any.
func (lo *logObserver) notify() {
	if lo.lastProcessedBlock == nil {
		return
	}
	repo.NotifyLastObservedBlock(lo.lastProcessedBlock)
	log.Infof("last processed block is #%d", lo.lastProcessedBlock.Uint64())
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
func (lo *logObserver) topicsList() []common.Hash {
	list := make([]common.Hash, len(lo.topics))

	var i int
	for h := range lo.topics {
		list[i] = h
		i++
	}

	return list
}

// contractTypeByFactory provides type of contract based on known factory address.
// We use pre-loaded NFT types map to perform this.
func (lo *logObserver) contractTypeByFactory(fa *common.Address) (string, error) {
	tp, ok := lo.nftTypes[*fa]
	if !ok {
		return "", fmt.Errorf("factory %s unknown", fa.String())
	}
	return tp, nil
}
