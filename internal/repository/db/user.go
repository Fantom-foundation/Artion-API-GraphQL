package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// CoUsers is the name of database collection.
	coUsers = "users"
)

func (db *SharedMongoDbBridge) GetUser(address common.Address) (user *types.User, err error) {
	col := db.client.Database(db.dbName).Collection(coUsers)

	filter := bson.D{{ Key: fieldId, Value: address.String() }}
	result := col.FindOne(context.Background(), filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var row types.User
	if err = result.Decode(&row); err != nil {
		log.Errorf("can not decode user; %s", err.Error())
		return nil, err
	}

	return &row, err
}

func (db *SharedMongoDbBridge) UpsertUser(User *types.User) error {
	if User == nil {
		return fmt.Errorf("no value to store")
	}
	col := db.client.Database(db.dbName).Collection(coUsers)

	filter := bson.D{{ Key: fieldId, Value: User.Address.String() }}
	if _, err := col.ReplaceOne(context.Background(), filter, User, options.Replace().SetUpsert(true)); err != nil {
		log.Errorf("can not update User; %s", err)
		return err
	}
	return nil
}
