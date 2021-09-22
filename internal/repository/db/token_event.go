package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// initTokenEventCollection initializes collection with indexes and additional parameters.
func (db *MongoDbBridge) initTokenEventCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index sender and recipient
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiTokenEventSeller, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiTokenEventBuyer, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiTokenEventTime, Value: 1}}})

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
	col := db.client.Database(db.dbName).Collection(types.CoTokenEvents)

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

func (db *MongoDbBridge) ListTokenEvents(filter *bson.D, cursor string, count int) (out *types.TokenEventList, err error) {
	var list types.TokenEventList
	ctx := context.Background()
	col := db.client.Database(db.dbName).Collection(types.CoTokenEvents)

	list.TotalCount, err = col.CountDocuments(ctx, *filter)
	if err != nil {
		db.log.Errorf("can not get total count of token events")
		return nil, err
	}

	if cursor != "" {
		if count > 0 {
			*filter = append(*filter, bson.E{Key: types.FiTokenEventId, Value: bson.D{{Key: "$lt", Value: cursor}}})
		} else {
			*filter = append(*filter, bson.E{Key: types.FiTokenEventId, Value: bson.D{{Key: "$gt", Value: cursor}}})
		}
	}

	var countAbs int
	opt := options.Find()
	if count >= 0 {
		countAbs = count
		opt.SetSort(bson.D{{Key: types.FiTokenEventId, Value: -1 }})
	} else {
		countAbs = -count
		opt.SetSort(bson.D{{Key: types.FiTokenEventId, Value: 1 }})
	}
	opt.SetLimit(int64(countAbs + 1))

	ld, err := col.Find(ctx, filter, opt)
	if err != nil {
		db.log.Errorf("error loading token events list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing token events list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < countAbs {
			var row types.TokenEvent
			if err = ld.Decode(&row); err != nil {
				db.log.Errorf("can not decode the token event in list; %s", err.Error())
				return nil, err
			}
			list.Collection = append(list.Collection, &row)
		} else {
			list.HasNext = true
		}
	}

	if count < 0 {
		list.Reverse()
	}
	return &list, nil
}
