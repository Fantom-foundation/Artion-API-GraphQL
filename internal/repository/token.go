// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
)

// Token reads NFT detail from the persistent database.
func (p *Proxy) Token(contract *common.Address, tokenId *hexutil.Big) (*types.Token, error) {
	var key strings.Builder
	key.WriteString("Token")
	key.WriteString(contract.String())
	key.WriteString(tokenId.String())

	token, err, _ := p.callGroup.Do(key.String(), func() (interface{}, error) {
		return p.db.TokenGet(contract, (*big.Int)(tokenId))
	})
	return token.(*types.Token), err
}

// TokenStore puts the given token into the persistent storage.
// The function is used for both insert and update operation.
func (p *Proxy) TokenStore(token *types.Token) error {
	return p.db.TokenStore(token)
}

// TokenUpdateMetadata updates basic metadata of the NFT token.
func (p *Proxy) TokenUpdateMetadata(nft *types.Token) error {
	return p.db.TokenUpdateMetadata(nft)
}

// TokenUpdateMetadataRefreshSchedule sets the NFT metadata update schedule time.
func (p *Proxy) TokenUpdateMetadataRefreshSchedule(nft *types.Token) error {
	return p.db.TokenUpdateMetadataRefreshSchedule(nft)
}

// TokenMetadataRefreshSet pulls s set of NFT tokens scheduled to be updated up to this time.
func (p *Proxy) TokenMetadataRefreshSet() ([]*types.Token, error) {
	return p.db.TokenMetadataRefreshSet()
}

func (p *Proxy) ListTokens(cursor types.Cursor, count int, backward bool) (list *types.TokenList, err error) {
	return p.db.ListTokens(cursor, count, backward)
}

func (p *Proxy) ListCollectionTokens(collection common.Address, cursor types.Cursor, count int, backward bool) (out *types.TokenList, err error) {
	return p.db.ListCollectionTokens(collection, cursor, count, backward)
}

func (p *Proxy) GetTokenJsonMetadata(uri string) (*types.JsonMetadata, error) {
	// TODO: in-memory cache
	var key strings.Builder
	key.WriteString("GetTokenJsonMetadata")
	key.WriteString(uri)

	jsonMetadata, err, _ := p.callGroup.Do(key.String(), func() (interface{}, error) {
		return p.uri.GetJsonMetadata(uri)
	})
	return jsonMetadata.(*types.JsonMetadata), err
}

func (p *Proxy) GetImage(uri string) (image *types.Image, err error) {
	// TODO: in-memory cache
	key := "GetImage" + uri
	data, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.uri.GetImage(uri)
	})
	return data.(*types.Image), err
}
