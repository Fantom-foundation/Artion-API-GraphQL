// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository"
	"sync"
)

// cfg holds reference to the server configuration.
var cfg *config.Config

// log represents the logger to be used by services to report state.
var log logger.Logger

// repo holds instance of the repository Proxy.
var repo *repository.Proxy

// manager holds singleton instance of the service manager.
var manager *Manager

// onceManager controls instance creation and ensures it has been created only once.
var onceManager sync.Once

// managerMux controls access to the manager singleton instance.
var managerMux sync.Mutex

// SetConfig sets the repository configuration to be used to establish
// and maintain external repository connections.
func SetConfig(c *config.Config) {
	cfg = c
}

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l.ModuleLogger("svc")
}

// Mgr provides access to the singleton instance of the SVC manager.
func Mgr() *Manager {
	managerMux.Lock()
	defer managerMux.Unlock()

	// make sure to instantiate the Repository only once
	onceManager.Do(func() {
		manager = newManager()
	})
	return manager
}

// Close terminates the singleton instance of the service manager.
func Close() {
	managerMux.Lock()
	defer managerMux.Unlock()

	if manager != nil {
		manager.Close()
	}
}
