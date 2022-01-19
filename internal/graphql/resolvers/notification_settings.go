package resolvers

import (
	"artion-api-graphql/internal/auth"
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"context"
)

type NotificationSettings types.NotificationSettings

func (rs *RootResolver) NotificationSettings(ctx context.Context) (*NotificationSettings, error) {
	user, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return nil, err
	}
	settings, err := repository.R().GetNotificationSettings(*user)
	if settings == nil {
		return nil, err
	} else {
		return (*NotificationSettings)(settings), err
	}
}

func (rs *RootResolver) UpdateNotificationSettings(ctx context.Context, args struct {
	Settings types.NotificationSettings
}) (bool, error) {
	user, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return false, err
	}
	err = repository.R().StoreNotificationSettings(*user, args.Settings)
	return err == nil, err
}
