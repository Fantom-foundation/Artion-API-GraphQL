package repository

import (
	"artion-api-graphql/internal/types"
)

func (p *Proxy) ListCategories() (out []types.Category, err error) {
	categories, err, _ := p.callGroup.Do("ListCategories", func() (interface{}, error) {
		return p.cache.ListCategories(p.shared.ListCategories)
	})
	return categories.([]types.Category), err
}
