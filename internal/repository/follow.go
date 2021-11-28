package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// AddFollow adds a new follower record.
func (p *Proxy) AddFollow(follow *types.Follow) error {
	return p.shared.AddFollow(follow)
}

// RemoveFollow removes given follow from the backend.
func (p *Proxy) RemoveFollow(follow *types.Follow) error {
	return p.shared.RemoveFollow(follow)
}

// ListUserFollowed lists set of followed users of the given user.
func (p *Proxy) ListUserFollowed(user common.Address, cursor types.Cursor, count int, backward bool) (out *types.FollowList, err error) {
	return p.shared.ListUserFollowed(user, cursor, count, backward)
}

// ListUserFollowers lists set of followers of the given user.
func (p *Proxy) ListUserFollowers(user common.Address, cursor types.Cursor, count int, backward bool) (out *types.FollowList, err error) {
	return p.shared.ListUserFollowers(user, cursor, count, backward)
}

// MustFollowers provides list of all followers of the given user, or an empty slice - if not available.
func (p *Proxy) MustFollowers(user common.Address) []common.Address {
	list, err := p.shared.Followers(user)
	if err != nil {
		log.Errorf("followers of %s not available; %s", user.String(), err.Error())
		return []common.Address{}
	}
	return list
}
