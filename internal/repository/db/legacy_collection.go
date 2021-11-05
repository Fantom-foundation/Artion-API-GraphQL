package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

const (
	// coCollection is the name of the collection of NFT contracts in shared database.
	coLegacyCollection = "collections"

	// fiCollectionAddress is the name of the field keeping the NFT contract address.
	fiLegacyCollectionAddress = "erc721Address"

	// fiLegacyCollectionIsAppropriate is the name of field keeping status if the NFT contract.
	fiLegacyCollectionIsAppropriate = "isAppropriate"
)

func (sdb *SharedMongoDbBridge) GetLegacyCollection(address common.Address) (collection *types.LegacyCollection, err error) {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)
	ctx := context.Background()
	filter := bson.D{
		{ Key: fiLegacyCollectionAddress, Value: bson.D{{Key: "$eq", Value: strings.ToLower(address.String()) }} },
	}
	result := col.FindOne(ctx, filter)

	var row types.LegacyCollection
	if err = result.Decode(&row); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Errorf("can not decode LegacyCollection; %s", err.Error())
		return nil, err
	}

	return &row, err
}

func (sdb *SharedMongoDbBridge) ListLegacyCollections(cursor types.Cursor, count int, backward bool) (out *types.LegacyCollectionList, err error) {
	db := (*MongoDbBridge)(sdb)
	var list types.LegacyCollectionList
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)
	ctx := context.Background()

	filter := bson.D{
		{Key: fiLegacyCollectionIsAppropriate, Value: bson.D{{Key: "$eq", Value: true}}},
	}

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.LegacyCollectionSortingName, backward)
	if err != nil {
		log.Errorf("error loading LegacyCollections list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			log.Errorf("error closing LegacyCollections list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.LegacyCollection
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the LegacyCollection in list; %s", err.Error())
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
