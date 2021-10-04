package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
	"strings"
)

// StoreToken puts the given token into the persistent storage.
// The function is used for both insert and update operation.
func (p *Proxy) StoreToken(token *types.Token) error {
	return p.db.StoreToken(token)
}

// GetToken reads NFT token detail from the persistent database.
func (p *Proxy) GetToken(nft *common.Address, tokenId *hexutil.Big) (*types.Token, error) {
	var key strings.Builder
	key.WriteString("GetToken")
	key.WriteString(nft.String())
	key.WriteString(tokenId.String())

	token, err, _ := p.callGroup.Do(key.String(), func() (interface{}, error) {
		return p.db.GetToken(nft, (*big.Int)(tokenId))
	})
	return token.(*types.Token), err
}

func (p *Proxy) ListTokens(cursor types.Cursor, count int, backward bool) (list *types.TokenList, err error) {
	key := "ListTokens-" + string(cursor) + "-" + strconv.Itoa(count) + strconv.FormatBool(backward)
	tokens, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.db.ListTokens(cursor, count, backward)
	})
	return tokens.(*types.TokenList), err
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
