package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

func (p *Proxy) GetUser(address common.Address) (*types.User, error) {
	key := "GetUser-" + address.String()
	user, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.db.GetUser(address)
	})
	return user.(*types.User), err
}

func (p *Proxy) UpsertUser(User *types.User) error {
	return p.db.UpsertUser(User)
}
