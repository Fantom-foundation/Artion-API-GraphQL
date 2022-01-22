// Package cache provides in-memory shared cache.
package cache

import (
	"artion-api-graphql/internal/types"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// legacyCollectionKey generates a cache key for a user's notification setting.
func legacyCollectionKey(contract common.Address) string {
	var sb strings.Builder
	sb.WriteString("lcol_")
	sb.Write(contract.Bytes())
	return sb.String()
}

// GetLegacyCollection try to get a user's notification setting from cache, or backend.
func (c *MemCache) GetLegacyCollection(contract common.Address, loader func(address common.Address) (*types.LegacyCollection, error)) (*types.LegacyCollection, error) {
	key := legacyCollectionKey(contract)

	data, err := c.cache.Get(key)
	if err == nil {
		lc := new(types.LegacyCollection)
		if err = json.Unmarshal(data, &lc); err != nil {
			return nil, err
		}
		return lc, nil
	}

	// load slow way
	lc, err := loader(contract)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(lc)
	if err != nil {
		log.Errorf("can not encode legacy collection into cache; %s", err)
		return lc, err
	}
	err = c.cache.Set(key, data)
	if err != nil {
		log.Errorf("could not store legacy collection; %s", err.Error())
	}
	return lc, nil
}

// FlushLegacyCollection removes user's notification setting from in-memory cache.
func (c *MemCache) FlushLegacyCollection(contract common.Address) {
	err := c.cache.Delete(legacyCollectionKey(contract))
	if err != nil {
		log.Warningf("could not clear legacy collection cache; %s", err.Error())
	}
}
