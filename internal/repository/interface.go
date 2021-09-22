// Package repository implements persistent data access and processing.
package repository

import "artion-api-graphql/internal/types"

// Proxy defines interface used to interact with the persistent storage
// and the blockchain node.
type Proxy interface {

	StoreTokenEvent(*types.TokenEvent) error

}