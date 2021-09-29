// Package types provides high level structures for the API server.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"time"
)

// NFTCollection represents an Artion token collection, represented by an NFT contract.
// Artion basically recognizes NFT contracts deployed form a designated factory.
type NFTCollection struct {
	Address  common.Address `bson:"_id"`
	Type     string         `bson:"type"`
	Name     string         `bson:"name"`
	Symbol   string         `bson:"symbol"`
	BaseUri  string         `bson:"base_uri"`
	Created  time.Time      `bson:"created"`
	IsActive bool           `bson:"is_active"`
}
