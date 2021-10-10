// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
	"time"
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

// TokenMarkListed marks the given NFT as listed for direct sale for the given price.
func (p *Proxy) TokenMarkListed(contract *common.Address, tokenID *big.Int, price *hexutil.Big, ts *time.Time) error {
	return p.db.TokenMarkListed(contract, tokenID, price, ts)
}

// TokenMarkUnlisted marks the given NFT as not listed for direct sale.
func (p *Proxy) TokenMarkUnlisted(contract *common.Address, tokenID *big.Int) error {
	return p.db.TokenMarkUnlisted(contract, tokenID)
}

// TokenMarkSold marks the given NFT as sold on a listing for direct sale for the given price.
func (p *Proxy) TokenMarkSold(contract *common.Address, tokenID *big.Int, price *hexutil.Big, ts *time.Time) error {
	return p.db.TokenMarkSold(contract, tokenID, price, ts)
}

func (p *Proxy) ListTokens(sorting sorting.TokenSorting, sortDesc bool, cursor types.Cursor, count int, backward bool) (list *types.TokenList, err error) {
	return p.db.ListTokens(sorting, sortDesc, cursor, count, backward)
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
