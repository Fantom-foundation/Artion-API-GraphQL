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

	// fiContractType is the name of the field keeping the contract type.
	fiContractType = "type"
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

// ObservedContractAddressByType provides address of an observed contract by its type, if available.
func (db *MongoDbBridge) ObservedContractAddressByType(t string) (*common.Address, error) {
	col := db.client.Database(db.dbName).Collection(coObservedContracts)
	sr := col.FindOne(
		context.Background(),
		bson.D{{Key: "type", Value: t}},
		options.FindOne().SetProjection(bson.D{{Key: "_id", Value: 1}}),
	)

	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			log.Warningf("contract of type %s not found", t)
			return nil, sr.Err()
		}
		log.Errorf("failed to lookup contract of type %s; %s", t, sr.Err().Error())
		return nil, sr.Err()
	}

	var row struct {
		ID common.Address `bson:"_id"`
	}
	if err := sr.Decode(&row); err != nil {
		log.Errorf("failed to decode contract address of type %s; %s", t, sr.Err().Error())
		return nil, sr.Err()
	}
	return &row.ID, nil
}

// isObservedContractKnown checks if the given observed contract is already stored in the database.
func (db *MongoDbBridge) isObservedContractKnown(col *mongo.Collection, oc *types.ObservedContract) bool {
	return db.exists(col, &bson.D{{Key: fiContractAddress, Value: oc.Address.String()}})
}

// MinObservedBlockNumber finds the lowest required block number based on observed deployed contracts.
func (db *MongoDbBridge) MinObservedBlockNumber(def uint64) uint64 {
	col := db.client.Database(db.dbName).Collection(coObservedContracts)
	c, err := col.Aggregate(context.Background(), mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "blk", Value: bson.D{{Key: "$min", Value: "$block"}}},
		}}},
	})
	if err != nil {
		log.Errorf("can not find min block of observed contracts; %s", err.Error())
		return def
	}

	defer func() {
		if err := c.Close(context.Background()); err != nil {
			log.Errorf("failed to close cursor; %s", err.Error())
		}
	}()

	if !c.Next(context.Background()) {
		log.Error("min observed block not available")
		return def
	}

	var row struct {
		Blk int64 `bson:"blk"`
	}
	if err := c.Decode(&row); err != nil {
		log.Errorf("can not decode min observed block; %s", err.Error())
		return def
	}

	log.Infof("the first observed block is #%d", row.Blk)
	return uint64(row.Blk)
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

// NFTContractsTypeMap provides a map of observed contract addresses to corresponding
// contract type for ERC721 and ERC1155 contracts including their factory.
// In case of a factory contract, we need the deployed NFT type for processing.
func (db *MongoDbBridge) NFTContractsTypeMap() map[common.Address]string {
	col := db.client.Database(db.dbName).Collection(coObservedContracts)
	list := make(map[common.Address]string, 0)

	fi, err := col.Find(
		context.Background(),
		bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: fiContractType, Value: types.ContractTypeERC721}},
			bson.D{{Key: fiContractType, Value: types.ContractTypeERC1155}},
		}}},
		options.Find().SetProjection(
			bson.D{
				{Key: fiContractAddress, Value: true},
				{Key: fiContractType, Value: true},
			}),
	)
	if err != nil {
		log.Errorf("can not pull observed contracts; %s", err.Error())
		return list
	}

	defer func() {
		if err := fi.Close(context.Background()); err != nil {
			log.Errorf("can not close observed contracts map cursor; %s", err.Error())
		}
	}()

	var row struct {
		Addr string `bson:"_id"`
		Type string `bson:"type"`
	}
	for fi.Next(context.Background()) {
		if err := fi.Decode(&row); err != nil {
			log.Errorf("failed to decode observed contract; %s", err.Error())
			break
		}
		list[common.HexToAddress(row.Addr)] = row.Type
	}

	log.Noticef("%d NFT contracts known", len(list))
	return list
}
