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

	// StoreListing stores the token listing into database.
	StoreListing(listing *types.Listing) error

	// ListListings allows to browse all tokens listings, preferably filtered.
	ListListings(nft *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (*types.ListingList, error)

	// StoreOffer stores the token offer into database.
	StoreOffer(offer *types.Offer) error

	// ListOffers allows to browse all tokens offers, preferably filtered.
	ListOffers(nft *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (*types.OfferList, error)

	// GetTokenJsonMetadata obtains token metadata from JSON at given URI
	GetTokenJsonMetadata(uri string) (*types.JsonMetadata, error)

	// GetTokenImage obtains image (including mimetype) from given URI
	GetTokenImage(uri string) (image *types.Image, err error)

	// ListTokens allows to browse all tokens available for Artion.
	ListTokens(cursor types.Cursor, count int, backward bool) (*types.TokenList, error)

}