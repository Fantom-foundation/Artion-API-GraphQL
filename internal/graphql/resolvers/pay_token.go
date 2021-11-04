package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
)

// PayTokens provides list of tokens supported for payments on the marketplace
func (rs *RootResolver) PayTokens() ([]types.PayToken, error) {
	return repository.R().ListPayTokens()
}
