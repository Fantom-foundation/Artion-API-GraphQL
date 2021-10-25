// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"strings"
	"time"
)

// nftMetadataWorker represents a service responsible for processing NFT token metadata
// update queue from the metadata updater service.
type nftMetadataWorker struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the router
	sigStop chan bool

	// inTokens represents the channel receiving tokens to be updated.
	inTokens chan *types.Token
}

// newNFTMetadataUpdater creates a new instance of the NFT metadata worker service.
func newNFTMetadataWorker(mgr *Manager) *nftMetadataWorker {
	return &nftMetadataWorker{
		mgr:     mgr,
		sigStop: make(chan bool, 1),
	}
}

// name provides the name of the service.
func (mw *nftMetadataWorker) name() string {
	return "nft metadata worker"
}

// init initializes the block scanner and registers it with the manager.
func (mw *nftMetadataWorker) init() {
	// connect queues
	mw.inTokens = mw.mgr.nftMetaUpdater.outTokens

	// add the updater to the manager
	mw.mgr.add(mw)
}

// close signals the block observer to terminate
func (mw *nftMetadataWorker) close() {
	mw.sigStop <- true
}

// run processes incoming NFT metadata update requests from the
// incoming queue.
func (mw *nftMetadataWorker) run() {
	defer func() {
		mw.mgr.closed(mw)
	}()

	for {
		// pull next token
		select {
		case <-mw.sigStop:
			return
		case tok, ok := <-mw.inTokens:
			if !ok {
				return
			}

			// process the token metadata update
			if err := mw.update(tok); err != nil {
				log.Errorf("NFT update failed; %s", err.Error())
			}
		}
	}
}

// update the given NFT metadata from external metadata source.
func (mw *nftMetadataWorker) update(tok *types.Token) error {
	// get metadata
	if tok.Uri == "" {
		log.Infof("token %s/%s metadata URI not available", tok.Contract.String(), tok.TokenId.String())
		return nil
	}

	// get the metadata
	log.Debugf("loading metadata for %s/%s", tok.Contract.String(), tok.TokenId.String())
	md, err := repo.GetTokenJsonMetadata(tok.Uri)
	if err != nil {
		log.Errorf("NFT metadata [%s] failed on %s/%s; %s", tok.Uri, tok.Contract.String(), tok.TokenId.String(), err.Error())
		handleTokenMetaUpdateFailure(tok)
		return err
	}

	// get image type
	if md.Image != nil {
		img, err := repo.GetImage(*md.Image)
		if err != nil {
			log.Errorf("NFT image [%s] failed on %s/%s; %s", md.Image, tok.Contract.String(), tok.TokenId.String(), err.Error())
			handleTokenMetaUpdateFailure(tok)
			return err
		}
		log.Warningf("NFT image [%s] has type %d", *md.Image, img.Type)
		tok.ImageType = img.Type
	}

	// update the data
	tok.ScheduleMetaUpdateOnSuccess()
	tok.Name = strings.TrimSpace(md.Name)
	tok.Description = strings.TrimSpace(md.Description)
	if md.Image != nil {
		tok.ImageURI = strings.TrimSpace(*md.Image)
	}

	updateTokenCategoriesFromCollection(tok)

	// does this token make a sense?
	tok.IsActive = tok.Name != "" || tok.Description != "" || tok.ImageURI != ""

	// update the token in persistent storage
	if err := repo.UpdateTokenMetadata(tok); err != nil {
		log.Errorf("failed metadata update on %s/%s; %s", tok.Contract.String(), tok.TokenId.String(), err.Error())
		return err
	}
	log.Infof("NFT %s/%s metadata updated [%s]", tok.Contract.String(), tok.TokenId.String(), tok.Name)
	return nil
}

func handleTokenMetaUpdateFailure(tok *types.Token) {
	tok.ScheduleMetaUpdateOnFailure()
	if e := repo.UpdateTokenMetadataRefreshSchedule(tok); e != nil {
		log.Errorf("token schedule update failed;", e.Error())
	}

	log.Infof("next update #%d of %s/%s at %s",
		tok.MetaFailures, tok.Contract.String(), tok.TokenId.String(),
		time.Time(tok.MetaUpdate).Format(time.Stamp))
}

func updateTokenCategoriesFromCollection(tok *types.Token) {
	lc, err := repo.GetLegacyCollection(tok.Contract)
	if err != nil {
		log.Errorf("NFT legacy collection not available for %s; %s", tok.Contract.String(), err.Error())
	}

	if lc != nil {
		tok.Categories, err = lc.CategoriesAsInts()
		if err != nil {
			log.Errorf("failed to decode categories for token contract %s; %s", tok.Contract.String(), err.Error())
		}
	}
}
