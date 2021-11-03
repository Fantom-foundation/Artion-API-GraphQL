package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// GetNotificationSettings provides notification settings of the user, if available.
func (p *Proxy) GetNotificationSettings(address common.Address) (user *types.NotificationSettings, err error) {
	return p.shared.GetNotificationSettings(address)
}

// StoreNotificationSettings inserts or updates existing notification settings of the user.
func  (p *Proxy) StoreNotificationSettings(user common.Address, settings types.NotificationSettings) error {
	return p.shared.StoreNotificationSettings(user, settings)
}
