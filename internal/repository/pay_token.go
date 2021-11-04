package repository

import (
	"artion-api-graphql/internal/types"
)

// ListPayTokens provides list of tokens allowed for market payments
func (p *Proxy) ListPayTokens() ([]types.PayToken, error) {
	tokens, err, _ := p.callGroup.Do("ListPayTokens", func() (interface{}, error) {
		return p.cache.ListPayTokens(p.rpc.ListPayTokens)
	})
	return tokens.([]types.PayToken), err
}
