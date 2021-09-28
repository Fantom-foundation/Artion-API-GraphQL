// Package svc implements monitoring and scanning services of the API server.
package svc

import "sync"

// Manager represents the manager controlling services lifetime.
type Manager struct {
	wg  *sync.WaitGroup
	svc []service

	// managed services
	blkRouter   *blkRouter
	blkScanner  *blkScanner
	blkObserver *blkObserver
	logObserver *logObserver
}

// newManager creates a new instance of the svc Manager.
func newManager() *Manager {
	// prep the manager
	mgr := Manager{
		wg:  new(sync.WaitGroup),
		svc: make([]service, 0),
	}

	// init & start services
	mgr.blkRouter = newBlkRouter(&mgr)
	mgr.blkScanner = newBlkScanner(&mgr)
	mgr.blkObserver = newBlkObserver(&mgr)
	mgr.logObserver = newLogObserver(&mgr)
	mgr.init()

	log.Notice("all services are running")
	return &mgr
}

// init initializes the services in the correct order.
func (mgr *Manager) init() {
	mgr.blkRouter.init()
	mgr.blkScanner.init()
	mgr.blkObserver.init()
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
