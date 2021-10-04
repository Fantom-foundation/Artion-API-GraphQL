// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// coNFTCollection is the name of the collection of NFT contracts.
	coNFTCollection = "nft"

	// fiNFTAddress is the name of the field keeping the NFT contract address.
	fiNFTAddress = "_id"
)

// AddNFTCollection adds the specified NFT collection contract record.
func (db *MongoDbBridge) AddNFTCollection(nft *types.NFTCollection) error {
	col := db.client.Database(db.dbName).Collection(coNFTCollection)
	if db.isNFTCollectionKnown(col, nft) {
		return nil
	}

	if _, err := col.InsertOne(context.Background(), nft); err != nil {
		log.Errorf("can not add NFT collection %s", nft.Address.String(), err.Error())
		return err
	}
	return nil
}

// isNFTCollectionKnown checks if the given NFT collection is already stored in the database.
func (db *MongoDbBridge) isNFTCollectionKnown(col *mongo.Collection, nft *types.NFTCollection) bool {
	return db.exists(col, &bson.D{{Key: fiNFTAddress, Value: nft.Address.String()}})
}

func (db *MongoDbBridge) GetNFTCollection(address common.Address) (NFTCollection *types.NFTCollection, err error) {
	col := db.client.Database(db.dbName).Collection(coNFTCollection)
	ctx := context.Background()
	filter := bson.D{
		{ Key: fiNFTAddress, Value: bson.D{{Key: "$eq", Value: address.String() }} },
	}
	result := col.FindOne(ctx, filter)

	var row types.NFTCollection
	if err = result.Decode(&row); err != nil {
		log.Errorf("can not decode NFTCollection; %s", err.Error())
		return nil, err
	}

	return &row, err
}

func (db *MongoDbBridge) ListNFTCollections(cursor types.Cursor, count int, backward bool) (out *types.NFTCollectionList, err error) {
	filter := bson.D{}
	return db.listNFTCollections(&filter, cursor, count, backward)
}

func (db *MongoDbBridge) listNFTCollections(filter *bson.D, cursor types.Cursor, count int, backward bool) (out *types.NFTCollectionList, err error) {
	var list types.NFTCollectionList
	col := db.client.Database(db.dbName).Collection(coNFTCollection)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, backward)
	if err != nil {
		log.Errorf("error loading NFTCollections list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			log.Errorf("error closing NFTCollections list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.NFTCollection
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the NFTCollection in list; %s", err.Error())
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
