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

	// fiListingClosed represents the name of the DB column storing date/time of listing having been closed.
	fiListingClosed = "closed"

	// fiListingStartTime represents the name of the DB column storing start time of the listing.
	fiListingStartTime = "start"

	// fiListingUnifiedPrice represents the name of the DB column storing listed price of the token in USD.
	fiListingUnifiedPrice = "uprice"

	// fiListingIsActive represents the name of the DB column storing if the listing creator currently own the token.
	fiListingIsActive = "is_active"
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

// SetListingActive sets IsActive state of listing when token ownership changes.
func (db *MongoDbBridge) SetListingActive(contract *common.Address, tokenID *big.Int, owner *common.Address, isActive bool) error {
	// get the collection
	col := db.client.Database(db.dbName).Collection(coListings)
	rs, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: fieldId, Value: types.ListingID(contract, tokenID, owner)}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: fiListingIsActive, Value: isActive},
		}}},
	)
	if err != nil {
		log.Errorf("can not update is_active of listing %s/%s; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err.Error())
		return err
	}
	if rs.UpsertedCount > 0 {
		log.Infof("listing %s/%s is_active updated", contract.String(), (*hexutil.Big)(tokenID).String())
	}
	return nil
}

// OpenListingSince pulls the earliest date of open listing for the token.
// If there is no open listing, it returns nil.
func (db *MongoDbBridge) OpenListingSince(contract *common.Address, tokenID *big.Int) *types.Time {
	var row struct {
		Since types.Time `bson:"val"`
	}

	col := db.client.Database(db.dbName).Collection(coListings)
	err := db.AggregateSingle(col, &mongo.Pipeline{
		{
			{Key: "$match", Value: bson.D{
				{Key: fiListingContract, Value: *contract},
				{Key: fiListingTokenId, Value: hexutil.Big(*tokenID)},
				{Key: fiListingClosed, Value: bson.D{{Key: "$type", Value: 10}}}, // closed = null - not sold yet
				{Key: fiListingIsActive, Value: bson.D{{Key: "$ne", Value: false}}}, // listing creator own the token
			}},
		},
		{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: nil},
				{Key: "val", Value: bson.D{{Key: "$min", Value: "$start"}}},
			}},
		},
	}, &row)

	if err != nil {
		// no listing at all?
		if err == mongo.ErrNoDocuments {
			log.Infof("no open listing available for %s/%s", contract.String(), (*hexutil.Big)(tokenID).String())
			return nil
		}
		log.Criticalf("failed earliest listing check of %s/%s; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err.Error())
		return nil
	}

	log.Infof("open listing on %s/%s since %s", contract.String(), (*hexutil.Big)(tokenID).String(), time.Time(row.Since).Format(time.RFC1123))
	return &row.Since
}

// MinListingPrice obtains minimal listed price of the token and time until when is the information valid.
func (db *MongoDbBridge) MinListingPrice(contract *common.Address, tokenID *big.Int) (tokenPrice types.TokenPrice, minListValid *types.Time) {
	col := db.client.Database(db.dbName).Collection(coListings)
	now := time.Now()

	// get minimal price starting before now
	var minListing types.Listing
	sr := col.FindOne(context.Background(), bson.D{
		{Key: fiListingContract, Value: *contract},
		{Key: fiListingTokenId, Value: hexutil.Big(*tokenID)},
		{Key: fiListingClosed, Value: bson.D{{Key: "$type", Value: 10}}}, // closed=null = not closed (sold) yet
		{Key: fiListingStartTime, Value: bson.D{{Key: "$lte", Value: now}}}, // already started
		{Key: fiListingIsActive, Value: bson.D{{Key: "$ne", Value: false}}}, // listing creator own the token
	}, options.FindOne().SetSort(bson.D{{Key: fiListingUnifiedPrice, Value: 1}}))
	if sr.Err() != nil {
		if sr.Err() != mongo.ErrNoDocuments {
			log.Errorf("error loading min listing price; %s", sr.Err())
		}
		return
	}
	err := sr.Decode(&minListing)
	if err != nil {
		log.Errorf("error decoding min listing price; %s", err)
		return
	}
	tokenPrice = types.TokenPrice{
		Usd: minListing.UnifiedPrice,
		Amount: minListing.UnitPrice,
		PayToken: minListing.PayToken,
	}

	// get end of validity - start of first future minimum (listing starting to be valid in future)
	var nextStartingListing types.Listing
	sr = col.FindOne(context.Background(), bson.D{
		{Key: fiListingContract, Value: *contract},
		{Key: fiListingTokenId, Value: hexutil.Big(*tokenID)},
		{Key: fiListingClosed, Value: bson.D{{Key: "$type", Value: 10}}}, // not closed yet
		{Key: fiListingStartTime, Value: bson.D{{Key: "$gt", Value: now}}}, // already started
		{Key: fiListingUnifiedPrice, Value: bson.D{{Key: "$lt", Value: minListing.UnifiedPrice}}},
	}, options.FindOne().SetSort(bson.D{{Key: fiListingStartTime, Value: 1}}))
	if sr.Err() != nil {
		if sr.Err() != mongo.ErrNoDocuments {
			log.Errorf("error loading nextStartingListing; %s", sr.Err())
		}
		return
	}
	err = sr.Decode(&nextStartingListing)
	if err != nil {
		log.Errorf("error decoding nextStartingListing; %s", err)
		return
	}
	minListValid = &nextStartingListing.StartTime
	return
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

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.ListingSortingCreated, !backward)
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
