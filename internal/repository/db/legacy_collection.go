package db

import (
	"artion-api-graphql/internal/types"
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
