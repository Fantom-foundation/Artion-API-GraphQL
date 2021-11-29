// Package cache provides in-memory shared cache.
package cache

import (
	"artion-api-graphql/internal/types"
	"fmt"
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
func (c *MemCache) GetNotificationSetting(user *common.Address, loader func(address common.Address) (*types.NotificationSettings, error)) (*types.NotificationSettings, error) {
	key := notificationSettingKey(user)

	data, err := c.cache.Get(key)
	if err == nil {
		ns := new(types.NotificationSettings)
		if err = ns.Unmarshal(data); err != nil {
			return nil, err
		}
		return ns, nil
	}

	// load slow way
	ns, err := loader(*user)
	if err != nil {
		return nil, err
	}

	// do we have any data here?
	if nil == ns {
		return nil, fmt.Errorf("notification settings not availble for %s", user.String())
	}

	err = c.cache.Set(key, ns.Marshal())
	if err != nil {
		log.Errorf("could not store notification setting; %s", err.Error())
	}
	return ns, nil
}

// FlushNotificationSetting removes user's notification setting from in-memory cache.
func (c *MemCache) FlushNotificationSetting(user *common.Address) {
	err := c.cache.Delete(notificationSettingKey(user))
	if err != nil {
		log.Warningf("could not clear notification setting; %s", err.Error())
	}
}
