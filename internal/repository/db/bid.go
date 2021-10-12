// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
)

const (
	// coAuctionBids is the name of database collection.
	coAuctionBids = "bids"

	// fiAuctionBidContract is the name of the DB column storing NFT contract address.
	fiAuctionBidContract = "contract"

	// fiAuctionBidTokenId represents the name of the DB column storing NFT token ID.
	fiAuctionBidTokenId = "token"
)

// StoreAuctionBid stores given auction bid into the database.
func (db *MongoDbBridge) StoreAuctionBid(bid *types.AuctionBid) error {
	if bid == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coAuctionBids)

	// try to do the insert
	id := bid.ID()
	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: fieldId, Value: id}},
		bson.D{
			{Key: "$set", Value: bid},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fieldId, Value: id},
			}},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		log.Errorf("can not add auction bid; %s", err)
		return err
	}
	return nil
}

// DeleteAuctionBid removes a bid of specific bidder from specific auction.
func (db *MongoDbBridge) DeleteAuctionBid(contract *common.Address, tokenID *big.Int, bidder *common.Address) error {
	col := db.client.Database(db.dbName).Collection(coAuctionBids)

	dr, err := col.DeleteOne(context.Background(), bson.D{
		{Key: fieldId, Value: types.AuctionBidID(contract, tokenID, bidder)},
	})
	if err != nil {
		log.Errorf("auction bid delete failed; %s", err.Error())
		return err
	}

	if dr.DeletedCount > 0 {
		log.Infof("%s bid removed from auction on %s/%s", bidder.String(), contract.String(), (*hexutil.Big)(tokenID).String())
	}
	return nil
}

// ClearAuctionBids removes all the bids stored for the given auction.
func (db *MongoDbBridge) ClearAuctionBids(contract *common.Address, tokenID *big.Int) error {
	col := db.client.Database(db.dbName).Collection(coAuctionBids)
	dr, err := col.DeleteMany(context.Background(), bson.D{
		{Key: fiAuctionBidContract, Value: *contract},
		{Key: fiAuctionBidTokenId, Value: (hexutil.Big)(*tokenID)},
	})
	if err != nil {
		log.Errorf("auction bids cleanup failed; %s", err.Error())
		return err
	}

	if dr.DeletedCount > 0 {
		log.Infof("%d bids cleared from auction on %s/%s", dr.DeletedCount, contract.String(), (*hexutil.Big)(tokenID).String())
	}
	return nil
}
