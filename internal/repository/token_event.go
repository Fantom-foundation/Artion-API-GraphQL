package repository

import "artion-api-graphql/internal/types"

func (p *proxy) StoreTokenEvent(event *types.TokenEvent) error {
	return p.db.StoreTokenEvent(event)
}
