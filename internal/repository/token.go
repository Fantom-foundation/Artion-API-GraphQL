package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strconv"
)

func (p *proxy) StoreToken(token *types.Token) error {
	return p.db.StoreToken(token)
}

func (p *proxy) GetToken(nft common.Address, tokenId hexutil.Big) (*types.Token, error) {
	key := "GetToken-" + nft.String() + "-" + tokenId.String()
	token, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.db.GetToken(nft, tokenId)
	})
	return token.(*types.Token), err
}

func (p *proxy) ListTokens(cursor types.Cursor, count int, backward bool) (list *types.TokenList, err error) {
	key := "ListTokens-" + string(cursor) + "-" + strconv.Itoa(count) + strconv.FormatBool(backward)
	tokens, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.db.ListTokens(cursor, count, backward)
	})
	return tokens.(*types.TokenList), err
}
