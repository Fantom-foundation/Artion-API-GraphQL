package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (db *MongoDbBridge) ListTokenEvents(nft *common.Address, tokenId *hexutil.Big, account *common.Address, cursor types.Cursor, count int, backward bool) (out *types.TokenEventList, err error) {
	filter := bson.D{}
	if nft != nil {
		filter = append(filter, primitive.E{Key: types.FiTokenEventNft, Value: nft.String() })
	}
	if tokenId != nil {
		filter = append(filter, primitive.E{Key: types.FiTokenEventTokenId, Value: tokenId.String() })
	}
	if account != nil {
		filter = append(filter, bson.E{
			Key: "$or",
			Value: bson.A{bson.D{{
				Key:   types.FiTokenEventSeller,
				Value: account.String(),
			}}, bson.D{{
				Key:   types.FiTokenEventBuyer,
				Value: account.String(),
			}}},
		})
	}
	return db.listTokenEvents(&filter, cursor, count, backward)
}

func (db *MongoDbBridge) listTokenEvents(filter *bson.D, cursor types.Cursor, count int, backward bool) (out *types.TokenEventList, err error) {
	var list types.TokenEventList
	col := db.client.Database(db.dbName).Collection(types.CoTokenEvents)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, backward)
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
		if len(list.Collection) < count {
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

	if cursor != "" {
		list.HasPrev = true
	}
	if backward {
		list.Reverse()
	}
	return &list, nil
}
