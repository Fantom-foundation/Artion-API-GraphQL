// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// coObservedContracts is the name of the collection of observed/monitored contracts.
	coObservedContracts = "observed"

	// fiContractAddress is the name of the field keeping the contract address.
	fiContractAddress = "_id"
)

// AddObservedContract adds the specified observed contract record to the collection.
func (db *MongoDbBridge) AddObservedContract(oc *types.ObservedContract) error {
	col := db.client.Database(db.dbName).Collection(coObservedContracts)
	if db.isObservedContractKnown(col, oc) {
		return nil
	}

	if _, err := col.InsertOne(context.Background(), oc); err != nil {
		log.Errorf("can not add observed contract %s", oc.Address.String(), err.Error())
		return err
	}
	return nil
}

// isObservedContractKnown checks if the given observed contract is already stored in the database.
func (db *MongoDbBridge) isObservedContractKnown(col *mongo.Collection, oc *types.ObservedContract) bool {
	sr := col.FindOne(
		context.Background(),
		bson.D{{Key: fiContractAddress, Value: oc.Address.String()}},
		options.FindOne().SetProjection(bson.D{{Key: fiContractAddress, Value: true}}),
	)
	if sr.Err() != nil {
		if sr.Err() != mongo.ErrNoDocuments {
			log.Errorf("check failed for %s; %s", oc.Address.String(), sr.Err().Error())
		}
		return false
	}
	return true
}

// ObservedContractsAddressList provides list of addresses of all observed contracts
// stored in the persistent database.
func (db *MongoDbBridge) ObservedContractsAddressList() []common.Address {
	col := db.client.Database(db.dbName).Collection(coObservedContracts)

	list := make([]common.Address, 0)
	fi, err := col.Find(context.Background(), bson.D{}, options.Find().SetProjection(bson.D{{Key: fiContractAddress, Value: true}}))
	if err != nil {
		log.Errorf("can not pull observed contracts; %s", err.Error())
		return list
	}

	defer func() {
		if err := fi.Close(context.Background()); err != nil {
			log.Errorf("can not close observed contracts list cursor; %s", err.Error())
		}
	}()

	var row struct {
		Addr string `bson:"_id"`
	}
	for fi.Next(context.Background()) {
		if err := fi.Decode(&row); err != nil {
			log.Errorf("failed to decode observed contract; %s", err.Error())
			break
		}
		list = append(list, common.HexToAddress(row.Addr))
	}
	return list
}
