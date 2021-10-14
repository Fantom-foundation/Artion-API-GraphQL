package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

const (
	// coUsers is the name of database collection.
	coUsers = "accounts"

	// fiUserAddress is the column storing the address of the user.
	fiUserAddress = "address"

	fiUserUsername = "alias"

	fiUserEmail = "email"

	fiUserBio = "bio"

	fiUserAvatar = "imageHash"

	fiUserBanner = "bannerHash"

	fiUserCreated = "createdAt"

	fiUserUpdated = "updatedAt"
)

func (sdb *SharedMongoDbBridge) GetUser(address common.Address) (user *types.User, err error) {
	col := sdb.client.Database(sdb.dbName).Collection(coUsers)

	filter := bson.D{{ Key: fiUserAddress, Value: strings.ToLower(address.String()) }}
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

func (sdb *SharedMongoDbBridge) UpsertUser(User *types.User) error {
	if User == nil {
		return fmt.Errorf("no value to store")
	}
	col := sdb.client.Database(sdb.dbName).Collection(coUsers)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{{ Key: fiUserAddress, Value: strings.ToLower(User.Address.String()) }},
		bson.D{
			{ Key: "$set", Value: bson.D{
				{Key: fiUserUsername, Value: User.Username},
				{Key: fiUserEmail, Value: User.Email},
				{Key: fiUserBio, Value: User.Bio},
				{Key: fiUserUpdated, Value: types.Time(time.Now())},
			} },
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fiUserAddress, Value: strings.ToLower(User.Address.String())},
				{Key: fiUserCreated, Value: types.Time(time.Now())},
			}},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		log.Errorf("can not update User; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) SetUserAvatar(user common.Address, imageCid string) error {
	col := sdb.client.Database(sdb.dbName).Collection(coUsers)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{{ Key: fiUserAddress, Value: strings.ToLower(user.String()) }},
		bson.D{
			{ Key: "$set", Value: bson.D{
				{Key: fiUserAvatar, Value: imageCid},
				{Key: fiUserUpdated, Value: types.Time(time.Now())},
			} },
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fiUserAddress, Value: strings.ToLower(user.String())},
				{Key: fiUserCreated, Value: types.Time(time.Now())},
			}},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		log.Errorf("can not update User; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) SetUserBanner(user common.Address, imageCid string) error {
	col := sdb.client.Database(sdb.dbName).Collection(coUsers)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{{ Key: fiUserAddress, Value: strings.ToLower(user.String()) }},
		bson.D{
			{ Key: "$set", Value: bson.D{
				{Key: fiUserBanner, Value: imageCid},
				{Key: fiUserUpdated, Value: types.Time(time.Now())},
			} },
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fiUserAddress, Value: strings.ToLower(user.String())},
				{Key: fiUserCreated, Value: types.Time(time.Now())},
			}},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		log.Errorf("can not update User; %s", err)
		return err
	}
	return nil
}
