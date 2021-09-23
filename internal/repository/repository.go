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

// repo represents an instance of the Repository manager.
var repo Repository

// onceRepo is the sync object used to make sure the Repository
// is instantiated only once on the first demand.
var onceRepo sync.Once

// config represents the configuration setup used by the repository
// to establish and maintain required connectivity to external services
// as needed.
var cfg *config.Config

// log represents the logger to be used by the repository.
var log logger.Logger

// instance is the singleton of the Repository interface.
var instance *proxy

// oneInstance is the sync guarding Repository singleton creation.
var oneInstance sync.Once

// R provides access to the singleton instance of the Repository.
func R() Repository {
	// make sure to instantiate the Repository only once
	onceRepo.Do(func() {
		repo = newRepository()
	})
	return repo
}

// proxy is the implementation of the Repository interface
type proxy struct {
	rpc rpc.Blockchain
	db    *db.MongoDbBridge
	log   logger.Logger
	cfg *config.Config
}

// SetConfig sets the repository configuration to be used to establish
// and maintain external repository connections.
func SetConfig(c *config.Config) {
	cfg = c
}

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l
}

// newRepository creates new instance of Repository implementation, namely proxy structure.
func newRepository() Repository {
	if cfg == nil {
		panic(fmt.Errorf("missing configuration"))
	}
	if log == nil {
		panic(fmt.Errorf("missing logger"))
	}

	// create connections
	dbBridge, err := connect(cfg, log)
	if err != nil {
		log.Fatal("repository init failed")
		return nil
	}

	// construct the proxy instance
	p := proxy{
		db:    dbBridge,
		log:   log,
		cfg:   cfg,
	}

	// return the proxy
	return &p
}

// connect opens connections to the external sources we need.
func connect(cfg *config.Config, log logger.Logger) (*db.MongoDbBridge, error) {
	// create new database connection bridge
	dbBridge, err := db.New(cfg, log)
	if err != nil {
		log.Criticalf("can not connect backend persistent storage, %s", err.Error())
		return nil, err
	}

	return dbBridge, nil
}