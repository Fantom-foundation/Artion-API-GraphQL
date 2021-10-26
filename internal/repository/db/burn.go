// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// coTokenBurns is the name of the collection keeping token burns.
	coTokenBurns = "burns"
)

// StoreBurn updates NFT burn in the persistent storage.
func (db *MongoDbBridge) StoreBurn(bu *types.NFTBurn) error {
	if bu == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coTokenBurns)
	ctx, cancel := context.WithTimeout(context.Background(), coTokenOwnershipsQueryTimeout)
	defer func() {
		cancel()
	}()

	// try to do the insert
	id := bu.ID()
	rs, err := col.UpdateOne(
		ctx,
		bson.D{{Key: fieldId, Value: id}},
		bson.D{
			{Key: "$set", Value: bu},
			{Key: "$setOnInsert", Value: bson.D{{
				Key: fieldId, Value: id,
			}}},
		},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Errorf("can not store burn %s of %s / #%s; %s",
			bu.Owner.String(), bu.Contract.String(), bu.TokenId.String(), err.Error())
		return err
	}
	if rs.UpsertedCount > 0 {
		log.Infof("token %s / #%s burn by %s stored", bu.Contract.String(), bu.TokenId.String(), bu.Owner.String())
	}
	return nil
}
