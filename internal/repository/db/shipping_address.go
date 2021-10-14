package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	// coShippingAddress is the name of database collection.
	coShippingAddress = "shippingaddress"

	// fiUserAddress is the column storing the address of the user.
	fiShippingAddressUser = "user"
)

func (sdb *SharedMongoDbBridge) GetShippingAddress(user common.Address) (address *types.ShippingAddress, err error) {
	col := sdb.client.Database(sdb.dbName).Collection(coShippingAddress)

	filter := bson.D{{ Key: fiShippingAddressUser, Value: user.String() }}
	result := col.FindOne(context.Background(), filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var row types.ShippingAddress
	if err = result.Decode(&row); err != nil {
		log.Errorf("can not decode shipping address; %s", err.Error())
		return nil, err
	}

	return &row, err
}

func (sdb *SharedMongoDbBridge) UpsertShippingAddress(address *types.ShippingAddress) error {
	if address == nil {
		return fmt.Errorf("no value to store")
	}
	address.Updated = types.Time(time.Now())

	col := sdb.client.Database(sdb.dbName).Collection(coShippingAddress)

	if _, err := col.ReplaceOne(
		context.Background(),
		bson.D{{ Key: fiShippingAddressUser, Value: address.User.String() }},
		address,
		options.Replace().SetUpsert(true),
	); err != nil {
		log.Errorf("can not update shipping address; %s", err)
		return err
	}
	return nil
}
