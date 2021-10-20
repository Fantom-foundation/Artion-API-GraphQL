// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
	"time"
)

// TokenPriceDecimalsDiff is the amount of decimals we need to remove from prices.
// Prices come in 18 decimals, we want to preserve 6 decimals on multiplied result by removing 3 from each number.
var TokenPriceDecimalsDiff = new(big.Int).SetInt64(1_000_000_000_000_000)

// TokenPriceDecimalsCorrection represents the value used to reduce price to stored fixed (6) decimals.
var TokenPriceDecimalsCorrection = new(big.Int).Mul(TokenPriceDecimalsDiff, TokenPriceDecimalsDiff)

// Token reads NFT detail from the persistent database.
func (p *Proxy) Token(contract *common.Address, tokenId *hexutil.Big) (*types.Token, error) {
	var key strings.Builder
	key.WriteString("Token")
	key.WriteString(contract.String())
	key.WriteString(tokenId.String())

	token, err, _ := p.callGroup.Do(key.String(), func() (interface{}, error) {
		return p.db.GetToken(contract, (*big.Int)(tokenId))
	})
	return token.(*types.Token), err
}

// GetUnitPriceAt converts token price in pay-tokens to value in unified units for storing in database.
func (p *Proxy) GetUnitPriceAt(marketplace *common.Address, payToken *common.Address, block *big.Int, val *big.Int) int64 {
	// get the unit price of the payToken
	unit, err := p.rpc.GetPayTokenPrice(marketplace, payToken, block)
	if err != nil {
		log.Warningf("unable to get price of pay token %s on marketplace %s (block %d); %s", payToken.String(), marketplace.String(), block.Uint64(), err.Error())
		return 0
	}

	// recalculate to total updated to fixed decimals int64
	return new(big.Int).Div(new(big.Int).Mul(val, unit), TokenPriceDecimalsCorrection).Int64()
}

// StoreToken puts the given token into the persistent storage.
// The function is used for both insert and update operation.
func (p *Proxy) StoreToken(token *types.Token) error {
	return p.db.StoreToken(token)
}

// UpdateTokenMetadata updates basic metadata of the NFT token.
func (p *Proxy) UpdateTokenMetadata(nft *types.Token) error {
	return p.db.UpdateTokenMetadata(nft)
}

// UpdateTokenMetadataRefreshSchedule sets the NFT metadata update schedule time.
func (p *Proxy) UpdateTokenMetadataRefreshSchedule(nft *types.Token) error {
	return p.db.UpdateTokenMetadataRefreshSchedule(nft)
}

// TokenMetadataRefreshSet pulls s set of NFT tokens scheduled to be updated up to this time.
func (p *Proxy) TokenMetadataRefreshSet() ([]*types.Token, error) {
	return p.db.TokenMetadataRefreshSet()
}

// TokenMarkListed marks the given NFT as listed for direct sale for the given price.
func (p *Proxy) TokenMarkListed(contract *common.Address, tokenID *big.Int, price int64, ts *time.Time) error {
	return p.db.TokenMarkListed(contract, tokenID, price, ts)
}

// TokenMarkOffered marks the given NFT as having offer for the given price.
func (p *Proxy) TokenMarkOffered(contract *common.Address, tokenID *big.Int, price int64, ts *time.Time) error {
	return p.db.TokenMarkOffered(contract, tokenID, price, ts)
}

// TokenMarkAuctioned marks the given NFT as having auction for the given price.
func (p *Proxy) TokenMarkAuctioned(contract *common.Address, tokenID *big.Int, price int64, ts *time.Time) error {
	return p.db.TokenMarkAuctioned(contract, tokenID, price, ts)
}

// TokenMarkBid marks the given NFT as having auction bid for the given price.
func (p *Proxy) TokenMarkBid(contract *common.Address, tokenID *big.Int, price int64, ts *time.Time) error {
	return p.db.TokenMarkBid(contract, tokenID, price, ts)
}

// TokenMarkUnlisted marks the given NFT as not listed for direct sale.
func (p *Proxy) TokenMarkUnlisted(contract *common.Address, tokenID *big.Int) error {
	return p.db.TokenMarkUnlisted(contract, tokenID)
}

// TokenMarkUnOffered marks the given NFT as not having offer anymore.
func (p *Proxy) TokenMarkUnOffered(contract *common.Address, tokenID *big.Int) error {
	return p.db.TokenMarkUnOffered(contract, tokenID)
}

// TokenMarkUnAuctioned marks the given NFT as not auctioned.
func (p *Proxy) TokenMarkUnAuctioned(contract *common.Address, tokenID *big.Int) error {
	return p.db.TokenMarkUnAuctioned(contract, tokenID)
}

// TokenMarkUnBid marks the given NFT as not having a bid anymore.
func (p *Proxy) TokenMarkUnBid(contract *common.Address, tokenID *big.Int) error {
	return p.db.TokenMarkUnBid(contract, tokenID)
}

// TokenMarkSold marks the given NFT as sold on a listing for direct sale for the given price.
func (p *Proxy) TokenMarkSold(contract *common.Address, tokenID *big.Int, price int64, ts *time.Time) error {
	return p.db.TokenMarkSold(contract, tokenID, price, ts)
}

func (p *Proxy) ListTokens(filter *types.TokenFilter, sorting sorting.TokenSorting, sortDesc bool, cursor types.Cursor, count int, backward bool) (list *types.TokenList, err error) {
	return p.db.ListTokens(filter, sorting, sortDesc, cursor, count, backward)
}

func (p *Proxy) GetTokenJsonMetadata(uri string) (*types.JsonMetadata, error) {
	var key strings.Builder
	key.WriteString("GetTokenJsonMetadata")
	key.WriteString(uri)

	jsonMetadata, err, _ := p.callGroup.Do(key.String(), func() (interface{}, error) {
		return p.uri.GetJsonMetadata(uri)
	})
	return jsonMetadata.(*types.JsonMetadata), err
}

func (p *Proxy) GetImage(uri string) (image *types.Image, err error) {
	key := "GetImage" + uri
	data, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.uri.GetImage(uri)
	})
	return data.(*types.Image), err
}

func (p *Proxy) GetImageThumbnail(uri string) (image *types.Image, err error) {
	key := "GetImageThumbnail" + uri
	data, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		image, err := p.GetImage(uri)
		if err != nil || image == nil {
			return nil, fmt.Errorf("getImage failed; %s", err)
		}
		thumb, err := createThumbnail(*image)
		if err != nil {
			return nil, fmt.Errorf("createThumbnail failed; %s", err)
		}
		return &thumb, nil
	})
	return data.(*types.Image), err
}
