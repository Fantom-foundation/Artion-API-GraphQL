// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
)

// nftMetadataUpdaterQueueCapacity is the capacity of NFT metadata updater queue.
const nftMetadataUpdaterQueueCapacity = 1000

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

	// workQueue represents the channel receiving both new NFTs and periodic updates for processing.
	outTokens chan *types.Token
}

// newNFTMetadataUpdater creates a new instance of the NFT metadata updater service.
func newNFTMetadataUpdater(mgr *Manager) *nftMetadataUpdater {
	return &nftMetadataUpdater{
		mgr:       mgr,
		sigStop:   make(chan bool, 1),
		outTokens: make(chan *types.Token, nftMetadataUpdaterQueueCapacity),
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

// run processes metadata update requests from new NFT and also ensures
// updates on existing tokens' metadata as needed.
func (mu *nftMetadataUpdater) run() {
	defer func() {
		close(mu.outTokens)
		mu.mgr.closed(mu)
	}()

	for {
		select {
		case <-mu.sigStop:
			return
		case tok, ok := <-mu.inTokens:
			if !ok {
				return
			}

			log.Debugf("received token %s / #%d for processing", tok.Nft.String(), tok.TokenId.ToInt().Uint64())
			select {
			case <-mu.sigStop:
				return
			case mu.outTokens <- tok:
			}
		}
	}
}

// close signals the block observer to terminate
func (mu *nftMetadataUpdater) close() {
	mu.sigStop <- true
}
