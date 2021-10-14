package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

func (p *Proxy) AddFollow(follow *types.Follow) error {
	return p.shared.AddFollow(follow)
}

func (p *Proxy) RemoveFollow(follow *types.Follow) error {
	return p.shared.RemoveFollow(follow)
}

func (p *Proxy) ListUserFollowed(user common.Address, cursor types.Cursor, count int, backward bool) (out *types.FollowList, err error) {
	return p.shared.ListUserFollowed(user, cursor, count, backward)
}

func (p *Proxy) ListUserFollowers(user common.Address, cursor types.Cursor, count int, backward bool) (out *types.FollowList, err error) {
	return p.shared.ListUserFollowers(user, cursor, count, backward)
}
