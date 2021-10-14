package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

const (
	// coFollows is the name of database collection.
	coFollows = "follows"

	// fiFollowFollower is the column storing the address of follower.
	fiFollowFollower = "from"

	// fiFollowFollowed is the column storing the address of followed user.
	fiFollowFollowed = "to"
)

func (sdb *SharedMongoDbBridge) AddFollow(follow *types.Follow) error {
	if follow == nil {
		return fmt.Errorf("no value to store")
	}

	col := sdb.client.Database(sdb.dbName).Collection(coFollows)
	_, err := col.InsertOne(context.Background(), bson.D{
		{Key: fiFollowFollower, Value: strings.ToLower(follow.Follower.String())},
		{Key: fiFollowFollowed, Value: strings.ToLower(follow.Followed.String())},
	})
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil // ignore already-exists error
		}
		log.Errorf("can not add follow; %s", err.Error())
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) RemoveFollow(follow *types.Follow) error {
	if follow == nil {
		return fmt.Errorf("no value to store")
	}

	col := sdb.client.Database(sdb.dbName).Collection(coFollows)
	_, err := col.DeleteOne(
		context.Background(),
		bson.D{
			{Key: fiFollowFollower, Value: strings.ToLower(follow.Follower.String())},
			{Key: fiFollowFollowed, Value: strings.ToLower(follow.Followed.String())},
		},
	)
	if err != nil {
		log.Errorf("can not remove follow; %s", err.Error())
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) ListUserFollowed(user common.Address, cursor types.Cursor, count int, backward bool) (out *types.FollowList, err error) {
	filter := bson.D{ {Key: fiFollowFollower, Value: strings.ToLower(user.String())} }
	return sdb.listFollows(filter, cursor, count, backward)
}

func (sdb *SharedMongoDbBridge) ListUserFollowers(user common.Address, cursor types.Cursor, count int, backward bool) (out *types.FollowList, err error) {
	filter := bson.D{ {Key: fiFollowFollowed, Value: strings.ToLower(user.String())} }
	return sdb.listFollows(filter, cursor, count, backward)
}

func (sdb *SharedMongoDbBridge) listFollows(filter bson.D, cursor types.Cursor, count int, backward bool) (out *types.FollowList, err error) {
	db := (*MongoDbBridge)(sdb)
	var list types.FollowList
	col := db.client.Database(db.dbName).Collection(coFollows)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.FollowSortingNone, backward)
	if err != nil {
		log.Errorf("error loading follows list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			log.Errorf("error closing follow list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.Follow
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the follow in list; %s", err.Error())
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
