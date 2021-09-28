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

// initListingCollection initializes collection with indexes and additional parameters.
func (db *MongoDbBridge) initListingCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index sender and recipient
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiListingNft, Value: 1}, {Key: types.FiListingTokenId, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiListingOwner, Value: 1}}})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for transaction collection; %s", err.Error())
	}

	// log we are done that
	db.log.Debugf("transactions collection initialized")
}

func (db *MongoDbBridge) AddListing(listing *types.Listing) error {
	if listing == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(types.CoListings)

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), listing); err != nil {
		db.log.Errorf("can not add Listing; %s", err)
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
	col := db.client.Database(db.dbName).Collection(types.CoListings)

	filter := bson.D{
		{ Key: types.FiListingOwner, Value: listing.Owner.String() },
		{ Key: types.FiListingNft, Value: listing.Nft.String() },
		{ Key: types.FiListingTokenId, Value: listing.TokenId.String() },
	}

	if _, err := col.ReplaceOne(context.Background(), filter, listing); err != nil {
		db.log.Errorf("can not update Listing; %s", err)
		return err
	}
	return nil
}

func (db *MongoDbBridge) RemoveListing(owner common.Address, nft common.Address, tokenId hexutil.Big) error {
	col := db.client.Database(db.dbName).Collection(types.CoListings)

	filter := bson.D{
		{ Key: types.FiListingOwner, Value: owner.String() },
		{ Key: types.FiListingNft, Value: nft.String() },
		{ Key: types.FiListingTokenId, Value: tokenId.String() },
	}

	if _, err := col.DeleteOne(context.Background(), filter); err != nil {
		db.log.Errorf("can not update Listing; %s", err)
		return err
	}
	return nil
}

func (db *MongoDbBridge) ListListings(nft *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (out *types.ListingList, err error) {
	filter := bson.D{}
	if nft != nil {
		filter = append(filter, primitive.E{ Key: types.FiListingNft, Value: nft.String() })
	}
	if tokenId != nil {
		filter = append(filter, primitive.E{ Key: types.FiListingTokenId, Value: tokenId.String() })
	}
	if owner != nil {
		filter = append(filter, primitive.E{ Key: types.FiListingOwner, Value: owner.String() })
	}
	return db.listListings(&filter, cursor, count, backward)
}

func (db *MongoDbBridge) listListings(filter *bson.D, cursor types.Cursor, count int, backward bool) (out *types.ListingList, err error) {
	var list types.ListingList
	col := db.client.Database(db.dbName).Collection(types.CoListings)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, backward)
	if err != nil {
		db.log.Errorf("error loading listings list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing listings list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.Listing
			if err = ld.Decode(&row); err != nil {
				db.log.Errorf("can not decode the listing in list; %s", err.Error())
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
