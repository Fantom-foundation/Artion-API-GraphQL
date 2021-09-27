// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/svc/blkcache"
	eth "github.com/ethereum/go-ethereum/core/types"
)

const (
	// outBlockQueueCapacity represents the size of outgoing block queue.
	outBlockQueueCapacity = 1000

	// blkCacheCapacity represents the amount of blocks kept in cache to prevent head miss
	blkCacheCapacity = 25
)

// blkRouter represents a router of blockchain heads records
// coming from the RPC heads subscription, or the block scanner service
type blkRouter struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the router
	sigStop chan bool

	// inNewBlock represents an input channel receiving new block headers from the blockchain
	inNewBlocks chan *eth.Header

	// inScanBlock represents an input channel receiving blocks from the scanner service
	inScanBlocks chan *eth.Header

	// inScanStateChange represents the channel receiving scanner state change notifications
	inScanStateChange chan int

	// outBlock represents an output channel being fed
	// with block headers to be processed by the block observer service
	outBlocks chan *eth.Header

	// cache is a circular cache used to store incoming blocks on scan state to prevent block loss
	cache *blkcache.Cache
}

// newBlkRouter creates a new instance of the block router service.
func newBlkRouter(mgr *Manager) *blkRouter {
	return &blkRouter{
		mgr:       mgr,
		sigStop:   make(chan bool, 1),
		outBlocks: make(chan *eth.Header, outBlockQueueCapacity),
		cache:     blkcache.NewCache(blkCacheCapacity),
	}
}

// name provides the name of the service
func (br *blkRouter) name() string {
	return "block router"
}

// init prepares the block router to process incoming blocks from multiple sources.
func (br *blkRouter) init() {
	br.inNewBlocks = repository.R().NewHeaders()
	br.inScanBlocks = br.mgr.blkScanner.outBlocks
	br.inScanStateChange = br.mgr.blkScanner.outStateChange
	br.mgr.add(br)
}

// run pulls block headers from multiple sources and routes based on the API server state.
func (br *blkRouter) run() {

}

// close signals the block observer to terminate
func (br *blkRouter) close() {
	br.sigStop <- true
}
