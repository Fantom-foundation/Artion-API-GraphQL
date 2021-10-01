// Package db provides access to the persistent storage.
package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// exists checks if the given filter matches at least one document in the given DB collection.
func (db *MongoDbBridge) exists(col *mongo.Collection, filter *bson.D) bool {
	sr := col.FindOne(
		context.Background(),
		filter,
		options.FindOne().SetProjection(bson.D{{Key: fieldId, Value: true}}),
	)
	if sr.Err() != nil {
		if sr.Err() != mongo.ErrNoDocuments {
			log.Errorf("check failed; %s", sr.Err().Error())
		}
		return false
	}
	return true
}
