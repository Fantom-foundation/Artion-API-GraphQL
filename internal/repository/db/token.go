package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
	"time"
)

const (
	// CoTokens is the name of database collection.
	coTokens = "tokens"

	// fiTokenContract is the column storing the address of the NFT token contract.
	fiTokenContract = "contract"

	// fiTokenMetadataURI is the column storing the NFT token metadata URI.
	fiTokenMetadataURI = "uri"

	// FiTokenName is the column storing the name of the NFT token.
	fiTokenName = "name"

	// fiTokenDescription is the column storing the description of the NFT token.
	fiTokenDescription = "desc"

	// fiTokenImageURI is the column storing the image URI of the NFT token.
	fiTokenImageURI = "image"

	// fiTokenMetadataUpdate is the column storing the time
	// of the metadata update schedule of the NFT token.
	fiTokenMetadataUpdate = "meta_update"

	// fiTokenMetadataUpdate is the column storing the time
	// of the metadata update schedule of the NFT token.
	fiTokenMetadataUpdateFailures = "meta_failures"

	// fiTokenIsListed is the column marking listed token.
	fiTokenIsListed = "is_listed"

	// fiTokenListedPrice is the column storing the latest price token is listed at.
	fiTokenListedPrice = "amo_list"

	// fiTokenTradePrice is the column storing the latest price token was sold for.
	fiTokenTradePrice = "amo_trade"

	// fiTokenLastTrade is the column storing the latest trade date/time.
	fiTokenLastTrade = "last_trade"

	// fiTokenLastListed is the column storing the latest listing date/time.
	fiTokenLastListed = "last_list"
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

// TokenStore inserts new NFT token or updates existing token in persistent database.
func (db *MongoDbBridge) TokenStore(token *types.Token) error {
	if token == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coTokens)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	// try to do the insert
	id := token.ID()
	rs, err := col.UpdateOne(
		ctx,
		bson.D{{Key: fieldId, Value: id}},
		bson.D{
			{Key: "$set", Value: token},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fieldId, Value: id},
			}},
		},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Errorf("can not store token %s at %s; %s", token.TokenId.String(), token.Contract.String(), err.Error())
		return err
	}
	if rs.UpsertedCount > 0 {
		log.Infof("token %s on contract %s added to database", token.TokenId.String(), token.Contract.String())
	}

	// make sure gas price collection is initialized
	if db.initTokens != nil {
		db.initTokens.Do(func() { db.initTokenCollection(col); db.initTokens = nil })
	}
	return nil
}

// TokenUpdate updates the token data i the database using provided update data set.
func (db *MongoDbBridge) TokenUpdate(contract *common.Address, tokenID *big.Int, data bson.D) error {
	// get the collection
	col := db.client.Database(db.dbName).Collection(coTokens)
	rs, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: fieldId, Value: types.TokenIdFromAddress(contract, tokenID)}},
		bson.D{{Key: "$set", Value: data}},
	)
	if err != nil {
		log.Errorf("can not update token %s/%s; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err.Error())
		return err
	}
	if rs.UpsertedCount > 0 {
		log.Infof("token %s/%s updated", contract.String(), (*hexutil.Big)(tokenID).String())
	}
	return nil
}

// TokenUpdateMetadata updates basic metadata of the NFT token.
func (db *MongoDbBridge) TokenUpdateMetadata(nft *types.Token) error {
	if nft == nil {
		return fmt.Errorf("no value to store")
	}

	return db.TokenUpdate(&nft.Contract, (*big.Int)(&nft.TokenId), bson.D{
		{Key: fiTokenName, Value: nft.Name},
		{Key: fiTokenDescription, Value: nft.Description},
		{Key: fiTokenImageURI, Value: nft.ImageURI},
		{Key: fiTokenMetadataUpdate, Value: nft.MetaUpdate},
		{Key: fiTokenMetadataUpdateFailures, Value: nft.MetaFailures},
	})
}

// TokenUpdateMetadataRefreshSchedule sets the NFT metadata update schedule time.
func (db *MongoDbBridge) TokenUpdateMetadataRefreshSchedule(nft *types.Token) error {
	if nft == nil {
		return fmt.Errorf("no value to store")
	}

	return db.TokenUpdate(&nft.Contract, (*big.Int)(&nft.TokenId), bson.D{
		{Key: fiTokenMetadataUpdate, Value: nft.MetaUpdate},
		{Key: fiTokenMetadataUpdateFailures, Value: nft.MetaFailures},
	})
}

// TokenMarkListed marks the given NFT as listed for direct sale for the given price.
func (db *MongoDbBridge) TokenMarkListed(contract *common.Address, tokenID *big.Int, price *hexutil.Big, ts *time.Time) error {
	return db.TokenUpdate(contract, tokenID, bson.D{
		{Key: fiTokenLastListed, Value: *ts},
		{Key: fiTokenIsListed, Value: true},
		{Key: fiTokenListedPrice, Value: *price},
	})
}

// TokenMarkUnlisted marks the given NFT as listed for direct sale for the given price.
func (db *MongoDbBridge) TokenMarkUnlisted(contract *common.Address, tokenID *big.Int) error {
	return db.TokenUpdate(contract, tokenID, bson.D{
		{Key: fiTokenIsListed, Value: false},
	})
}

// TokenMarkSold marks the given NFT as sold for the given price.
func (db *MongoDbBridge) TokenMarkSold(contract *common.Address, tokenID *big.Int, price *hexutil.Big, ts *time.Time) error {
	return db.TokenUpdate(contract, tokenID, bson.D{
		{Key: fiTokenLastTrade, Value: *ts},
		{Key: fiTokenIsListed, Value: false},
		{Key: fiTokenTradePrice, Value: *price},
	})
}

// TokenGet loads specific NFT token for the given contract address and token ID
func (db *MongoDbBridge) TokenGet(nft *common.Address, tokenId *big.Int) (token *types.Token, err error) {
	col := db.client.Database(db.dbName).Collection(coTokens)
	result := col.FindOne(context.Background(), bson.D{{Key: fieldId, Value: types.TokenIdFromAddress(nft, tokenId)}})

	var row types.Token
	if err = result.Decode(&row); err != nil {
		log.Errorf("can not decode token; %s", err.Error())
		return nil, err
	}
	return &row, err
}

// TokenMetadataRefreshSet pulls s set of NFT tokens scheduled to be updated up to this time.
func (db *MongoDbBridge) TokenMetadataRefreshSet() ([]*types.Token, error) {
	list := make([]*types.Token, types.MetadataRefreshSetSize)
	col := db.client.Database(db.dbName).Collection(coTokens)

	// load the set from database
	cur, err := col.Find(
		context.Background(),
		bson.D{
			{Key: fiTokenMetadataUpdate, Value: bson.D{{"$lt", time.Now()}}},
			{Key: fiTokenMetadataURI, Value: bson.D{{"$ne", ""}}},
		},
		options.Find().SetSort(bson.D{{Key: fiTokenMetadataUpdate, Value: -1}}).SetLimit(types.MetadataRefreshSetSize),
	)
	if err != nil {
		log.Errorf("can not pull metadata refresh set; %s", err.Error())
		return nil, err
	}
	defer func() {
		if err := cur.Close(context.Background()); err != nil {
			log.Errorf("can not close cursor; %s", err.Error())
		}
	}()

	var i int
	for cur.Next(context.Background()) {
		var row types.Token
		if err := cur.Decode(&row); err != nil {
			log.Errorf("can not decode Token; %s", err.Error())
			return nil, err
		}
		list[i] = &row
		i++
	}
	return list[:i], nil
}

func (db *MongoDbBridge) ListTokens(cursor types.Cursor, count int, backward bool) (out *types.TokenList, err error) {
	filter := bson.D{}
	return db.listTokens(&filter, cursor, count, backward)
}

func (db *MongoDbBridge) ListCollectionTokens(collection common.Address, cursor types.Cursor, count int, backward bool) (out *types.TokenList, err error) {
	filter := bson.D{
		{Key: fiTokenContract, Value: collection.String()},
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
