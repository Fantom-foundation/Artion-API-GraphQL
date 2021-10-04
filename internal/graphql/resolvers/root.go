// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"artion-api-graphql/cmd/artionapi/build"
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

// instance represents a singleton instance of the root resolver
var instance *RootResolver

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

// Resolver returns a singleton instance fo the root instance.
func Resolver() *RootResolver {
	// make sure to instantiate the Repository only once
	oneInstance.Do(func() {
		instance = newResolver()
	})
	return instance
}

// new creates an instance of root resolver and initializes its internal structure.
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

func (rs *RootResolver) NftCollection(args struct {
	Address common.Address
}) (*NftCollection, error) {
	Collection := NftCollection{Address: args.Address}
	return &Collection, nil
}

func (rs *RootResolver) NftCollections(args struct{ PaginationInput }) (con *NftCollectionConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListNFTCollections(cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewCollectionConnection(list)
}

func (rs *RootResolver) Token(args struct {
	Address common.Address
	TokenId hexutil.Big
}) (*Token, error) {
	token := Token{Address: args.Address, TokenId: args.TokenId}
	return &token, nil
}

func (rs *RootResolver) Tokens(args struct{ PaginationInput }) (con *TokenConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListTokens(cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewTokenConnection(list)
}

func (rs *RootResolver) Listings(args struct{ PaginationInput }) (con *ListingConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListListings(nil, nil, nil, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewListingConnection(list)
}

func (rs *RootResolver) Offers(args struct{ PaginationInput }) (con *OfferConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListOffers(nil, nil, nil, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewOfferConnection(list)
}
