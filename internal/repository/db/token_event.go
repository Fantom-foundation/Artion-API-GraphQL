package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// coTransaction is the name of database collection.
	coTokenEvents = "token_events"

	// BSON attributes names used in the database collection.
	fiTokenEventSeller = "seller"
	fiTokenEventBuyer  = "buyer"
	fiTokenEventTime   = "evtTime"
)

// initTokenEventCollection initializes collection with indexes and additional parameters.
func (db *MongoDbBridge) initTokenEventCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index sender and recipient
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: fiTokenEventSeller, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: fiTokenEventBuyer, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: fiTokenEventTime, Value: 1}}})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for transaction collection; %s", err.Error())
	}

	// log we are done that
	db.log.Debugf("transactions collection initialized")
}

func (db *MongoDbBridge) StoreTokenEvent(event *types.TokenEvent) error {
	if event == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coTokenEvents)

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), event); err != nil {
		db.log.Errorf("can not store TokenEvent; %s", err)
		return err
	}
	// make sure gas price collection is initialized
	if db.initTokenEvents != nil {
		db.initTokenEvents.Do(func() { db.initTokenEventCollection(col); db.initTokenEvents = nil })
	}
	return nil
}
