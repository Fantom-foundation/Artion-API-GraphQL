// Package types provides high level structures for the API server.
package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// Collection represents an Artion token collection, represented by an NFT contract.
// Artion basically recognizes NFT contracts deployed form a designated factory.
type Collection struct {
	Address            common.Address  `bson:"_id"`
	Type               string          `bson:"type"`
	Name               string          `bson:"name"`
	Symbol             string          `bson:"symbol"`
	Created            Time            `bson:"created"`
	Categories         []int32         `bson:"categories"`
	IsActive           bool            `bson:"is_active"`
	VerifiedBy         *common.Address `bson:"verified_by"`
}
