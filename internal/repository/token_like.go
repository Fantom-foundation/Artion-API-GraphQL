package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
	"time"
)

func (p *Proxy) AddTokenLike(tokenLike *types.TokenLike) (err error) {
	err = p.shared.AddTokenLike(tokenLike)
	if err == nil {
		p.updateTokenLikesViews(&tokenLike.Contract, big.NewInt(int64(tokenLike.TokenId32)))
	}
	return err
}

func (p *Proxy) RemoveTokenLike(tokenLike *types.TokenLike) (err error) {
	err = p.shared.RemoveTokenLike(tokenLike)
	if err == nil {
		p.updateTokenLikesViews(&tokenLike.Contract, big.NewInt(int64(tokenLike.TokenId32)))
	}
	return err
}

// GetTokenLikesCount get amount of likes for given token
func (p *Proxy) GetTokenLikesCount(contract *common.Address, token *big.Int) (count int64, err error) {
	return p.shared.GetTokenLikesCount(contract, token)
}

// IsTokenLiked get like status of the token for given user
func (p *Proxy) IsTokenLiked(user *common.Address, contract *common.Address, tokenId *big.Int) (bool, error) {
	return p.shared.IsTokenLiked(user, contract, tokenId)
}

func (p *Proxy) ListUserTokenLikes(user *common.Address, cursor types.Cursor, count int, backward bool) (out *types.TokenLikeList, err error) {
	return p.shared.ListUserTokenLikes(user, cursor, count, backward)
}

// updateTokenLikesViews updates likes/views in tokens collection from shared db.
// To be called when likes/views are updated locally.
func (p *Proxy) updateTokenLikesViews(contract *common.Address, token *big.Int) {
	var key strings.Builder
	key.WriteString("TokenLikesViewsRefresh")
	key.WriteString(contract.String())
	key.WriteString(token.String())

	_, _, _ = p.callGroup.Do(key.String(), func() (interface{}, error) {
		likes, err := p.shared.GetTokenLikesCount(contract, token)
		if err != nil {
			log.Errorf("unable to get token likes from shared db; %s", err)
		}

		views, err := p.shared.GetTokenViews(*contract, *token)
		if err != nil {
			log.Errorf("unable to get token views from shared db; %s", err)
		}

		tok := types.Token{
			Contract:    *contract,
			TokenId:     *(*hexutil.Big)(token),
			CachedLikes: likes,
			CachedViews: views.Int64(),
			LikesUpdate: types.Time(time.Now()),
		}
		err = p.db.TokenLikesViewsStore(&tok)
		if err != nil {
			log.Errorf("unable to store token likes/views; %s", err)
		}

		return nil, nil
	})
}
