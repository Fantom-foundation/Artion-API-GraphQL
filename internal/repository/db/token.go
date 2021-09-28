package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// initTokenCollection initializes collection with indexes and additional parameters.
func (db *MongoDbBridge) initTokenCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// indexes
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiTokenName, Value: 1}}})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for transaction collection; %s", err.Error())
	}

	// log we are done that
	db.log.Debugf("transactions collection initialized")
}

func (db *MongoDbBridge) StoreToken(token *types.Token) error {
	if token == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(types.CoTokens)

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), token); err != nil {
		db.log.Errorf("can not store Token; %s", err)
		return err
	}
	// make sure gas price collection is initialized
	if db.initTokens != nil {
		db.initTokens.Do(func() { db.initTokenCollection(col); db.initTokens = nil })
	}
	return nil
}

func (db *MongoDbBridge) GetToken(nft common.Address, tokenId hexutil.Big) (token *types.Token, err error) {
	col := db.client.Database(db.dbName).Collection(types.CoTokens)
	ctx := context.Background()
	filter := bson.D{
		{ Key: types.FiTokenNft, Value: bson.D{{Key: "$eq", Value: nft.String() }} },
		{ Key: types.FiTokenTokenId, Value: bson.D{{Key: "$eq", Value: tokenId.String() }} },
	}
	result := col.FindOne(ctx, filter)

	var row types.Token
	if err = result.Decode(&row); err != nil {
		db.log.Errorf("can not decode token; %s", err.Error())
		return nil, err
	}

	return &row, err
}

func (db *MongoDbBridge) ListTokens(cursor types.Cursor, count int, backward bool) (out *types.TokenList, err error) {
	filter := bson.D{}
	return db.listTokens(&filter, cursor, count, backward)
}

func (db *MongoDbBridge) listTokens(filter *bson.D, cursor types.Cursor, count int, backward bool) (out *types.TokenList, err error) {
	var list types.TokenList
	col := db.client.Database(db.dbName).Collection(types.CoTokens)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, backward)
	if err != nil {
		db.log.Errorf("error loading tokens list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing tokens list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.Token
			if err = ld.Decode(&row); err != nil {
				db.log.Errorf("can not decode the token in list; %s", err.Error())
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
