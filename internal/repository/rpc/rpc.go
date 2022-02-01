// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository/rpc/contracts"
	"artion-api-graphql/internal/types"
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
	auctionContractsProps      map[common.Address]types.AuctionProps
	auctionContracts           map[common.Address]IAuctionContract
	marketplaceContracts       map[common.Address]IMarketplaceContract
	payTokenPriceContract      IMarketplaceContract // for token royalty or pay token price
	tokenRegistryContract      *contracts.FantomTokenRegistry
	royaltyRegistryContract    *contracts.FantomRoyaltyRegistry
	rngFeedContract            *contracts.RandomNumberOracle

	basicContracts types.Contracts
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
		var ac AuctionContractV1
		ac.auctionV1aContract, err = contracts.NewFantomAuction(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded V0 auction contract at %s", addr.String())
		}
		ac.auctionV1Contract, err = contracts.NewFantomAuctionV1(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded V1 auction contract at %s", addr.String())
		}
		o.auctionContracts[*addr] = &ac
		o.auctionContractsProps[*addr] = types.AuctionV1Props

	case "auction2":
		var ac AuctionContractV2
		ac.auctionV2Contract, err = contracts.NewFantomAuctionV2(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded V2 auction contract at %s", addr.String())
		}
		o.auctionContracts[*addr] = &ac
		o.auctionContractsProps[*addr] = types.AuctionV2Props

	case "auction3":
		var ac AuctionContractV2 // V3 use the same ABI as V2
		ac.auctionV2Contract, err = contracts.NewFantomAuctionV2(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded V3 auction contract at %s", addr.String())
		}
		o.auctionContracts[*addr] = &ac
		o.auctionContractsProps[*addr] = types.AuctionV3Props
		o.basicContracts.AuctionHall = *addr
		log.Warningf("setting basic auction to %s", addr.String())

	case "market":
		var mc MarketplaceContractV1
		mc.marketplace, err = contracts.NewFantomMarketplace(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded %s contract at %s", ct, addr.String())
		}
		o.marketplaceContracts[*addr] = &mc
		o.payTokenPriceContract = &mc

	case "market2":
		var mc MarketplaceContractV1 // V2 use the same ABI as V1
		mc.marketplace, err = contracts.NewFantomMarketplace(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded %s contract at %s", ct, addr.String())
		}
		o.marketplaceContracts[*addr] = &mc

	case "market3":
		var mc MarketplaceContractV1 // V3 use the same ABI as V1
		mc.marketplace, err = contracts.NewFantomMarketplace(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded %s contract at %s", ct, addr.String())
		}
		o.marketplaceContracts[*addr] = &mc
		o.basicContracts.Marketplace = *addr
		log.Warningf("setting basic marketplace to %s", ct, addr.String())

	case "rng":
		o.rngFeedContract, err = contracts.NewRandomNumberOracle(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded %s contract at %s", ct, addr.String())
		}

	case "token_registry":
		o.tokenRegistryContract, err = contracts.NewFantomTokenRegistry(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded %s contract at %s", ct, addr.String())
		}

	case "royalty_registry":
		o.royaltyRegistryContract, err = contracts.NewFantomRoyaltyRegistry(*addr, o.ftm)
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

		auctionContractsProps: make(map[common.Address]types.AuctionProps),
		auctionContracts:     make(map[common.Address]IAuctionContract),
		marketplaceContracts: make(map[common.Address]IMarketplaceContract),
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

// BasicContracts provides addresses of basic Artion contracts.
func (o *Opera) BasicContracts() *types.Contracts {
	return &o.basicContracts
}

// IsEscrowContract test if the address is address of auction or other escrow contract.
func (o *Opera) IsEscrowContract(addr common.Address) bool {
	_, exists := o.auctionContracts[addr]
	return exists
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
