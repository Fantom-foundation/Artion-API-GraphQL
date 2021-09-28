package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

type User struct {
	Address common.Address
	dbUser *types.User // data for user loaded from Mongo
}

func (user User) Username() (*string, error) {
	if user.dbUser == nil {
		return nil, nil
	}
	return &user.dbUser.Username, nil
}

func (user User) Bio() (*string, error) {
	if user.dbUser == nil {
		return nil, nil
	}
	return &user.dbUser.Bio, nil
}

func (user User) Email() (*string, error) {
	if user.dbUser == nil {
		return nil, nil
	}
	// TODO: check permission
	return &user.dbUser.Email, nil
}

func (user User) Avatar() (*string, error) {
	if user.dbUser == nil || user.dbUser.Avatar == "" {
		return nil, nil
	}
	// TODO: replace IPFS links by proxy?
	return &user.dbUser.Avatar, nil
}

func (rs *RootResolver) User(args struct{ Address common.Address }) (user User, err error) {
	dbUser, err := repository.R().GetUser(args.Address)
	if err != nil {
		return User{}, err
	} else {
		return User{ Address: args.Address, dbUser: dbUser }, nil
	}
}

func (rs *RootResolver) UpdateUser(args struct{ User types.User }) (bool, error) {
	// TODO: check permission
	err := repository.R().UpsertUser(&args.User)
	return err == nil, err
}
