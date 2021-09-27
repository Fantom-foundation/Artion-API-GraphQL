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

	// GetTokenJsonMetadata obtains token metadata from JSON at given URI
	GetTokenJsonMetadata(uri string) (*types.JsonMetadata, error)

	// GetTokenImage obtains image (including mimetype) from given URI
	GetTokenImage(uri string) (image *types.Image, err error)

	// ListTokens allows to browse all tokens available for Artion.
	ListTokens(cursor types.Cursor, count int, backward bool) (*types.TokenList, error)

}