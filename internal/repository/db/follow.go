package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// AddFollow adds the follow into the shared backed database.
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

// RemoveFollow removes the follow from the shared database.
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

// ListUserFollowed gets a list of profiles followed by the user.
func (sdb *SharedMongoDbBridge) ListUserFollowed(user common.Address, cursor types.Cursor, count int, backward bool) (out *types.FollowList, err error) {
	filter := bson.D{{Key: fiFollowFollower, Value: strings.ToLower(user.String())}}
	return sdb.listFollows(filter, cursor, count, backward)
}

// ListUserFollowers gets a list of profiles of followers of the given user.
func (sdb *SharedMongoDbBridge) ListUserFollowers(user common.Address, cursor types.Cursor, count int, backward bool) (out *types.FollowList, err error) {
	filter := bson.D{{Key: fiFollowFollowed, Value: strings.ToLower(user.String())}}
	return sdb.listFollows(filter, cursor, count, backward)
}

// Followers provides a list of followers' addresses for the given user.
func (sdb *SharedMongoDbBridge) Followers(user common.Address) ([]common.Address, error) {
	col := sdb.client.Database(sdb.dbName).Collection(coFollows)
	ld, err := col.Find(context.Background(),
		bson.D{{Key: fiFollowFollowed, Value: strings.ToLower(user.String())}},
		options.Find().SetProjection(bson.D{{Key: fiFollowFollowed, Value: 1}}),
	)
	if err != nil {
		log.Errorf("can not load followers of %s; %s", user.String(), err.Error())
		return nil, err
	}

	defer closeFindCursor("followers", ld)

	list := make([]common.Address, 0)
	for ld.Next(context.Background()) {
		var row struct {
			To common.Address `bson:"to"`
		}

		if err := ld.Decode(&row); err != nil {
			log.Errorf("can not decode follower; %s", err.Error())
			continue
		}

		list = append(list, row.To)
	}

	return list, nil
}

// listFollows performs the followers loading for the given filter and set of limits.
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

	defer closeFindCursor("follow list", ld)

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
