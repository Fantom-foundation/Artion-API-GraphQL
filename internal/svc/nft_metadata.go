// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"time"
)

// nftMetadataUpdaterQueueCapacity is the capacity of NFT metadata updater queue.
const nftMetadataUpdaterQueueCapacity = 5000

// nftMetadataUpdater represents a service responsible for updating NFT token metadata
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
	workQueue chan *types.Token
}

// newNFTMetadataUpdater creates a new instance of the NFT metadata updater service.
func newNFTMetadataUpdater(mgr *Manager) *nftMetadataUpdater {
	return &nftMetadataUpdater{
		mgr:       mgr,
		sigStop:   make(chan bool, 1),
		workQueue: make(chan *types.Token, nftMetadataUpdaterQueueCapacity),
	}
}

// name provides the name of the service.
func (mu *nftMetadataUpdater) name() string {
	return "nft metadata updater"
}

// init initializes the block scanner and registers it with the manager.
func (mu *nftMetadataUpdater) init() {
	// connect queues
	mu.inTokens = mu.mgr.logObserver.outNftTokens

	// run worker
	mu.mgr.wg.Add(1)
	go mu.worker()

	// add the updater to the manager
	mu.mgr.add(mu)
}

// run scans past blocks one by one until it reaches top
// after the top is reached, it idles and checks the head state to make sure
// the API server keeps up with the most recent block
func (mu *nftMetadataUpdater) run() {
	defer func() {
		close(mu.workQueue)
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
			select {
			case <-mu.sigStop:
				return
			case mu.workQueue <- tok:
			}
		}
	}
}

// worker processes all the tokens from the worker queue updating persistent storage metadata
// with the data loaded from metadata JSON file.
func (mu *nftMetadataUpdater) worker() {
	defer func() {
		log.Noticef("NFT metadata update worker closed")
		mu.mgr.wg.Done()
	}()

	for {
		// pull next token
		tok, ok := <-mu.workQueue
		if !ok {
			return
		}

		// process the token
		if err := mu.update(tok); err != nil {
			log.Errorf("NFT update failed; %s", err.Error())
		}
	}
}

// update the given NFT metadata from external metadata source.
func (mu *nftMetadataUpdater) update(tok *types.Token) error {
	// get metadata
	if tok.Uri == "" {
		log.Infof("token %d on contract %s URI not available", tok.TokenId.String(), tok.Nft.String())
		return nil
	}

	// get the metadata
	md, err := repo.GetTokenJsonMetadata(tok.Uri)
	if err != nil {
		log.Errorf("NFT metadata failed for token %d on contract %s; %s", tok.TokenId.String(), tok.Nft.String(), err.Error())
		return err
	}

	// update the data
	tok.Name = md.Name
	tok.Description = md.Description
	if md.Image != nil {
		tok.ImageURI = *md.Image
	}
	tok.MetadataAge = types.Time(time.Now())

	// update the token in persistent storage
	if err := repo.StoreToken(tok); err != nil {
		log.Errorf("failed NFT metadata update for token %d on contract %s; %s", tok.TokenId.String(), tok.Nft.String(), err.Error())
		return err
	}
	return nil
}

// close signals the block observer to terminate
func (mu *nftMetadataUpdater) close() {
	mu.sigStop <- true
}
