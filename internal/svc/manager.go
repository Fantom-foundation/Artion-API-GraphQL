// Package svc implements monitoring and scanning services of the API server.
package svc

import "sync"

// Manager represents the manager controlling services lifetime.
type Manager struct {
	wg  *sync.WaitGroup
	svc []service
}

// newManager creates a new instance of the svc Manager.
func newManager() *Manager {
	// prep the manager
	mgr := Manager{
		wg:  new(sync.WaitGroup),
		svc: make([]service, 0),
	}

	// init & start services
	mgr.load()

	log.Notice("services are running")
	return &mgr
}

// load initializes the services in the correct order.
func (mgr *Manager) load() {

}

// add managed service instance to the Manager and run it.
func (mgr *Manager) add(s service) {
	// keep track of running services
	mgr.svc = append(mgr.svc, s)

	// run the service
	mgr.wg.Add(1)
	go s.run()
}

// Close terminates the service manager
// and all the managed services along with it.
func (mgr *Manager) Close() {
	log.Notice("closing services")
	for _, s := range mgr.svc {
		s.close()
	}

	mgr.wg.Wait()
	log.Notice("services closed")
}
