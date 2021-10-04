// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"time"
)

const (
	// logEventQueueCapacity represents the log events queue capacity.
	logEventQueueCapacity = 25000

	// observedBlocksCapacity represents the capacity of channel for observed block IDs.
	observedBlocksCapacity = 10000
)

// blkObserver represents a service monitoring incoming blocks
// and pulling event logs of those blocks for processing.
type blkObserver struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the observer
	sigStop chan bool

	// inBlock represents an input channel receiving block headers to be observed
	inBlocks chan *eth.Header

	// outEvent represents an output channel being fed
	// with recognized block events for processing
	outEvents chan *eth.Log

	// topics represents the topics observed by the API server
	topics []common.Hash

	// outObservedBlock is fed with numbers of processed blocks.
	outObservedBlocks chan uint64
}

// newBlkObserver creates a new instance of the block observer service.
func newBlkObserver(mgr *Manager) *blkObserver {
	return &blkObserver{
		mgr:               mgr,
		sigStop:           make(chan bool, 1),
		outEvents:         make(chan *eth.Log, logEventQueueCapacity),
		outObservedBlocks: make(chan uint64, observedBlocksCapacity),
		topics:            nil,
	}
}

// name provides the name of the service
func (bo *blkObserver) name() string {
	return "block observer"
}

// init prepares the block observer to capture blocks.
func (bo *blkObserver) init() {
	bo.inBlocks = bo.mgr.blkRouter.outBlocks
	bo.topics = bo.mgr.logObserver.topicsList()
	bo.mgr.add(bo)
}

// run pulls block headers from the input queue and processes them.
func (bo *blkObserver) run() {
	// start the notification ticker
	tick := time.NewTicker(obsBlocksNotificationTickInterval)

	defer func() {
		tick.Stop()
		close(bo.outEvents)
		close(bo.outObservedBlocks)
		bo.mgr.closed(bo)
	}()

	for {
		select {
		case <-bo.sigStop:
			return
		case hdr, ok := <-bo.inBlocks:
			if !ok {
				return
			}
			bo.process(hdr)
		}
	}
}

// close signals the block observer to terminate
func (bo *blkObserver) close() {
	bo.sigStop <- true
}

// process an incoming block header by investigating its events.
func (bo *blkObserver) process(hdr *eth.Header) {
	// pull events for the block
	logs, err := repo.BlockLogs(hdr.Number, bo.topics)
	if err != nil {
		log.Errorf("block #%d event logs not available; %s", hdr.Number.Uint64(), err.Error())
		return
	}

	// any logs?
	if 0 < len(logs) {
		log.Infof("processing %d events on block #%d", len(logs), hdr.Number.Uint64())
	}

	// push interesting events into the output queue, if any
	for _, evt := range logs {
		select {
		case bo.outEvents <- &evt:
		case <-bo.sigStop:
			bo.sigStop <- true
			return
		}
	}

	// notify the scanner we did process this block
	// the scanner uses the info to decide if the server keeps up
	// with the top head of the blockchain
	select {
	case bo.outObservedBlocks <- hdr.Number.Uint64():
	default:
	}
}
