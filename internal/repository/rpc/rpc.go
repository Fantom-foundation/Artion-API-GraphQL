// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import "artion-api-graphql/internal/config"

// opera represents the implementation of the Blockchain interface for Fantom Opera node.
type opera struct {
}

// New provides a new instance of the RPC access point.
func New(cfg *config.Config) Blockchain {
	return &opera{}
}
