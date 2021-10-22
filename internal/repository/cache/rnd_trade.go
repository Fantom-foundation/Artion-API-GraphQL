package cache

import (
	"artion-api-graphql/internal/types"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

// headerCacheKey generates a string key for the block header.
func randomTradeCacheKey(adr *common.Address) string {
	var sb strings.Builder
	sb.WriteString("rngt")
	sb.WriteString(types.RandomTradeID(adr).Hex())
	return sb.String()
}

// PullRandomTrade tries to pull the random trade from the in-memory cache, if possible.
func (c *MemCache) PullRandomTrade(adr *common.Address) *types.RandomTrade {
	data, err := c.cache.Get(randomTradeCacheKey(adr))
	if err != nil {
		return nil
	}

	var rt types.RandomTrade
	if err := json.Unmarshal(data, &rt); err != nil {
		return nil
	}
	return &rt
}

// PushRandomTrade tries to store the given random trade to in-memory cache.
func (c *MemCache) PushRandomTrade(rt *types.RandomTrade) {
	if nil == rt {
		return
	}

	data, err := bson.Marshal(rt)
	if err != nil {
		log.Errorf("can not encode random trade; %s", err.Error())
		return
	}

	err = c.cache.Set(randomTradeCacheKey(&rt.Contract), data)
	if err != nil {
		log.Errorf("can not store random trade in cache; %s", err.Error())
	}
}
