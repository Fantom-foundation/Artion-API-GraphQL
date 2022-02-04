package resolvers

import (
	"artion-api-graphql/internal/auth"
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type Follow types.Follow

func (f Follow) FollowerUser() (User, error) {
	return getUserByAddress(f.Follower)
}

func (f Follow) FollowedUser() (User, error) {
	return getUserByAddress(f.Followed)
}

type FollowEdge struct {
	Node *Follow
}

func (edge FollowEdge) Cursor() (types.Cursor, error) {
	return sorting.FollowSortingNone.GetCursor((*types.Follow)(edge.Node))
}

type FollowConnection struct {
	Edges      []FollowEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

func NewFollowConnection(list *types.FollowList) (con *FollowConnection, err error) {
	con = new(FollowConnection)
	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]FollowEdge, len(list.Collection))
	for i := 0; i < len(list.Collection); i++ {
		con.Edges[i].Node = (*Follow)(list.Collection[i])
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

func (rs *RootResolver) FollowUser(ctx context.Context, args struct {
	User common.Address
}) (bool, error) {
	logged, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return false, err
	}
	balance, err := repository.R().GetBalance(*logged)
	if err != nil {
		return false, err
	}
	if balance.Int64() == 0 {
		return false, fmt.Errorf("your account needs to have a non-zero balance before you can follow a user")
	}
	follow := types.Follow{
		Follower: *logged,
		Followed: args.User,
	}
	err = repository.R().AddFollow(&follow)
	return err == nil, err
}

func (rs *RootResolver) UnfollowUser(ctx context.Context, args struct {
	User common.Address
}) (bool, error) {
	logged, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return false, err
	}
	follow := types.Follow{
		Follower: *logged,
		Followed: args.User,
	}
	err = repository.R().RemoveFollow(&follow)
	return err == nil, err
}
