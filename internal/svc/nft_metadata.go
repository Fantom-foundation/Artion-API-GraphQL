// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"time"
)

const (
	// nftMetadataWorkerQueueCapacity is the capacity of NFT metadata worker queue.
	nftMetadataWorkerQueueCapacity = 200

	// nftMetadataRefreshTick is the tick used to queue NFT metadata refresh candidates.
	nftMetadataRefreshTick = 15 * time.Second

	// nftMetadataRefreshPullTick is the tick used to pull NFT metadata refresh candidates.
	nftMetadataRefreshPullTick = 200 * time.Millisecond

	// nftMetadataRefreshSetSize is the max size of metadata refresh set pulled at once.
	nftMetadataRefreshSetSize = 100
)

// nftMetadataUpdater represents a service responsible for periodic update of NFT token metadata
// from remote URI/IFS in a local persistent storage. The metadata download is time
// costly operation, so we don't do it during the NFT token detection.
type nftMetadataUpdater struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the router
	sigStop chan bool

	// inTokens represents the channel receiving tokens to be updated.
	inTokens chan *types.Token

	// outTokens represents the channel receiving NFTs for actual processing.
	outTokens chan *types.Token

	// refreshQueue is the queue used to store NFT metadata update candidates.
	refreshQueue chan *types.Token
}

// newNFTMetadataUpdater creates a new instance of the NFT metadata updater service.
func newNFTMetadataUpdater(mgr *Manager) *nftMetadataUpdater {
	return &nftMetadataUpdater{
		mgr:          mgr,
		sigStop:      make(chan bool, 1),
		outTokens:    make(chan *types.Token, nftMetadataWorkerQueueCapacity),
		refreshQueue: make(chan *types.Token, nftMetadataRefreshSetSize),
	}
}

// name provides the name of the service.
func (mu *nftMetadataUpdater) name() string {
	return "nft metadata updater"
}

// init initializes the metadata scanner and registers it with the manager.
func (mu *nftMetadataUpdater) init() {
	// connect queues
	mu.inTokens = mu.mgr.logObserver.outNftTokens

	// add the updater to the manager
	mu.mgr.add(mu)
}

// close signals the block observer to terminate
func (mu *nftMetadataUpdater) close() {
	mu.sigStop <- true
}

// run processes metadata update requests from new NFT and also ensures
// updates on existing tokens' metadata as needed.
func (mu *nftMetadataUpdater) run() {
	// make the metadata refresh ticker
	refreshTick := time.NewTicker(nftMetadataRefreshTick)
	refreshPullTick := time.NewTicker(nftMetadataRefreshPullTick)

	defer func() {
		refreshTick.Stop()
		refreshPullTick.Stop()

		close(mu.outTokens)
		close(mu.refreshQueue)

		mu.mgr.closed(mu)
	}()

	var (
		ok  = true
		nft *types.Token
	)
	for ok {
		// pull next token
		nft, ok = mu.pullNext(refreshTick, refreshPullTick)

		// do we have a token to be pushed to the worker?
		if nft != nil {
			select {
			case <-mu.sigStop:
				return
			case mu.outTokens <- nft:
				log.Debugf("token %s/%s sent for processing", nft.Contract.String(), nft.TokenId.String())
				nft = nil
			}
		}
	}
}

// pullNext pulls a next available NFT from one of the pools for metadata processing
// It calls for refresh pool fill, if there is a time to do it.
func (mu *nftMetadataUpdater) pullNext(tiPool *time.Ticker, tiRefresh *time.Ticker) (*types.Token, bool) {
	var ok bool
	var nft *types.Token

	select {
	case <-mu.sigStop:
		return nil, false
	case nft, ok = <-mu.inTokens:
		if !ok {
			log.Noticef("input queue closed, terminating")
			return nil, false
		}
	case <-tiPool.C:
		mu.scheduleMetadataRefreshSet()
	case <-tiRefresh.C:
		// try to pull from the refresh queue instead
		select {
		case nft, ok = <-mu.refreshQueue:
			if !ok {
				log.Noticef("NFT metadata refresh queue closed, terminating")
				return nil, false
			}
		default:
		}
	}
	return nft, false
}

// scheduleMetadataRefreshSet pulls new refresh set from repository and pushes it into the scheduler.
func (mu *nftMetadataUpdater) scheduleMetadataRefreshSet() {
	defer func() {
		if p := recover(); p != nil {
			log.Errorf("could not collect metadata refresh set; ", p)
		}
	}()

	rs, err := repo.TokenMetadataRefreshSet(nftMetadataRefreshSetSize)
	if err != nil {
		log.Errorf("metadata refresh set not available; %s", err.Error())
		return
	}

	// log data
	log.Infof("loaded %d tokens in metadata refresh set", len(rs))

	// push the refresh set into the refresh queue
	// please note we don't wait for tokens to be stored
	for _, nft := range rs {
		select {
		case mu.refreshQueue <- nft:
			log.Debugf("scheduled metadata refresh on %s/%s", nft.Contract.String(), nft.TokenId.String())
		default:
		}
	}
}
