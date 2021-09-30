// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/svc/blkcache"
	eth "github.com/ethereum/go-ethereum/core/types"
)

const (
	// outBlockQueueCapacity represents the size of outgoing block queue.
	outBlockQueueCapacity = 10000

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

	// state represents the routing state received from the block scanner
	state int
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
	br.inNewBlocks = repo.NewHeaders()
	br.inScanBlocks = br.mgr.blkScanner.outBlocks
	br.inScanStateChange = br.mgr.blkScanner.outStateChange
	br.mgr.add(br)
}

// close signals the block observer to terminate
func (br *blkRouter) close() {
	br.sigStop <- true
}

// run pulls block headers from multiple sources and routes based on the API server state.
func (br *blkRouter) run() {
	defer func() {
		br.mgr.closed(br)
	}()

	for {
		select {
		case <-br.sigStop:
			return
		case st := <-br.inScanStateChange:
			br.state = st
			if st == blkIsIdling {
				br.cleanCache()
			}
		case hdr := <-br.inNewBlocks:
			if br.state == blkIsScanning {
				br.cache.Add(hdr)
				log.Debugf("block #%d cached", hdr.Number.Uint64())
			} else {
				br.push(hdr)
			}
		case hdr := <-br.inScanBlocks:
			br.push(hdr)
		}
	}
}

// push the given block header to the output channel, make sure not to mis stop signal.
func (br *blkRouter) push(hdr *eth.Header) {
	select {
	case <-br.sigStop:
		br.sigStop <- true
	case br.outBlocks <- hdr:
		log.Debugf("block #%d pushed for processing", hdr.Number.Uint64())
	}
}

// cleanCache cleans the block headers cache into the output channel on scanner idle state switch.
func (br *blkRouter) cleanCache() {
	cl := br.cache.Pull()
	for _, hdr := range cl {
		select {
		case <-br.sigStop:
			br.sigStop <- true
			return
		case br.outBlocks <- hdr:
			log.Debugf("block #%d unloaded from cache", hdr.Number.Uint64())
		}
	}
}
