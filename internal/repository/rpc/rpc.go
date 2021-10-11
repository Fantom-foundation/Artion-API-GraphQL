// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository/rpc/contracts"
	"bytes"
	"embed"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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

// contractsMap represents a map of contract type to address as provided
// by the repository initializer.
var contractsMap = map[string]common.Address{}

// Opera represents the implementation of the Blockchain interface for Fantom Opera node.
type Opera struct {
	// basics of the connection
	rpc *client.Client
	ftm *ethclient.Client

	// sync tools
	wg       *sync.WaitGroup
	sigClose chan bool

	// captured header queue
	headers chan *eth.Header

	// decode ABI structures
	abiFantom721   *abi.ABI
	abiFantom1155  *abi.ABI
	abiMarketplace *abi.ABI

	// contracts
	auctionContract *contracts.FantomAuction
}

// RegisterContract adds a new contract address to the RPC provider.
func (o *Opera) RegisterContract(ct string, addr *common.Address) (err error) {
	// address provided?
	if nil == addr {
		return fmt.Errorf("empty address on %s", ct)
	}

	// load the contract instance
	switch ct {
	case "auction":
		o.auctionContract, err = contracts.NewFantomAuction(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded %s contract at %s", ct, addr.String())
		}
	default:
		err = fmt.Errorf("unknown contract type %s", ct)
	}

	return err
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
	o.abiFantom721, err = loadABIFile("contracts/abi/FantomNFTTradable.json")
	if err != nil {
		return err
	}

	o.abiFantom1155, err = loadABIFile("contracts/abi/FantomArtTradable.json")
	if err != nil {
		return err
	}

	o.abiMarketplace, err = loadABIFile("contracts/abi/FantomMarketplace.json")
	if err != nil {
		return err
	}

	return nil
}

// loadABIFile reads specified ABI file and returns the decoded ABI.
func loadABIFile(path string) (*abi.ABI, error) {
	// FantomNFTTradable
	data, err := abiFiles.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// parse ABI
	decoded, err := abi.JSON(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return &decoded, nil
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
