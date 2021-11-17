// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"strings"
	"sync"
	"time"
)

// nftMetadataWorkerThreads is the number of threads working on NFT metadata updates.
const nftMetadataWorkerThreads = 15

// nftMetadataWorker represents a service responsible for processing NFT token metadata
// update queue from the metadata updater service.
type nftMetadataWorker struct {
	// mgr represents the Manager instance
	mgr *Manager

	// workers is the list of work threads
	workers []*nftMetadataWorkerThread

	// wg is the wait group keeping eye on worker threads
	wg *sync.WaitGroup
}

// nftMetadataWorkerThread represents a thread working on NFT metadata updates.
type nftMetadataWorkerThread struct {
	// inTokens represents the channel receiving tokens to be updated.
	inTokens chan *types.Token

	// sigStop represents the signal for closing the router
	sigStop chan bool

	// wg is the wait group keeping eye on worker threads
	wg *sync.WaitGroup
}

// newNFTMetadataUpdater creates a new instance of the NFT metadata worker service.
func newNFTMetadataWorker(mgr *Manager) *nftMetadataWorker {
	return &nftMetadataWorker{
		mgr:     mgr,
		workers: make([]*nftMetadataWorkerThread, nftMetadataWorkerThreads),
		wg:      new(sync.WaitGroup),
	}
}

// name provides the name of the service.
func (mw *nftMetadataWorker) name() string {
	return "nft metadata worker"
}

// init initializes the block scanner and registers it with the manager.
func (mw *nftMetadataWorker) init() {
	// prep the threads
	for i := 0; i < nftMetadataWorkerThreads; i++ {
		mw.workers[i] = &nftMetadataWorkerThread{
			inTokens: mw.mgr.nftMetaUpdater.outTokens,
			sigStop:  make(chan bool, 1),
			wg:       mw.wg,
		}
	}

	// add the updater to the manager
	mw.mgr.add(mw)
}

// close signals the block observer to terminate
func (mw *nftMetadataWorker) close() {
	for i := 0; i < nftMetadataWorkerThreads; i++ {
		mw.workers[i].sigStop <- true
	}
}

// run start the work threads
func (mw *nftMetadataWorker) run() {
	for i := 0; i < nftMetadataWorkerThreads; i++ {
		mw.wg.Add(1)
		go mw.workers[i].run()
	}

	mw.wg.Wait()
	mw.mgr.closed(mw)
}

// run processes incoming NFT metadata update requests from the
// incoming queue.
func (mwt *nftMetadataWorkerThread) run() {
	defer func() {
		mwt.wg.Done()
	}()

	for {
		// pull next token
		select {
		case <-mwt.sigStop:
			return
		case tok, ok := <-mwt.inTokens:
			if !ok {
				return
			}

			// process the token metadata update
			if err := mwt.update(tok); err != nil {
				log.Errorf("NFT update failed; %s", err.Error())
			}
		}
	}
}

// update the given NFT metadata from external metadata source.
func (mwt *nftMetadataWorkerThread) update(tok *types.Token) error {
	// get metadata
	if tok.Uri == "" {
		log.Infof("token %s/%s metadata URI not available", tok.Contract.String(), tok.TokenId.String())
		mwt.tryLegacyUpdate(tok)
		return nil
	}

	// get the metadata
	log.Debugf("loading metadata for %s/%s", tok.Contract.String(), tok.TokenId.String())
	md, err := repo.GetTokenJsonMetadata(tok.Uri)
	if err != nil {
		log.Errorf("NFT metadata [%s] failed on %s/%s; %s", tok.Uri, tok.Contract.String(), tok.TokenId.String(), err.Error())
		mwt.tryLegacyUpdate(tok)

		handleTokenMetaUpdateFailure(tok)
		return err
	}

	// get image type (skip if the image URI has not changed)
	if md.Image != nil && *md.Image != tok.ImageURI {
		img, err := repo.GetImage(*md.Image)
		if err != nil {
			log.Errorf("NFT image [%s] failed on %s/%s; %s", md.Image, tok.Contract.String(), tok.TokenId.String(), err.Error())
			mwt.tryLegacyUpdate(tok)

			handleTokenMetaUpdateFailure(tok)
			return err
		}
		log.Debugf("NFT image [%s] has type %d", *md.Image, img.Type)
		tok.ImageType = img.Type
	}

	// update the data
	tok.ScheduleMetaUpdateOnSuccess()

	tok.Name = strings.TrimSpace(md.Name)
	if md.Description != nil {
		tok.Description = strings.TrimSpace(*md.Description)
	}
	if md.Image != nil {
		tok.ImageURI = strings.TrimSpace(*md.Image)
	}
	if md.Properties.Symbol != nil {
		tok.Symbol = *md.Properties.Symbol
	}
	if md.Properties.IpRights != nil {
		tok.IpRights = *md.Properties.IpRights
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

// tryLegacyUpdate tries to load critical details of the given NFT metadata from legacy source.
func (mwt *nftMetadataWorkerThread) tryLegacyUpdate(tok *types.Token) {
	t, err := repo.ExtendLegacyToken(tok)
	if err != nil {
		log.Errorf("could not extend %s / %s from legacy; %s", tok.Contract.String(), tok.TokenId.String(), err.Error())
		return
	}

	if t == nil {
		log.Errorf("not found legacy extension of %s / %s", tok.Contract.String(), tok.TokenId.String())
		return
	}

	// does this token make a sense?
	tok.IsActive = tok.Name != "" || tok.Description != "" || tok.ImageURI != ""
	log.Warningf("token %s / %s updated from legacy DB [%s]", tok.Contract.String(), tok.TokenId.String(), tok.Name)

	// update the token in persistent storage
	if err := repo.UpdateTokenMetadata(tok); err != nil {
		log.Errorf("failed metadata update on %s/%s; %s", tok.Contract.String(), tok.TokenId.String(), err.Error())
		return
	}
	log.Infof("NFT %s/%s metadata updated [%s]", tok.Contract.String(), tok.TokenId.String(), tok.Name)
}

// handleTokenMetaUpdateFailure updates the token Metadata update schedule on failure.
func handleTokenMetaUpdateFailure(tok *types.Token) {
	tok.ScheduleMetaUpdateOnFailure()
	if e := repo.UpdateTokenMetadataRefreshSchedule(tok); e != nil {
		log.Errorf("token schedule update failed;", e.Error())
	}

	log.Infof("next update #%d of %s/%s at %s",
		tok.MetaFailures, tok.Contract.String(), tok.TokenId.String(),
		time.Time(tok.MetaUpdate).Format(time.Stamp))
}

// updateTokenCategoriesFromCollection loads list of categories of a token from legacy collection setup.
func updateTokenCategoriesFromCollection(tok *types.Token) {
	lc, err := repo.GetLegacyCollection(tok.Contract)
	if err != nil {
		log.Errorf("NFT legacy collection not available for %s; %s", tok.Contract.String(), err.Error())
	}

	if lc != nil {
		tok.Categories, err = lc.CategoriesAsInt()
		if err != nil {
			log.Errorf("failed to decode categories for token contract %s; %s", tok.Contract.String(), err.Error())
		}
	}
}
