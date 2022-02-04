package svc

import (
	"artion-api-graphql/internal/types"
	"time"
)

// 50 items per 15 seconds = 300_000 items per 25 hours
// 50 items per 10 seconds = 300_000 items per 16.6 hours
const (
	// likesRefreshTick is the tick used to pull likes refresh candidates.
	likesRefreshTick = 10 * time.Second

	// likesRefreshSetSize is the max size of likes refresh set pulled at once.
	likesRefreshSetSize = 50
)

// likesViewsUpdater represents a service responsible for updating token views/likes in local tokens collection.
type likesViewsUpdater struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the router
	sigStop chan bool

	// refreshQueue is the queue used to store NFT metadata update candidates.
	refreshQueue chan *types.Token
}

// newLikesViewsUpdater creates a new instance of the likes/views updater service.
func newLikesViewsUpdater(mgr *Manager) *likesViewsUpdater {
	return &likesViewsUpdater{
		mgr:          mgr,
		sigStop:      make(chan bool, 1),
		refreshQueue: make(chan *types.Token, likesRefreshSetSize),
	}
}

// name provides the name of the service.
func (mu *likesViewsUpdater) name() string {
	return "likes/views updater"
}

// init initializes the metadata scanner and registers it with the manager.
func (mu *likesViewsUpdater) init() {
	// add the updater to the manager
	mu.mgr.add(mu)
}

// close signals the block observer to terminate
func (mu *likesViewsUpdater) close() {
	mu.sigStop <- true
}

// run processes metadata update requests from new NFT and also ensures
// updates on existing tokens' metadata as needed.
func (mu *likesViewsUpdater) run() {
	// make the metadata refresh ticker
	refreshTick := time.NewTicker(likesRefreshTick)

	defer func() {
		refreshTick.Stop()
		close(mu.refreshQueue)
		mu.mgr.closed(mu)
	}()

	var token *types.Token
	for {
		select {
		case <-mu.sigStop:
			return
		case <-refreshTick.C:
			mu.scheduleLikesViewsRefreshSet()
		case token = <-mu.refreshQueue:
		}

		// do we have a token to be pushed to the worker?
		if token != nil {
			err := repo.TokenLikesViewsRefresh(token)
			if err != nil {
				log.Errorf("likes views refresh failed for %s/%s; %s", token.Contract.String(), token.TokenId.String(), err.Error())
				return
			}
			token = nil
		}
	}
}

// scheduleLikesViewsRefreshSet pulls new refresh set from repository and pushes it into the scheduler.
func (mu *likesViewsUpdater) scheduleLikesViewsRefreshSet() {
	defer func() {
		if p := recover(); p != nil {
			log.Errorf("could not collect likes refresh set; ", p)
		}
	}()

	rs, err := repo.TokenLikesViewsRefreshSet(likesRefreshSetSize)
	if err != nil {
		log.Errorf("likes refresh set not available; %s", err.Error())
		return
	}

	// log data
	log.Infof("loaded %d tokens in likes refresh set", len(rs))

	// push the refresh set into the refresh queue
	// please note we don't wait for tokens to be stored
	for _, nft := range rs {
		select {
		case mu.refreshQueue <- nft:
			log.Debugf("scheduled likes refresh on %s/%s", nft.Contract.String(), nft.TokenId.String())
		default:
		}
	}
}
