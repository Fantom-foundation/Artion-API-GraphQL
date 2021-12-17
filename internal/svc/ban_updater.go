package svc

import (
	"math/big"
	"time"
)

const (
	// banRefreshTick is the tick used to pull ban status refresh candidates.
	banRefreshTick = 5 * time.Minute

	banRefreshOverlap = 30 * time.Second

	banNftRefreshSetSize = 50

	banCollectionRefreshSetSize = 50
)

// banUpdater represents a service responsible for updating tokens ban status
// (ban status is obtained from collection.isAppropriate and from presence in bannednfts in shared database)
type banUpdater struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the router
	sigStop chan bool
}

func newBanUpdater(mgr *Manager) *banUpdater {
	return &banUpdater{
		mgr:     mgr,
		sigStop: make(chan bool, 1),
	}
}

// name provides the name of the service.
func (mu *banUpdater) name() string {
	return "ban updater"
}

// init initializes the metadata scanner and registers it with the manager.
func (mu *banUpdater) init() {
	// add the updater to the manager
	mu.mgr.add(mu)
}

// close signals the block observer to terminate
func (mu *banUpdater) close() {
	mu.sigStop <- true
}

// run processes the token prices update.
func (mu *banUpdater) run() {
	// make the metadata refresh ticker
	refreshTick := time.NewTicker(banRefreshTick)

	defer func() {
		refreshTick.Stop()
		mu.mgr.closed(mu)
	}()

	for {
		select {
		case <-mu.sigStop:
			return
		case <-refreshTick.C:
		}
		mu.runUpdates()
	}
}

func (mu *banUpdater) runUpdates() {
	defer func() {
		if p := recover(); p != nil {
			log.Errorf("updating ban status of tokens failed; ", p)
		}
	}()

	log.Info("updating ban status of tokens started")

	newUpdateTime := time.Now().Add(-banRefreshOverlap)
	var after time.Time
	lastUpdate, err := repo.LastBanUpdate()
	if lastUpdate == nil {
		after = time.Unix(0,0)
	} else {
		after = *lastUpdate
	}

	// set collection-level banned
	cols, err := repo.ListCollectionsWithAppropriateUpdate(after, banCollectionRefreshSetSize)
	if err != nil {
		log.Errorf("collections with IsAppropriate changes not available; %s", err)
		return
	}
	for _, col := range cols {
		err = repo.TokenMarkCollectionBanned(&col.Address, ! col.IsAppropriate)
		if err != nil {
			log.Errorf("setting tokens as collection-banned failed; %s", err)
		}
	}

	// set token-level banned
	tokens, err := repo.ListBannedNftsAfter(after, banNftRefreshSetSize)
	if err != nil {
		log.Errorf("new banned nfts not available; %s", err)
		return
	}
	for _, token := range tokens {
		err = repo.TokenMarkBanned(&token.Contract, (*big.Int)(&token.TokenId), token.IsBanned)
		if err != nil {
			log.Errorf("setting token as banned failed; %s", err)
		}
	}

	// save the last update time
	repo.UpdateLastBanUpdate(newUpdateTime)

	log.Noticef("updating ban status of tokens done (updated %d collections, %d tokens)", len(cols), len(tokens))
}
