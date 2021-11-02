package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// coActivities is the name of database collection.
	coActivities = "activities"

	// BSON attributes names used in the database collection.
	fiActivityContract = "contract"
	fiActivityTokenId = "token"
	fiActivityFrom = "from"
	fiActivityTo   = "to"
	fiActivityType = "type"
)

func (db *MongoDbBridge) StoreActivity(activity *types.Activity) error {
	if activity == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coActivities)

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), activity); err != nil {
		log.Errorf("can not store activity; %s", err)
		return err
	}
	return nil
}

func (db *MongoDbBridge) ListActivities(contract *common.Address, tokenId *hexutil.Big, user *common.Address, actTypes []types.ActivityType, cursor types.Cursor, count int, backward bool) (out *types.ActivityList, err error) {
	filter := bson.D{}
	if contract != nil {
		filter = append(filter, primitive.E{Key: fiActivityContract, Value: contract.String() })
	}
	if tokenId != nil {
		filter = append(filter, primitive.E{Key: fiActivityTokenId, Value: tokenId.String() })
	}
	if user != nil {
		filter = append(filter, bson.E{
			Key: "$or",
			Value: bson.A{bson.D{{
				Key:   fiActivityFrom,
				Value: user.String(),
			}}, bson.D{{
				Key:   fiActivityTo,
				Value: user.String(),
			}}},
		})
	}
	if actTypes != nil {
		filter = append(filter, bson.E{Key: fiActivityType, Value: bson.D{{Key: "$in", Value: actTypes}}})
	}
	return db.listActivities(filter, cursor, count, backward)
}

func (db *MongoDbBridge) listActivities(filter bson.D, cursor types.Cursor, count int, backward bool) (out *types.ActivityList, err error) {
	var list types.ActivityList
	col := db.client.Database(db.dbName).Collection(coActivities)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.ActivitySortingNone, ! backward)
	if err != nil {
		log.Errorf("error loading activities list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			log.Errorf("error closing activities list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.Activity
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the activity in list; %s", err.Error())
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
