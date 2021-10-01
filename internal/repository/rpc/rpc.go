// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"bytes"
	"embed"
	"github.com/ethereum/go-ethereum/accounts/abi"
	eth "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	client "github.com/ethereum/go-ethereum/rpc"
	"sync"
)

//go:embed contracts/abi/*.json
var abiFiles embed.FS

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

// Opera represents the implementation of the Blockchain interface for Fantom Opera node.
type Opera struct {
	rpc *client.Client
	ftm *ethclient.Client

	wg       *sync.WaitGroup
	sigClose chan bool
	headers  chan *eth.Header

	abiFantom721 abi.ABI
}

// New provides a new instance of the RPC access point.
func New() *Opera {
	con, err := connect()
	if err != nil {
		log.Criticalf("can not connect Opera node; %s", err.Error())
		return nil
	}

	op := &Opera{
		rpc: con,
		ftm: ethclient.NewClient(con),

		wg:       new(sync.WaitGroup),
		sigClose: make(chan bool, 1),
		headers:  make(chan *eth.Header, headerObserverCapacity),
	}

	// load and parse ABIs
	if err := loadABI(op); err != nil {
		log.Criticalf("can not parse ABI files; %s", err.Error())
		return nil
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

// loadABI tries to load and parse expected ABI for contracts we need.
func loadABI(o *Opera) (err error) {
	// FantomNFTTradable
	data, err := abiFiles.ReadFile("contracts/abi/FantomNFTTradable.json")
	if err != nil {
		return err
	}

	// parse ABI
	o.abiFantom721, err = abi.JSON(bytes.NewReader(data))
	if err != nil {
		return err
	}

	return nil
}

// Close terminates the node connection.
func (o *Opera) Close() {
	// stop block observer
	o.sigClose <- true
	o.wg.Wait()

	o.ftm.Close()
	log.Noticef("blockchain connection closed")
}

// NewHeaders provides a channel receiving new blockchain headers.
func (o *Opera) NewHeaders() chan *eth.Header {
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
