// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	eth "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	client "github.com/ethereum/go-ethereum/rpc"
	"sync"
)

const (
	// headerObserverCapacity represents the capacity of new headers' observer channel
	headerObserverCapacity = 5000
)

// config represents the configuration setup used by the repository
// to establish and maintain required connectivity to external services
// as needed.
var cfg *config.Config

// log represents the logger to be used by the repository.
var log logger.Logger

// opera represents the implementation of the Blockchain interface for Fantom Opera node.
type opera struct {
	rpc *client.Client
	ftm *ethclient.Client

	wg       *sync.WaitGroup
	sigClose chan bool
	headers  chan *eth.Header
}

// New provides a new instance of the RPC access point.
func New() Blockchain {
	con, err := connect()
	if err != nil {
		log.Criticalf("can not connect Opera node; %s", err.Error())
		return nil
	}

	op := &opera{
		rpc: con,
		ftm: ethclient.NewClient(con),

		wg:       new(sync.WaitGroup),
		sigClose: make(chan bool, 1),
		headers:  make(chan *eth.Header, headerObserverCapacity),
	}

	// start the observer
	op.wg.Add(1)
	go op.observe()

	log.Notice("blockchain connection is open")
	return op
}

// connects opens RPC connection to the Opera node.
func connect() (*client.Client, error) {
	c, err := client.Dial(cfg.Node.Url)
	if err != nil {
		log.Criticalf("can not connect blockchain node; %s", err.Error())
		return nil, err
	}

	log.Noticef("blockchain node connected at %s", cfg.Node.Url)
	return c, nil
}

// Close terminates the node connection.
func (o *opera) Close() {
	// stop block observer
	o.sigClose <- true
	o.wg.Wait()

	o.ftm.Close()
	log.Noticef("blockchain connection closed")
}

// HeaderObserver provides a channel receiving new observed blockchain headers.
func (o *opera) HeaderObserver() chan *eth.Header {
	return o.headers
}

// SetConfig sets the repository configuration to be used to establish
// and maintain external repository connections.
func SetConfig(c *config.Config) {
	cfg = c
}

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l.ModuleLogger("rpc")
}
