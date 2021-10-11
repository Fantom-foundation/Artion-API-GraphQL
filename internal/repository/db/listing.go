// Package db provides access to the persistent storage.
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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
	"time"
)

const (
	// coListings is the name of database collection.
	coListings = "listings"

	// fiListingContract is the name of the DB column storing NFT contract address.
	fiListingContract = "contract"

	// fiListingTokenId represents the name of the DB column storing NFT token ID.
	fiListingTokenId = "token"

	// fiListingOwner represents the name of the DB column storing token owner.
	fiListingOwner = "owner"

	// fiListingStartTime represents the name of the DB column storing listing start.
	fiListingStartTime = "start"

	// fiListingClosed represents the name of the DB column storing date/time of listing having been closed.
	fiListingClosed = "closed"
)

// GetListing provides the token listing stored in the database, if available.
func (db *MongoDbBridge) GetListing(contract *common.Address, tokenID *big.Int, owner *common.Address) (*types.Listing, error) {
	// get the collection
	col := db.client.Database(db.dbName).Collection(coListings)

	sr := col.FindOne(context.Background(), bson.D{{Key: fieldId, Value: types.ListingID(contract, tokenID, owner)}})
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			log.Warningf("could not find listing %s/%s of owner %s; %s",
				contract.String(), (*hexutil.Big)(tokenID).String(), owner.String(), sr.Err().Error())
			return nil, sr.Err()
		}

		log.Errorf("failed to lookup listing %s/%s of owner %s; %s",
			contract.String(), (*hexutil.Big)(tokenID).String(), owner.String(), sr.Err().Error())
		return nil, sr.Err()
	}

	// decode
	var row types.Listing
	if err := sr.Decode(&row); err != nil {
		log.Errorf("could not decode listing %s/%s of owner %s; %s",
			contract.String(), (*hexutil.Big)(tokenID).String(), owner.String(), sr.Err().Error())
		return nil, sr.Err()
	}
	return &row, nil
}

// StoreListing adds the provided listing into the database.
func (db *MongoDbBridge) StoreListing(lst *types.Listing) error {
	if lst == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coListings)

	// try to do the insert
	id := lst.ID()
	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: fieldId, Value: id}},
		bson.D{
			{Key: "$set", Value: lst},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fieldId, Value: id},
			}},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		log.Errorf("can not add Listing; %s", err)
		return err
	}
	return nil
}

// HasOpenListings checks if the given NFT has standing listing(s).
func (db *MongoDbBridge) HasOpenListings(contract *common.Address, tokenID *big.Int) bool {
	col := db.client.Database(db.dbName).Collection(coListings)

	// check for count of any un-closed offers with future deadline; we need only 1
	count, err := col.CountDocuments(context.Background(), bson.D{
		{Key: fiListingContract, Value: *contract},
		{Key: fiListingTokenId, Value: hexutil.Big(*tokenID)},
		{Key: fiListingClosed, Value: bson.D{{Key: "$type", Value: 10}}},
		{Key: fiListingStartTime, Value: bson.D{{Key: "$lte", Value: types.Time(time.Now().UTC())}}},
	}, options.Count().SetLimit(1))

	if err != nil {
		log.Errorf("can not count listings; %s", err.Error())
		return false
	}
	return count > 0
}

func (db *MongoDbBridge) ListListings(contract *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (out *types.ListingList, err error) {
	filter := bson.D{}
	if contract != nil {
		filter = append(filter, primitive.E{Key: fiListingContract, Value: contract.String()})
	}
	if tokenId != nil {
		filter = append(filter, primitive.E{Key: fiListingTokenId, Value: tokenId.String()})
	}
	if owner != nil {
		filter = append(filter, primitive.E{Key: fiListingOwner, Value: owner.String()})
	}
	return db.listListings(filter, cursor, count, backward)
}

func (db *MongoDbBridge) listListings(filter bson.D, cursor types.Cursor, count int, backward bool) (out *types.ListingList, err error) {
	var list types.ListingList
	col := db.client.Database(db.dbName).Collection(coListings)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.ListingSortingNone, backward)
	if err != nil {
		log.Errorf("error loading listings list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			log.Errorf("error closing listings list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.Listing
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the listing in list; %s", err.Error())
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
