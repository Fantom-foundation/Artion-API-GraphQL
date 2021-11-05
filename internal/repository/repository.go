// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository/cache"
	"artion-api-graphql/internal/repository/db"
	"artion-api-graphql/internal/repository/email"
	"artion-api-graphql/internal/repository/pinner"
	"artion-api-graphql/internal/repository/rpc"
	"artion-api-graphql/internal/repository/uri"
	"artion-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/singleflight"
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
	rpc       *rpc.Opera
	uri       *uri.Downloader
	pinner    *pinner.Pinner
	db        *db.MongoDbBridge
	shared    *db.SharedMongoDbBridge
	cache     *cache.MemCache
	callGroup *singleflight.Group

	observedContracts []common.Address
	nftTypes          map[common.Address]string

	// notificationQueue processing channel
	notificationQueue  chan types.Notification
	newCollectionQueue chan common.Address
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

	// in-memory cache
	cache.SetLogger(log)
	cache.SetConfig(cfg)

	// emailing service
	email.SetLogger(log)
	email.SetConfig(cfg)

	// pinner
	pinner.SetLogger(log)
}

// newProxy creates new instance of Proxy, implementing the Repository interface.
func newProxy() *Proxy {
	// pass environment to the sub-modules
	passEnvironment()

	// make Proxy instance
	p := Proxy{
		rpc:       rpc.New(),
		uri:       uri.New(cfg),
		pinner:    pinner.New(cfg),
		db:        db.New(),
		shared:    db.NewShared(),
		cache:     cache.New(),
		callGroup: new(singleflight.Group),

		// NFT contracts info
		nftTypes: make(map[common.Address]string, 0),

		notificationQueue:  make(chan types.Notification, notificationQueueCapacity),
		newCollectionQueue: make(chan common.Address, addCollectionQueueCapacity),
	}

	if p.db == nil || p.rpc == nil || p.cache == nil {
		log.Panicf("repository init failed")
		return nil
	}

	// register contracts to the repository
	p.registerContracts()

	log.Notice("repository ready")
	return &p
}

// Close terminates repository connections.
func (p *Proxy) Close() {
	close(p.notificationQueue)

	if p.rpc != nil {
		p.rpc.Close()
	}

	if p.db != nil {
		p.db.Close()
	}
}

// registerContracts will pass contract addresses to the RPC provider.
func (p *Proxy) registerContracts() {
	var contractTypes = []string{"auction", "market", "rng"}

	for _, ct := range contractTypes {
		err := p.rpc.RegisterContract(ct, p.ObservedContractAddressByType(ct))
		if err != nil {
			log.Panicf("mandatory contract %s not available", ct)
		}
	}

	// load list of observed contract addresses
	log.Notice("loading observed contracts")
	p.observedContracts = p.ObservedContractsAddressList()
	p.loadObservedCollections()
}
