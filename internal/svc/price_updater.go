package svc

import (
	"artion-api-graphql/internal/types"
	"time"
)

const (
	// priceRefreshTick is the tick used to pull price refresh candidates.
	priceRefreshTick = 5 * time.Second

	// priceRefreshSetSize is the max size of price refresh set pulled at once.
	priceRefreshSetSize = 50
)

// priceUpdater represents a service responsible for updating token prices when the price validity expires
// (auction/offer is terminated by end-time, new listings starts to be valid by start-time etc.)
type priceUpdater struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the router
	sigStop chan bool

	// refreshQueue is the queue used to store token price update candidates.
	refreshQueue chan *types.Token
}

// newPriceUpdater creates a new instance of the price updater service.
func newPriceUpdater(mgr *Manager) *priceUpdater {
	return &priceUpdater{
		mgr:          mgr,
		sigStop:      make(chan bool, 1),
		refreshQueue: make(chan *types.Token, priceRefreshSetSize),
	}
}

// name provides the name of the service.
func (mu *priceUpdater) name() string {
	return "price updater"
}

// init initializes the metadata scanner and registers it with the manager.
func (mu *priceUpdater) init() {
	// add the updater to the manager
	mu.mgr.add(mu)
}

// close signals the block observer to terminate
func (mu *priceUpdater) close() {
	mu.sigStop <- true
}

// run processes the token prices update.
func (mu *priceUpdater) run() {
	// make the metadata refresh ticker
	refreshTick := time.NewTicker(priceRefreshTick)

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
			mu.schedulePriceRefreshSet()
		case token = <-mu.refreshQueue:
		}

		// do we have a token to be pushed to the worker?
		if token != nil {
			err := repo.TokenPriceRefresh(token)
			if err != nil {
				log.Errorf("price refresh failed for %s/%s; %s", token.Contract.String(), token.TokenId.String(), err.Error())
				return
			}
			token = nil
		}
	}
}

// schedulePriceRefreshSet pulls new refresh set from repository and pushes it into the scheduler.
func (mu *priceUpdater) schedulePriceRefreshSet() {
	defer func() {
		if p := recover(); p != nil {
			log.Errorf("could not collect price refresh set; ", p)
		}
	}()

	rs, err := repo.TokenPriceRefreshSet(priceRefreshSetSize)
	if err != nil {
		log.Errorf("price refresh set not available; %s", err.Error())
		return
	}

	// log data
	log.Infof("price refresh set - loaded %d tokens", len(rs))

	// push the refresh set into the refresh queue
	// please note we don't wait for tokens to be stored
	for _, nft := range rs {
		select {
		case mu.refreshQueue <- nft:
			log.Debugf("scheduled price refresh on %s/%s", nft.Contract.String(), nft.TokenId.String())
		default:
		}
	}
}
