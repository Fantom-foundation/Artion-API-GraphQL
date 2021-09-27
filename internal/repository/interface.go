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

	StoreTokenEvent(*types.TokenEvent) error

	ListTokenEvents(nft *common.Address, tokenId *hexutil.Big, account *common.Address, cursor types.Cursor, count int, backward bool) (*types.TokenEventList, error)

	StoreToken(*types.Token) error

	GetToken(nft common.Address, tokenId hexutil.Big) (*types.Token, error)

	GetTokenJsonMetadata(uri string) (*types.JsonMetadata, error)

	GetTokenImage(uri string) ([]byte, error)

	ListTokens(cursor types.Cursor, count int, backward bool) (*types.TokenList, error)

}