package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p *proxy) StoreToken(token *types.Token) error {
	return p.db.StoreToken(token)
}

func (p *proxy) GetToken(nft common.Address, tokenId hexutil.Big) (*types.Token, error) {
	return p.db.GetToken(nft, tokenId)
}

func (p *proxy) ListTokens(cursor types.Cursor, count int, backward bool) (list *types.TokenList, err error) {
	return p.db.ListTokens(cursor, count, backward)
}
