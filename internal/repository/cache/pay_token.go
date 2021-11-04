package cache

import (
	"artion-api-graphql/internal/types"
	"encoding/json"
)

const payTokensCacheKey = "payTokens"

// ListPayTokens provides list of tokens allowed for market payments
func (c *MemCache) ListPayTokens(loader func()(out []types.PayToken, err error)) (payTokens []types.PayToken, err error) {
	data, err := c.cache.Get(payTokensCacheKey)
	if err == nil {
		if err := json.Unmarshal(data, &payTokens); err != nil {
			return nil, err
		}
		return payTokens, nil // HIT
	}

	payTokens, err = loader() // load data from primary source

	data, err = json.Marshal(payTokens)
	if err != nil {
		log.Errorf("can not encode pay tokens into cache; %s", err)
		return
	}
	err = c.cache.Set(payTokensCacheKey, data)
	if err != nil {
		log.Errorf("can not store pay tokens in cache; %s", err)
	}
	return payTokens, nil // MIS
}
