package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// CoListings is the name of database collection.
	coListings = "listings"

	// FiListingNft is the name of the DB column storing NFT contract address.
	fiListingNft     = "nft"

	// FiListingTokenId represents the name of the DB column storing NFT token ID.
	fiListingTokenId = "tokenId"

	// FiListingOwner represents the name of the DB column storing token owner.
	fiListingOwner   = "owner"
)

// initListingCollection initializes collection with indexes and additional parameters.
func (db *MongoDbBridge) initListingCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index sender and recipient
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: fiListingNft, Value: 1}, {Key: fiListingTokenId, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: fiListingOwner, Value: 1}}})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		log.Panicf("can not create indexes for transaction collection; %s", err.Error())
	}

	// log we are done that
	log.Debugf("transactions collection initialized")
}

func (db *MongoDbBridge) AddListing(listing *types.Listing) error {
	if listing == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coListings)

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), listing); err != nil {
		log.Errorf("can not add Listing; %s", err)
		return err
	}

	if db.initListings != nil {
		db.initListings.Do(func() { db.initListingCollection(col); db.initListings = nil })
	}
	return nil
}

func (db *MongoDbBridge) UpdateListing(listing *types.Listing) error {
	if listing == nil {
		return fmt.Errorf("no value to store")
	}
	col := db.client.Database(db.dbName).Collection(coListings)

	filter := bson.D{
		{ Key: fiListingOwner, Value: listing.Owner.String() },
		{ Key: fiListingNft, Value: listing.Nft.String() },
		{ Key: fiListingTokenId, Value: listing.TokenId.String() },
	}

	if _, err := col.ReplaceOne(context.Background(), filter, listing); err != nil {
		log.Errorf("can not update Listing; %s", err)
		return err
	}
	return nil
}

func (db *MongoDbBridge) RemoveListing(owner common.Address, nft common.Address, tokenId hexutil.Big) error {
	col := db.client.Database(db.dbName).Collection(coListings)

	filter := bson.D{
		{ Key: fiListingOwner, Value: owner.String() },
		{ Key: fiListingNft, Value: nft.String() },
		{ Key: fiListingTokenId, Value: tokenId.String() },
	}

	if _, err := col.DeleteOne(context.Background(), filter); err != nil {
		log.Errorf("can not update Listing; %s", err)
		return err
	}
	return nil
}

func (db *MongoDbBridge) ListListings(nft *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (out *types.ListingList, err error) {
	filter := bson.D{}
	if nft != nil {
		filter = append(filter, primitive.E{ Key: fiListingNft, Value: nft.String() })
	}
	if tokenId != nil {
		filter = append(filter, primitive.E{ Key: fiListingTokenId, Value: tokenId.String() })
	}
	if owner != nil {
		filter = append(filter, primitive.E{ Key: fiListingOwner, Value: owner.String() })
	}
	return db.listListings(&filter, cursor, count, backward)
}

func (db *MongoDbBridge) listListings(filter *bson.D, cursor types.Cursor, count int, backward bool) (out *types.ListingList, err error) {
	var list types.ListingList
	col := db.client.Database(db.dbName).Collection(coListings)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, backward)
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
