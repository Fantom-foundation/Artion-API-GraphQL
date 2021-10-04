package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
)

const (
	// CoTokens is the name of database collection.
	coTokens = "tokens"

	// fiTokenNft is the column storing the name of the NFT token.
	fiTokenNft = "nft"

	// FiTokenName is the column storing the name of the NFT token.
	fiTokenName = "name"
)

// initTokenCollection initializes collection with indexes and additional parameters.
func (db *MongoDbBridge) initTokenCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// indexes
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: fiTokenName, Value: 1}}})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		log.Panicf("can not create indexes for transaction collection; %s", err.Error())
	}

	// log we are done that
	log.Debugf("transactions collection initialized")
}

// StoreToken inserts new NFT token or updates existing token in persistent database.
func (db *MongoDbBridge) StoreToken(token *types.Token) error {
	if token == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coTokens)

	// try to do the insert
	rs, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: fieldId, Value: types.TokenIdFromAddress(&token.Nft, (*big.Int)(&token.TokenId))}},
		bson.D{{Key: "$set", Value: token}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Errorf("can not store token %s at %s; %s", token.TokenId.String(), token.Nft.String(), err.Error())
		return err
	}
	if rs.UpsertedCount > 0 {
		log.Infof("token %s on contract %s added to database", token.TokenId.String(), token.Nft.String())
	}

	// make sure gas price collection is initialized
	if db.initTokens != nil {
		db.initTokens.Do(func() { db.initTokenCollection(col); db.initTokens = nil })
	}
	return nil
}

// GetToken loads specific NFT token for the given contract address and token ID
func (db *MongoDbBridge) GetToken(nft *common.Address, tokenId *big.Int) (token *types.Token, err error) {
	col := db.client.Database(db.dbName).Collection(coTokens)
	result := col.FindOne(context.Background(), bson.D{{Key: fieldId, Value: types.TokenIdFromAddress(nft, tokenId)}})

	var row types.Token
	if err = result.Decode(&row); err != nil {
		log.Errorf("can not decode token; %s", err.Error())
		return nil, err
	}

	return &row, err
}

func (db *MongoDbBridge) ListTokens(cursor types.Cursor, count int, backward bool) (out *types.TokenList, err error) {
	filter := bson.D{}
	return db.listTokens(&filter, cursor, count, backward)
}

func (db *MongoDbBridge) ListCollectionTokens(collection common.Address, cursor types.Cursor, count int, backward bool) (out *types.TokenList, err error) {
	filter := bson.D{
		{ Key: fiTokenNft, Value: collection.String() },
	}
	return db.listTokens(&filter, cursor, count, backward)
}

func (db *MongoDbBridge) listTokens(filter *bson.D, cursor types.Cursor, count int, backward bool) (out *types.TokenList, err error) {
	var list types.TokenList
	col := db.client.Database(db.dbName).Collection(coTokens)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, backward)
	if err != nil {
		log.Errorf("error loading tokens list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			log.Errorf("error closing tokens list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.Token
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the token in list; %s", err.Error())
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
