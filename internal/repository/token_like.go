package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (p *Proxy) AddTokenLike(tokenLike *types.TokenLike) error {
	return p.shared.AddTokenLike(tokenLike)
}

func (p *Proxy) RemoveTokenLike(tokenLike *types.TokenLike) error {
	return p.shared.RemoveTokenLike(tokenLike)
}

// GetTokenLikesCount get amount of likes for given token
func (p *Proxy) GetTokenLikesCount(contract *common.Address, token *big.Int) (count int64, err error) {
	return p.shared.GetTokenLikesCount(contract, token)
}

func (p *Proxy) ListUserTokenLikes(user *common.Address, cursor types.Cursor, count int, backward bool) (out *types.TokenLikeList, err error) {
	return p.shared.ListUserTokenLikes(user, cursor, count, backward)
}
