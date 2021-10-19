package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
)

type Category types.Category

func (rs *RootResolver) Categories() (out []Category, err error) {
	list, err := repository.R().ListCategories()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(list); i++ {
		out = append(out, Category(list[i]))
	}
	return out, nil
}
