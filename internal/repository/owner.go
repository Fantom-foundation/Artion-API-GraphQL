// Package repository implements persistent data access and processing.
package repository

import "artion-api-graphql/internal/types"

// StoreOwner stores the given NFT owner record in persistent storage.
func (p *Proxy) StoreOwner(to *types.NFTOwner) error {
	return p.db.StoreOwner(to)
}
