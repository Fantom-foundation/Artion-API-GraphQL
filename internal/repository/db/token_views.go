package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/big"
	"strings"
)

const (
	// coLegacyTokens is the name of database collection.
	coLegacyTokens = "nftitems"

	fiLegacyTokenContract = "contractAddress"

	fiLegacyTokenTokenId = "tokenID"

	fiLegacyTokenViewed = "viewed"
)

func (sdb *SharedMongoDbBridge) GetTokenViews(contract common.Address, tokenId big.Int) (views *big.Int, err error) {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyTokens)

	filter := bson.D{
		{Key: fiLegacyTokenContract, Value: strings.ToLower(contract.String())},
		{Key: fiLegacyTokenTokenId, Value: int32(tokenId.Int64())},
	}
	result := col.FindOne(context.Background(), filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return big.NewInt(0), nil
		}
		return nil, err
	}

	var row types.LegacyToken
	if err = result.Decode(&row); err != nil {
		log.Errorf("can not decode LegacyToken; %s", err.Error())
		return nil, err
	}

	return big.NewInt(int64(row.Viewed)), err
}

func (sdb *SharedMongoDbBridge) IncrementTokenViews(contract common.Address, tokenId big.Int) error {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyTokens)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyTokenContract, Value: strings.ToLower(contract.String())},
			{Key: fiLegacyTokenTokenId, Value: int32(tokenId.Int64())},
		},
		bson.D{
			{Key: "$inc", Value: bson.D{{Key: fiLegacyTokenViewed, Value: 1 }}},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fiLegacyTokenContract, Value: strings.ToLower(contract.String())},
				{Key: fiLegacyTokenTokenId, Value: int32(tokenId.Int64())},
			}},
		},
		// TODO: uncomment when Artion 1.0 will be down
		//options.Update().SetUpsert(true),
	); err != nil {
		log.Errorf("can not update LegacyToken views; %s", err)
		return err
	}
	return nil
}
