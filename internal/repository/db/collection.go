// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// coCollection is the name of the collection of NFT contracts.
	coCollection = "collections"

	// fiCollectionAddress is the name of the field keeping the NFT contract address.
	fiCollectionAddress = "_id"
)

// AddCollection adds the specified NFT collection contract record.
func (db *MongoDbBridge) AddCollection(nft *types.Collection) error {
	col := db.client.Database(db.dbName).Collection(coCollection)
	if db.isCollectionKnown(col, nft) {
		return nil
	}

	if _, err := col.InsertOne(context.Background(), nft); err != nil {
		log.Errorf("can not add NFT collection %s", nft.Address.String(), err.Error())
		return err
	}
	return nil
}

// isCollectionKnown checks if the given NFT collection is already stored in the database.
func (db *MongoDbBridge) isCollectionKnown(col *mongo.Collection, nft *types.Collection) bool {
	return db.exists(col, &bson.D{{Key: fiCollectionAddress, Value: nft.Address.String()}})
}

func (db *MongoDbBridge) GetCollection(address common.Address) (NFTCollection *types.Collection, err error) {
	col := db.client.Database(db.dbName).Collection(coCollection)
	ctx := context.Background()
	filter := bson.D{
		{ Key: fiCollectionAddress, Value: bson.D{{Key: "$eq", Value: address.String() }} },
	}
	result := col.FindOne(ctx, filter)

	var row types.Collection
	if err = result.Decode(&row); err != nil {
		log.Errorf("can not decode Collection; %s", err.Error())
		return nil, err
	}

	return &row, err
}

func (db *MongoDbBridge) ListCollections(cursor types.Cursor, count int, backward bool) (out *types.CollectionList, err error) {
	filter := bson.D{}
	return db.listCollections(&filter, cursor, count, backward)
}

func (db *MongoDbBridge) listCollections(filter *bson.D, cursor types.Cursor, count int, backward bool) (out *types.CollectionList, err error) {
	var list types.CollectionList
	col := db.client.Database(db.dbName).Collection(coCollection)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.CollectionSortingNone, backward)
	if err != nil {
		log.Errorf("error loading Collections list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			log.Errorf("error closing Collections list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.Collection
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the Collection in list; %s", err.Error())
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
