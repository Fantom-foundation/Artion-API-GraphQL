// Package cache provides in-memory shared cache.
package cache

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// notificationSettingKey generates a cache key for a user's notification setting.
func notificationSettingKey(user *common.Address) string {
	var sb strings.Builder
	sb.WriteString("ntf_set_")
	sb.Write(user.Bytes())
	return sb.String()
}

// GetNotificationSetting try to get a user's notification setting from cache, or backend.
func (c *MemCache) GetNotificationSetting(user *common.Address, loader func(address common.Address) (*types.NotificationSettings, error)) (result *types.NotificationSettings, err error) {
	key := notificationSettingKey(user)
	data, err := c.cache.Get(key)
	if err == nil {
		if err = result.Unmarshal(data); err != nil {
			return nil, err
		}
	}

	// load slow way
	result, err = loader(*user)
	if err != nil {
		return nil, err
	}

	err = c.cache.Set(key, result.Marshal())
	if err != nil {
		log.Errorf("could not store notification setting; %s", err.Error())
	}

	return result, nil
}

// FlushNotificationSetting removes user's notification setting from in-memory cache.
func (c *MemCache) FlushNotificationSetting(user *common.Address) {
	err := c.cache.Delete(notificationSettingKey(user))
	if err != nil {
		log.Warningf("could not clear notification setting; %s", err.Error())
	}
}
