package resolvers

import (
	"artion-api-graphql/internal/auth"
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type TokenLike types.TokenLike

type TokenLikeEdge struct {
	Node *TokenLike
}

func (edge TokenLikeEdge) Cursor() (types.Cursor, error) {
	return sorting.TokenLikeSortingNone.GetCursor((*types.TokenLike)(edge.Node))
}

type TokenLikeConnection struct {
	Edges      []TokenLikeEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

func NewTokenLikeConnection(list *types.TokenLikeList) (con *TokenLikeConnection, err error) {
	con = new(TokenLikeConnection)
	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]TokenLikeEdge, len(list.Collection))
	for i := 0; i < len(list.Collection); i++ {
		con.Edges[i].Node = (*TokenLike)(list.Collection[i])
	}
	con.PageInfo.HasNextPage = list.HasNext
	con.PageInfo.HasPreviousPage = list.HasPrev
	if len(list.Collection) > 0 {
		startCur, err := con.Edges[0].Cursor()
		if err != nil {
			return nil, err
		}
		endCur, err := con.Edges[len(con.Edges)-1].Cursor()
		if err != nil {
			return nil, err
		}
		con.PageInfo.StartCursor = &startCur
		con.PageInfo.EndCursor = &endCur
	}
	return con, err
}

func (o TokenLike) TokenId() (hexutil.Big, error) {
	return *(*hexutil.Big)(big.NewInt(int64(o.TokenId32))), nil
}

func (o TokenLike) Token() (*Token, error) {
	return NewToken(&o.Contract, (*hexutil.Big)(big.NewInt(int64(o.TokenId32))))
}

func (rs *RootResolver) LoggedUserTokenLikes(ctx context.Context, args struct{ PaginationInput }) (con *TokenLikeConnection, err error) {
	user, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return nil, err
	}
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListUserTokenLikes(user, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewTokenLikeConnection(list)
}

func (rs *RootResolver) LikeToken(ctx context.Context, args struct {
	Contract common.Address
	TokenId  hexutil.Big
}) (bool, error) {
	user, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return false, err
	}
	tokenLike := types.TokenLike{
		User:      *user,
		Contract:  args.Contract,
		TokenId32: int32(args.TokenId.ToInt().Int64()),
	}
	err = repository.R().AddTokenLike(&tokenLike)
	return err == nil, err
}

func (rs *RootResolver) UnlikeToken(ctx context.Context, args struct {
	Contract common.Address
	TokenId  hexutil.Big
}) (bool, error) {
	user, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return false, err
	}
	tokenLike := types.TokenLike{
		User:      *user,
		Contract:  args.Contract,
		TokenId32: int32(args.TokenId.ToInt().Int64()),
	}
	err = repository.R().RemoveTokenLike(&tokenLike)
	return err == nil, err
}
