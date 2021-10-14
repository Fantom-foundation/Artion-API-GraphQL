package resolvers

import (
	"artion-api-graphql/internal/auth"
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"context"
)

type ShippingAddress types.ShippingAddress

func (rs *RootResolver) LoggedUserShippingAddress(ctx context.Context) (address *ShippingAddress, err error) {
	user, err := auth.GetIdentityOrErr(ctx)
	if user == nil || err != nil {
		return nil, err
	}
	addr, err := repository.R().GetShippingAddress(*user)
	if err != nil {
		return nil, err
	} else {
		return (*ShippingAddress)(addr), nil
	}
}

func (rs *RootResolver) UpdateShippingAddress(ctx context.Context, args struct{
	Address types.ShippingAddress
}) (bool, error) {
	user, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return false, err
	}
	args.Address.User = *user
	err = repository.R().UpsertShippingAddress(&args.Address)
	return err == nil, err
}
