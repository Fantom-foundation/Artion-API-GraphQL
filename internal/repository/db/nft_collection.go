// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// coNFTCollection is the name of the collection of NFT contracts.
	coNFTCollection = "nft"

	// fiNFTAddress is the name of the field keeping the NFT contract address.
	fiNFTAddress = "_id"
)

// AddNFTCollection adds the specified NFT collection contract record.
func (db *MongoDbBridge) AddNFTCollection(nft *types.NFTCollection) error {
	col := db.client.Database(db.dbName).Collection(coNFTCollection)
	if db.isNFTCollectionKnown(col, nft) {
		return nil
	}

	if _, err := col.InsertOne(context.Background(), nft); err != nil {
		log.Errorf("can not add NFT collection %s", nft.Address.String(), err.Error())
		return err
	}
	return nil
}

// isNFTCollectionKnown checks if the given NFT collection is already stored in the database.
func (db *MongoDbBridge) isNFTCollectionKnown(col *mongo.Collection, nft *types.NFTCollection) bool {
	return db.exists(col, &bson.D{{Key: fiNFTAddress, Value: nft.Address.String()}})
}
