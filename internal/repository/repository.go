// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository/db"
	"artion-api-graphql/internal/repository/rpc"
	"sync"
)

// config represents the configuration setup used by the repository
// to establish and maintain required connectivity to external services
// as needed.
var cfg *config.Config

// log represents the logger to be used by the repository.
var log logger.Logger

// instance is the singleton of the Proxy interface.
var instance *proxy

// oneInstance is the sync guarding Proxy singleton creation.
var oneInstance sync.Once

// proxy is the implementation of the Proxy interface
type proxy struct {
	rpc rpc.Blockchain
	db    *db.MongoDbBridge
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

// Do provide access to singleton instance of the repository Proxy.
func Do() Proxy {
	oneInstance.Do(func() {
		instance = &proxy{
			rpc: rpc.New(cfg),
		}
	})
	return instance
}
