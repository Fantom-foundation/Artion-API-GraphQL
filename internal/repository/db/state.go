// Package db provides access to the persistent storage.
package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
)

const (
	// coSystemStateCollection is the name of the system status collection.
	coSystemStateCollection = "status"

	// keyLastSeenBlock represents the
	keyLastSeenBlock = "lastSeenBlock"

	// fiBlockNumber represents the name of the last seen block number field.
	fiBlockNumber = "block"
)

// UpdateLastSeenBlock stores the ID of the last seen block
// so the API server can continue where it left off.
func (db *MongoDbBridge) UpdateLastSeenBlock(blk *big.Int) {
	if nil == blk {
		return
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coSystemStateCollection)
	re, err := col.UpdateOne(context.Background(), bson.D{{Key: fieldId, Value: keyLastSeenBlock}}, bson.D{
		{Key: fieldId, Value: keyLastSeenBlock},
		{Key: fiBlockNumber, Value: blk.Int64()},
	}, options.Update().SetUpsert(true))
	if err != nil {
		log.Errorf("can not store last seen block #%d; %s", blk.Uint64(), err.Error())
		return
	}
	if re.UpsertedCount > 0 {
		log.Debugf("last seen block updated to #%d", blk.Uint64())
	}
}

// LastSeenBlockNumber pulls the ID of the last known block from the database.
func (db *MongoDbBridge) LastSeenBlockNumber() (uint64, error) {
	col := db.client.Database(db.dbName).Collection(coSystemStateCollection)
	fi := col.FindOne(context.Background(), bson.D{{Key: fieldId, Value: keyLastSeenBlock}})
	if fi.Err() != nil {
		// no match?
		if fi.Err() == mongo.ErrNoDocuments {
			log.Errorf("previous state not known")
			return 0, nil
		}
		log.Errorf("cannot get block state; %s", fi.Err().Error())
		return 0, fi.Err()
	}

	// decode
	var row struct {
		Num int64 `bson:"block"`
	}
	if err := fi.Decode(&row); err != nil {
		log.Errorf("failed to decode block state; %s", err.Error())
		return 0, err
	}
	log.Noticef("last known block is #%d", row.Num)
	return uint64(row.Num), nil
}
