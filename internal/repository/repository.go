// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository/db"
	"artion-api-graphql/internal/repository/rpc"
	"fmt"
	"sync"
)

// config represents the configuration setup used by the repository
// to establish and maintain required connectivity to external services
// as needed.
var cfg *config.Config

// log represents the logger to be used by the repository.
var log logger.Logger

// instance is the singleton of the Proxy, implementing Repository interface.
var instance *Proxy

// oneInstance is the sync guarding Repository singleton creation.
var oneInstance sync.Once

// instanceMux controls access to the repository instance
var instanceMux sync.Mutex

// Proxy is the implementation of the Repository interface
type Proxy struct {
	rpc *rpc.Opera
	db  *db.MongoDbBridge
}

// R provides access to the singleton instance of the Repository.
func R() *Proxy {
	instanceMux.Lock()
	defer instanceMux.Unlock()

	// make sure to instantiate the Repository only once
	oneInstance.Do(func() {
		instance = newProxy()
	})
	return instance
}

// Close will terminate the singleton instance of the repository, if started already.
func Close() {
	instanceMux.Lock()
	defer instanceMux.Unlock()

	if instance != nil {
		instance.Close()
	}
}

// SetConfig sets the repository configuration to be used to establish
// and maintain external repository connections.
func SetConfig(c *config.Config) {
	cfg = c
}

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l.ModuleLogger("repo")
}

// passEnvironment provides configuration and logger to the repository sub-modules.
func passEnvironment() {
	if cfg == nil {
		panic(fmt.Errorf("missing configuration"))
	}
	if log == nil {
		panic(fmt.Errorf("missing logger"))
	}

	// RPC module
	rpc.SetLogger(log)
	rpc.SetConfig(cfg)

	// persistent storage module
	db.SetLogger(log)
	db.SetConfig(cfg)
}

// newProxy creates new instance of Proxy, implementing the Repository interface.
func newProxy() *Proxy {
	// pass environment to the sub-modules
	passEnvironment()

	// make Proxy instance
	p := Proxy{
		db:  db.New(),
		rpc: rpc.New(),
	}

	if p.db == nil || p.rpc == nil {
		log.Panicf("repository init failed")
		return nil
	}

	log.Notice("repository ready")
	return &p
}

// Close terminates repository connections.
func (p *Proxy) Close() {
	if p.rpc != nil {
		p.rpc.Close()
	}
	if p.db != nil {
		p.db.Close()
	}
}
