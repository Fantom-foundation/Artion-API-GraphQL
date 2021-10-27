package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/big"
	"strings"
)

const (
	// coUnlockables is the name of database collection.
	coUnlockables = "unlockablecontents"

	fiUnlockableContract = "contractAddress"

	fiUnlockableTokenId = "tokenID"

	fiUnlockableContent = "content"
)

func (sdb *SharedMongoDbBridge) GetUnlockable(contract common.Address, tokenId big.Int) (unlockable *types.Unlockable, err error) {
	col := sdb.client.Database(sdb.dbName).Collection(coUnlockables)

	filter := bson.D{
		{ Key: fiUnlockableContract, Value: strings.ToLower(contract.String()) },
		{ Key: fiUnlockableTokenId, Value: int32(tokenId.Int64()) },
	}
	result := col.FindOne(context.Background(), filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var row types.Unlockable
	if err = result.Decode(&row); err != nil {
		log.Errorf("can not decode unlockable; %s", err.Error())
		return nil, err
	}

	return &row, err
}

func (sdb *SharedMongoDbBridge) InsertUnlockable(unlockable *types.Unlockable) error {
	if unlockable == nil {
		return fmt.Errorf("no value to store")
	}
	col := sdb.client.Database(sdb.dbName).Collection(coUnlockables)

	if _, err := col.InsertOne(
		context.Background(),
		bson.D{
			{Key: fiUnlockableContract, Value: strings.ToLower(unlockable.Contract.String())},
			{Key: fiUnlockableTokenId, Value: unlockable.TokenId},
			{Key: fiUnlockableContent, Value: unlockable.Content},
		},
	); err != nil {
		log.Errorf("can not insert unlockable; %s", err)
		return err
	}
	return nil
}
