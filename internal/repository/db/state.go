// Package db provides access to the persistent storage.
package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	// coSystemStateCollection is the name of the system status collection.
	coSystemStateCollection = "status"

	// keyLastSeenBlock represents a key into status collection.
	keyLastSeenBlock = "lastSeenBlock"

	// fiBlockNumber represents the name of the last seen block number field.
	fiBlockNumber = "block"

	// keyLastBanUpdate represents a key into status collection.
	keyLastBanUpdate = "lastBanUpdate"

	// fiTime represents the name of the time field.
	fiTime = "time"
)

// UpdateLastSeenBlock stores the ID of the last seen block
// so the API server can continue where it left off.
func (db *MongoDbBridge) UpdateLastSeenBlock(blk uint64) {
	// get the collection
	col := db.client.Database(db.dbName).Collection(coSystemStateCollection)
	re, err := col.UpdateOne(context.Background(), bson.D{{Key: fieldId, Value: keyLastSeenBlock}}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: fiBlockNumber, Value: int64(blk)},
		}},
		{Key: "$setOnInsert", Value: bson.D{
			{Key: fieldId, Value: keyLastSeenBlock},
		}},
	}, options.Update().SetUpsert(true))
	if err != nil {
		log.Errorf("can not store last seen block #%d; %s", blk, err.Error())
		return
	}
	if re.UpsertedCount > 0 {
		log.Debugf("last seen block updated to #%d", blk)
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

// UpdateLastBanUpdate stores the time of last NFT banned status update
// so the API server can continue where it left off.
func (db *MongoDbBridge) UpdateLastBanUpdate(time time.Time) {
	// get the collection
	col := db.client.Database(db.dbName).Collection(coSystemStateCollection)
	re, err := col.UpdateOne(context.Background(), bson.D{{Key: fieldId, Value: keyLastBanUpdate}}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: fiTime, Value: time},
		}},
		{Key: "$setOnInsert", Value: bson.D{
			{Key: fieldId, Value: keyLastBanUpdate},
		}},
	}, options.Update().SetUpsert(true))
	if err != nil {
		log.Errorf("can not store last ban update time %s; %s", time.String(), err.Error())
		return
	}
	if re.UpsertedCount > 0 {
		log.Debugf("last ban update updated to %s", time.String())
	}
}

// LastBanUpdate pulls the time of the last NFT banned status update
func (db *MongoDbBridge) LastBanUpdate() (*time.Time, error) {
	col := db.client.Database(db.dbName).Collection(coSystemStateCollection)
	fi := col.FindOne(context.Background(), bson.D{{Key: fieldId, Value: keyLastBanUpdate}})
	if fi.Err() != nil {
		// no match?
		if fi.Err() == mongo.ErrNoDocuments {
			log.Errorf("last NFT banned status update not known")
			return nil, nil
		}
		log.Errorf("cannot get last NFT banned status update; %s", fi.Err().Error())
		return nil, fi.Err()
	}

	// decode
	var row struct {
		Time time.Time `bson:"time"`
	}
	if err := fi.Decode(&row); err != nil {
		log.Errorf("failed to decode last NFT banned status update; %s", err.Error())
		return nil, err
	}
	log.Noticef("last NFT banned status update is %s", row.Time.String())
	return &row.Time, nil
}
