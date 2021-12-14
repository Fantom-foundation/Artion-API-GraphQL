// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/repository"
	"sync"
)

// Manager represents the manager controlling services lifetime.
type Manager struct {
	wg  *sync.WaitGroup
	svc []service

	// managed services
	blkScanner          *blkScanner
	blkObserver         *blkObserver
	logObserver         *logObserver
	nftMetaUpdater      *nftMetadataUpdater
	nftMetaWorker       *nftMetadataWorker
	notifyProcessor     *notificationProcessor
	collectionValidator *nftCollectionValidator
	priceUpdater        *priceUpdater
	likesViewsUpdater   *likesViewsUpdater
	banUpdater          *banUpdater
}

// newManager creates a new instance of the svc Manager.
func newManager() *Manager {
	// prep the manager
	mgr := Manager{
		wg:  new(sync.WaitGroup),
		svc: make([]service, 0),
	}

	// get the repository instance
	repo = repository.R()

	// make services
	mgr.blkScanner = newBlkScanner(&mgr)
	mgr.blkObserver = newBlkObserver(&mgr)
	mgr.logObserver = newLogObserver(&mgr)
	mgr.nftMetaUpdater = newNFTMetadataUpdater(&mgr)
	mgr.nftMetaWorker = newNFTMetadataWorker(&mgr)
	mgr.notifyProcessor = newNotificationProcessor(&mgr)
	mgr.collectionValidator = newNFTCollectionValidator(&mgr)
	mgr.priceUpdater = newPriceUpdater(&mgr)
	mgr.likesViewsUpdater = newLikesViewsUpdater(&mgr)
	mgr.banUpdater = newBanUpdater(&mgr)

	// init and run
	mgr.init()
	log.Notice("all services are running")
	return &mgr
}

// init initializes the services in the correct order.
func (mgr *Manager) init() {
	mgr.blkScanner.init()
	mgr.blkObserver.init()
	mgr.logObserver.init()
	mgr.nftMetaWorker.init()
	mgr.nftMetaUpdater.init()
	mgr.notifyProcessor.init()
	mgr.collectionValidator.init()
	mgr.priceUpdater.init()
	mgr.likesViewsUpdater.init()
	mgr.banUpdater.init()
}

// add managed service instance to the Manager and run it.
func (mgr *Manager) add(s service) {
	// keep track of running services
	mgr.svc = append(mgr.svc, s)

	// run the service
	mgr.wg.Add(1)
	go s.run()
	log.Noticef("service %s started", s.name())
}

// closed signals the manager a service terminated.
func (mgr *Manager) closed(s service) {
	mgr.wg.Done()
	log.Noticef("service %s stopped", s.name())
}

// Close terminates the service manager
// and all the managed services along with it.
func (mgr *Manager) Close() {
	log.Notice("services are being terminated")

	for _, s := range mgr.svc {
		log.Noticef("closing %s", s.name())
		s.close()
	}

	mgr.wg.Wait()
	log.Notice("services closed")
}
