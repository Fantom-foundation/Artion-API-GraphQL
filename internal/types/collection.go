// Package types provides high level structures for the API server.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection represents an Artion token collection, represented by an NFT contract.
// Artion basically recognizes NFT contracts deployed form a designated factory.
type Collection struct {
	Address  common.Address `bson:"_id"`
	Type     string         `bson:"type"`
	Name     string         `bson:"name"`
	Symbol   string         `bson:"symbol"`
	Created  Time           `bson:"created"`
	IsActive bool           `bson:"is_active"`
}

// CollectionsIndexes provides list of indexes expected on collections.
func CollectionsIndexes() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 1)
	
	ixType := "ix_type"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "type", Value: 1}}, Options: &options.IndexOptions{Name: &ixType}}
	return ix
}
