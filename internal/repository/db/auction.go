// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
	"time"
)

const (
	// coAuctions is the name of database collection.
	coAuctions = "auctions"

	// fiAuctionContract is the name of the DB column storing NFT contract address.
	fiAuctionContract = "contract"

	// fiAuctionTokenId represents the name of the DB column storing NFT token ID.
	fiAuctionTokenId = "token"

	// fiAuctionOwner represents the name of the DB column storing token owner.
	// fiAuctionOwner = "owner"

	// fiAuctionStartTime represents the name of the DB column storing auction start.
	fiAuctionStartTime = "start"

	// fiAuctionEndTime represents the name of the DB column storing auction end.
	fiAuctionEndTime = "end"

	// fiAuctionClosed represents the name of the DB column storing date/time of auction having been closed.
	fiAuctionClosed = "closed"
)

// GetAuction provides the auction stored in the database, if available.
func (db *MongoDbBridge) GetAuction(contract *common.Address, tokenID *big.Int) (*types.Auction, error) {
	// get the collection
	col := db.client.Database(db.dbName).Collection(coAuctions)

	sr := col.FindOne(context.Background(), bson.D{{Key: fieldId, Value: types.AuctionID(contract, tokenID)}})
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			log.Warningf("auction for %s/%s not found",
				contract.String(), (*hexutil.Big)(tokenID).String())
			return nil, sr.Err()
		}

		log.Errorf("failed to lookup auction for %s/%s; %s",
			contract.String(), (*hexutil.Big)(tokenID).String(), sr.Err().Error())
		return nil, sr.Err()
	}

	// decode
	var row types.Auction
	if err := sr.Decode(&row); err != nil {
		log.Errorf("could not decode auction for %s/%s; %s",
			contract.String(), (*hexutil.Big)(tokenID).String(), sr.Err().Error())
		return nil, sr.Err()
	}
	return &row, nil
}

// StoreAuction adds the provided auction into the database.
func (db *MongoDbBridge) StoreAuction(au *types.Auction) error {
	if au == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coListings)

	// try to do the insert
	id := au.ID()
	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: fieldId, Value: id}},
		bson.D{
			{Key: "$set", Value: au},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fieldId, Value: id},
			}},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		log.Errorf("can not add auction; %s", err)
		return err
	}
	return nil
}

// HasOpenAuctions checks if the given NFT has standing auction(s).
func (db *MongoDbBridge) HasOpenAuctions(contract *common.Address, tokenID *big.Int) bool {
	col := db.client.Database(db.dbName).Collection(coListings)

	// check for count of any started and un-closed auctions with future deadline; we need only 1
	now := types.Time(time.Now().UTC())
	count, err := col.CountDocuments(context.Background(), bson.D{
		{Key: fiAuctionContract, Value: *contract},
		{Key: fiAuctionTokenId, Value: hexutil.Big(*tokenID)},
		{Key: fiAuctionClosed, Value: bson.D{{Key: "$type", Value: 10}}},
		{Key: fiAuctionStartTime, Value: bson.D{{Key: "$lte", Value: now}}},
		{Key: fiAuctionEndTime, Value: bson.D{{Key: "$gt", Value: now}}},
	}, options.Count().SetLimit(1))

	if err != nil {
		log.Errorf("can not count auctions; %s", err.Error())
		return false
	}
	return count > 0
}
