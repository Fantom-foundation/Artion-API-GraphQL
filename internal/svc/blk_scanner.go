// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	eth "github.com/ethereum/go-ethereum/core/types"
	"time"
)

const (
	// blkIsScanning represents the state of active block scanning
	blkIsScanning = iota

	// blkIsIdling represents the state of passive head checks
	blkIsIdling

	// scanTickFrequency is the time delay for the regular scanner
	scanTickFrequency = 5 * time.Millisecond

	// idleTickFrequency is the time delay for the regular scanner
	idleTickFrequency = 1 * time.Second

	// topUpdateTickFrequency is the time delay for the block head update
	topUpdateTickFrequency = 5 * time.Second

	// blockQueueCapacity represents the capacity of block headers processor queue
	blockQueueCapacity = 5000

	// defStartingBlockNumber represents the first block we will scan from
	// if the previous state is unknown.
	defStartingBlockNumber = 16000000

	// blkScannerHysteresis represent the number of blocks we let slide
	// until we switch back to active scan state.
	blkScannerHysteresis = 10
)

// blkScanner represents a scanner of historical data from the blockchain.
type blkScanner struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the router
	sigStop chan bool

	// inObservedBlocks is a channel receiving IDs of observed blocks
	// we track the observed heads to recognize if we need to switch back to scan from idle
	inObservedBlocks chan uint64

	// inRescanBlocks is a channel receiving re-scan requests from given block number
	inRescanBlocks chan uint64

	// outBlocks represents a channel fed with past block headers being scanned.
	outBlocks chan *eth.Header

	// scanTicker represents the ticker for the scanner
	scanTicker *time.Ticker

	// state represents the current state of the scanner
	// it's scanning by default and turns idle later, if not needed
	state int

	// current represents the ID of the currently processed block
	current uint64

	// target represents the ID we need to reach
	target uint64

	// lastProcessedBlock represents the ID of the last processed block notified to us
	lastProcessedBlock uint64
}

// newBlkScanner creates a new instance of the block scanner service.
func newBlkScanner(mgr *Manager) *blkScanner {
	return &blkScanner{
		mgr:       mgr,
		sigStop:   make(chan bool, 1),
		outBlocks: make(chan *eth.Header, blockQueueCapacity),
	}
}

// name provides the name of the service.
func (bs *blkScanner) name() string {
	return "block scanner"
}

// init initializes the block scanner and registers it with the manager.
func (bs *blkScanner) init() {
	bs.inObservedBlocks = bs.mgr.logObserver.outObservedBlocks
	bs.inRescanBlocks = bs.mgr.collectionValidator.outRescanQueue

	bs.current, bs.target = bs.start(), bs.top()
	bs.mgr.add(bs)
}

// close signals the block observer to terminate
func (bs *blkScanner) close() {
	bs.sigStop <- true
}

// run scans past blocks one by one until it reaches top
// after the top is reached, it idles and checks the head state to make sure
// the API server keeps up with the most recent block
func (bs *blkScanner) run() {
	// make tickers
	topTick := time.NewTicker(topUpdateTickFrequency)
	bs.scanTicker = time.NewTicker(scanTickFrequency)

	defer func() {
		topTick.Stop()
		bs.scanTicker.Stop()

		bs.mgr.closed(bs)
	}()

	for {
		// make sure to check for terminate; but do not stay in
		select {
		case <-bs.sigStop:
			return

		case <-topTick.C:
			bs.target = bs.top()
			bs.notify()

		case bid, ok := <-bs.inObservedBlocks:
			if !ok {
				return
			}

			// we just casually follow the chain head
			if bs.state == blkIsIdling && bid > bs.current {
				bs.current = bid
				bs.lastProcessedBlock = bid
				continue
			}

			// we rush to catch the head, so we don't accept processed blocks above scanner head
			if bid > bs.lastProcessedBlock && bid <= bs.current {
				bs.lastProcessedBlock = bid
			}

		case bid, ok := <-bs.inRescanBlocks:
			if !ok {
				return
			}
			bs.rescan(bid)

		case <-bs.scanTicker.C:
		}

		bs.next()
		bs.checkTarget()
		bs.checkIdle()
	}
}

// rescan the blockchain from the given block, if relevant.
func (bs *blkScanner) rescan(from uint64) {
	// are we already on the track?
	if from > bs.current {
		return
	}

	// refresh target block
	bs.target = bs.top()

	// we know from <= current here; start at least <blkScannerHysteresis> back
	diff := bs.current - from
	if diff < blkScannerHysteresis {
		bs.current = bs.current - blkScannerHysteresis
		return
	}
	bs.current = from
}

// scanNext tries to advance the scanner to the next block, if possible
func (bs *blkScanner) next() {
	if bs.state == blkIsScanning && bs.current <= bs.target {
		hdr, err := repo.GetHeader(bs.current)
		if err != nil {
			log.Errorf("block header #%s not available; %s", bs.current, err.Error())
			return
		}

		// send the block to the observer; make sure not to miss stop signal
		select {
		case bs.outBlocks <- hdr:
			bs.current += 1
		case <-bs.sigStop:
			bs.sigStop <- true
		}
	}
}

// checkTarget checks if the scanner reached designated target head.
func (bs *blkScanner) checkTarget() {
	// reached target? make sure we are on target; switch state if so
	if bs.state == blkIsScanning && bs.current > bs.target {
		bs.target = bs.top()
		diff := int64(bs.target) - int64(bs.current)

		if diff <= 0 {
			bs.state = blkIsIdling
			bs.scanTicker.Reset(idleTickFrequency)
			log.Noticef("scanner idling since #%d", bs.current)
		}
	}
}

// checkIdle checks if the idle state should be switched back to active scan.
func (bs *blkScanner) checkIdle() {
	if bs.state != blkIsIdling {
		return
	}

	diff := int64(bs.target) - int64(bs.current)
	if diff >= blkScannerHysteresis {
		bs.state = blkIsScanning
		bs.scanTicker.Reset(scanTickFrequency)
		log.Noticef("scanner head at #%d of #%d with %d diff", bs.current, bs.target, diff)
	}
}

// start provides the starting point for the scanner
func (bs *blkScanner) start() uint64 {
	lnb, err := repo.LastSeenBlockNumber()
	if err != nil {
		log.Criticalf("can not pull the previous state; %s", err.Error())
		return 0
	}

	// if the state is unknown, use the lowest observed block number
	// use a default starting block number, if the query fails
	// we don't need to start scanning from the absolute start of the chain
	if lnb == 0 {
		lnb = repo.MinObservedBlockNumber(defStartingBlockNumber)
	}

	if lnb <= blkScannerHysteresis {
		return lnb
	}
	return lnb - blkScannerHysteresis
}

// top provides the number of the target block for the scanner.
func (bs *blkScanner) top() uint64 {
	cur, err := repo.CurrentHead()
	if err != nil {
		log.Criticalf("can not pull the latest head number; %s", err.Error())
		return 0
	}

	return cur
}

// notify the repository about the latest observed block, if any.
func (bs *blkScanner) notify() {
	if bs.lastProcessedBlock == 0 {
		return
	}
	repo.NotifyLastObservedBlock(bs.lastProcessedBlock)

	if bs.state == blkIsIdling {
		log.Infof("idle at #%d, head at #%d", bs.current, bs.target)
		return
	}
	log.Infof("scanner at #%d of #%d; processed #%d", bs.current, bs.target, bs.lastProcessedBlock)
}
