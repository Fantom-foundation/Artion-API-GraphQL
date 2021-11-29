// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"artion-api-graphql/cmd/artionapi/build"
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"sync"
)

// RootResolver is GraphQL resolver of root namespace.
type RootResolver struct {
	wg      sync.WaitGroup
	sigStop chan bool
}

// log represents the logger to be used by the repository.
var log logger.Logger

// config represents the configuration setup used by the repository
// to establish and maintain required connectivity to external services
// as needed.
var cfg *config.Config

// resolver represents a singleton resolver of the root resolver
var resolver *RootResolver

// oneInstance is the sync guarding root resolver singleton creation.
var oneInstance sync.Once

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l.ModuleLogger("graphql")
}

// SetConfig sets the repository configuration to be used to establish
// and maintain external repository connections.
func SetConfig(c *config.Config) {
	cfg = c
}

// Resolver returns a singleton resolver fo the root resolver.
func Resolver() *RootResolver {
	// make sure to instantiate the Repository only once
	oneInstance.Do(func() {
		resolver = newResolver()
	})
	return resolver
}

// new creates an resolver of root resolver and initializes its internal structure.
func newResolver() *RootResolver {
	if cfg == nil {
		panic(fmt.Errorf("missing configuration"))
	}
	if log == nil {
		panic(fmt.Errorf("missing logger"))
	}

	// create new resolver
	rs := RootResolver{
		// create terminator
		sigStop: make(chan bool, 1),
	}

	// handle broadcast and subscriptions in a separate routine
	rs.wg.Add(1)
	go rs.run()

	return &rs
}

// Close terminates resolver's broadcast service.
func (rs *RootResolver) Close() {
	// log
	log.Notice("GraphQL resolver is closing")

	// send the signal
	rs.sigStop <- true
	rs.wg.Wait()
}

// run monitors and handles subscriptions and broadcasts incoming events to their subscribers.
func (rs *RootResolver) run() {
	// sign off on leaving
	defer func() {
		// terminate
		log.Notice("GraphQL resolver done")
		rs.wg.Done()
	}()

	// log action
	log.Notice("GraphQL resolver started")

	// main loop waits for data on any channel and act upon it
	for {
		select {
		case <-rs.sigStop:
			return
		}
	}
}

// Version resolves the current version of the API server.
func (rs *RootResolver) Version() string {
	return build.Short(cfg)
}

// IsApprovedForAll checks if the operator has ApprovedForAll permission to manipulate with tokens of given owner.
func (rs *RootResolver) IsApprovedForAll(args struct {
	Contract common.Address
	Owner    common.Address
	Operator common.Address
}) (bool, error) {
	return repository.R().IsApprovedForAll(&args.Contract, &args.Owner, &args.Operator)
}

// Contracts provides primary used Artion contracts addresses.
func (rs *RootResolver) Contracts() (*types.Contracts, error) {
	return repository.R().BasicContracts(), nil
}
