package resolvers

import (
	"artion-api-graphql/internal/auth"
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"context"
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
	return &user.dbUser.Avatar, nil
}

func (user User) AvatarProxy() (*string, error) {
	if user.dbUser == nil || user.dbUser.Avatar == "" {
		return nil, nil
	}
	url := "/user-avatar/" + user.Address.String()
	return &url, nil
}

func (user User) Banner() (*string, error) {
	if user.dbUser == nil || user.dbUser.Banner == "" {
		return nil, nil
	}
	return &user.dbUser.Banner, nil
}

func (user User) Ownerships(args struct{ PaginationInput }) (con *OwnershipConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListOwnerships(nil, nil, &user.Address, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewOwnershipConnection(list)
}

func (user User) TokenLikes(args struct{ PaginationInput }) (con *TokenLikeConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListUserTokenLikes(&user.Address, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewTokenLikeConnection(list)
}

func (rs *RootResolver) User(args struct{ Address common.Address }) (user User, err error) {
	dbUser, err := repository.R().GetUser(args.Address)
	if err != nil {
		return User{}, err
	} else {
		return User{ Address: args.Address, dbUser: dbUser }, nil
	}
}

func (rs *RootResolver) LoggedUser(ctx context.Context) (user *User, err error) {
	address, err := auth.GetIdentityOrNil(ctx)
	if address == nil || err != nil {
		return nil, err
	}
	dbUser, err := repository.R().GetUser(*address)
	if err != nil {
		return nil, err
	} else {
		return &User{ Address: *address, dbUser: dbUser }, nil
	}
}

func (rs *RootResolver) UpdateUser(ctx context.Context, args struct{
	Username string
	Bio      string
	Email    string
	Avatar   string
	Banner   string
}) (bool, error) {
	address, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return false, err
	}
	user := types.User{
		Address: *address,
		Username: args.Username,
		Bio: args.Bio,
		Email: args.Email,
		Avatar: args.Avatar,
		Banner: args.Banner,
	}
	err = repository.R().UpsertUser(&user)
	return err == nil, err
}

func (rs *RootResolver) InitiateLogin() (string, error) {
	return auth.GetAuthenticator().GenerateChallenge()
}

func (rs *RootResolver) Login(args struct{
	User      common.Address
	Challenge string
	Signature string
}) (string, error) {
	return auth.GetAuthenticator().GenerateBearer(args.Challenge, args.User, args.Signature)
}
