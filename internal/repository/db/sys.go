// Package db provides access to the persistent storage.
package db

import (
	"context"
	eth "github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	// coSystemStateCollection is the name of the system status collection.
	coSystemStateCollection = "Artion"

	// keyLastSeenBlock represents the
	keyLastSeenBlock = "LastBlock"

	// fiBlockNumber represents the name of the last seen block number field.
	fiBlockNumber = "number"

	// fiBlockHash represents the name of the last seen block hash field.
	fiBlockHash = "hash"

	// fiBlockTimeStamp represents the name of the last seen block time stamp field.
	fiBlockTimeStamp = "stamp"
)

// UpdateLastSeenBlock stores the ID of the last seen block
// so the API server can continue where it left off.
func (db *MongoDbBridge) UpdateLastSeenBlock(blk *eth.Header) {
	if nil == blk {
		return
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coSystemStateCollection)
	re, err := col.UpdateOne(context.Background(), bson.D{{Key: fieldId, Value: keyLastSeenBlock}}, bson.D{
		{Key: fiBlockNumber, Value: blk.Number.Int64()},
		{Key: fiBlockHash, Value: blk.Hash().String()},
		{Key: fiBlockTimeStamp, Value: time.Unix(int64(blk.Time), 0)},
	}, options.Update().SetUpsert(true))
	if err != nil {
		log.Errorf("can not store last seen block #%d; %s", blk.Number.Uint64(), err.Error())
		return
	}
	if re.UpsertedCount > 0 {
		log.Debugf("last seen block updated to #%d", blk.Number.Uint64())
	}
}
