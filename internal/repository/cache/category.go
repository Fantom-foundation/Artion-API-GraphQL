package cache

import (
	"artion-api-graphql/internal/types"
	"encoding/json"
)

const categoriesCacheKey = "categories"

func (c *MemCache) ListCategories(loader func()(out []types.Category, err error)) (cats []types.Category, err error) {
	data, err := c.cache.Get(categoriesCacheKey)
	if err == nil {
		if err := json.Unmarshal(data, &cats); err != nil {
			return nil, err
		}
		return cats, nil // HIT
	}

	cats, err = loader() // load data from primary source

	data, err = json.Marshal(cats)
	if err != nil {
		log.Errorf("can not encode categories into cache; %s", err)
		return
	}
	err = c.cache.Set(categoriesCacheKey, data)
	if err != nil {
		log.Errorf("can not store categories in cache; %s", err)
	}
	return cats, nil // MIS
}
