// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Repository defines interface used to interact with the persistent storage
// and the blockchain node.
type Repository interface {

	// StoreTokenEvent stores the token event into database.
	StoreTokenEvent(*types.TokenEvent) error

	// ListTokenEvents allows to browse all events on tokens, preferably filtered.
	ListTokenEvents(nft *common.Address, tokenId *hexutil.Big, account *common.Address, cursor types.Cursor, count int, backward bool) (*types.TokenEventList, error)

	// StoreToken stores the token info in database.
	StoreToken(*types.Token) error

	// GetToken obtains token info stored in database.
	GetToken(nft common.Address, tokenId hexutil.Big) (*types.Token, error)

	// AddListing stores the token listing into database.
	AddListing(listing *types.Listing) error

	// UpdateListing updates existing token listing in database.
	UpdateListing(listing *types.Listing) error

	// RemoveListing removes token listing from database.
	RemoveListing(owner common.Address, nft common.Address, tokenId hexutil.Big) error

	// ListListings allows to browse all tokens listings, preferably filtered.
	ListListings(nft *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (*types.ListingList, error)

	// AddOffer stores the token offer into database.
	AddOffer(offer *types.Offer) error

	// UpdateOffer updates existing token offer in database.
	UpdateOffer(offer *types.Offer) error

	// RemoveOffer removes token offer from database.
	RemoveOffer(creator common.Address, nft common.Address, tokenId hexutil.Big) error

	// ListOffers allows to browse all tokens offers, preferably filtered.
	ListOffers(nft *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (*types.OfferList, error)

	// GetUser obtains user by address from database.
	GetUser(address common.Address) (user *types.User, err error)

	// UpsertUser stores profiles of existing of new user.
	UpsertUser(User *types.User) error

	// GetTokenJsonMetadata obtains token metadata from JSON at given URI
	GetTokenJsonMetadata(uri string) (*types.JsonMetadata, error)

	// GetImage obtains image (including mimetype) from given URI
	GetImage(uri string) (image *types.Image, err error)

	// ListTokens allows to browse all tokens available for Artion.
	ListTokens(cursor types.Cursor, count int, backward bool) (*types.TokenList, error)

}