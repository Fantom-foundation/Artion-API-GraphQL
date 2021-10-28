package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/big"
	"strings"
)

const (
	// coTokenLikes is the name of database collection.
	coTokenLikes = "likes"

	// fiTokenLikeContract is the column storing the address of the NFT token contract.
	fiTokenLikeContract = "contractAddress"

	// fiTokenLikeToken is the column storing the token ID.
	fiTokenLikeToken = "tokenID"

	// fiTokenLikeUser is the column storing the user.
	fiTokenLikeUser = "follower"
)

func (sdb *SharedMongoDbBridge) AddTokenLike(tokenLike *types.TokenLike) error {
	if tokenLike == nil {
		return fmt.Errorf("no value to store")
	}

	col := sdb.client.Database(sdb.dbName).Collection(coTokenLikes)
	_, err := col.InsertOne(context.Background(), bson.D{
		{Key: fiTokenLikeUser, Value: strings.ToLower(tokenLike.User.String())},
		{Key: fiTokenLikeContract, Value: strings.ToLower(tokenLike.Contract.String())},
		{Key: fiTokenLikeToken, Value: tokenLike.TokenId32},
	})
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil // ignore already-exists error
		}
		log.Errorf("can not add token like; %s", err.Error())
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) RemoveTokenLike(tokenLike *types.TokenLike) error {
	if tokenLike == nil {
		return fmt.Errorf("no value to store")
	}

	col := sdb.client.Database(sdb.dbName).Collection(coTokenLikes)
	_, err := col.DeleteOne(
		context.Background(),
		bson.D{
			{Key: fiTokenLikeUser, Value: strings.ToLower(tokenLike.User.String())},
			{Key: fiTokenLikeContract, Value: strings.ToLower(tokenLike.Contract.String())},
			{Key: fiTokenLikeToken, Value: tokenLike.TokenId32},
		},
	)
	if err != nil {
		log.Errorf("can not remove token like; %s", err.Error())
		return err
	}
	return nil
}

// GetTokenLikesCount get amount of likes for given token
func (sdb *SharedMongoDbBridge) GetTokenLikesCount(contract *common.Address, tokenId *big.Int) (count int64, err error) {
	col := sdb.client.Database(sdb.dbName).Collection(coTokenLikes)
	return col.CountDocuments(context.Background(), bson.D{
		{Key: fiTokenLikeContract, Value: strings.ToLower(contract.String()) },
		{Key: fiTokenLikeToken, Value: int32(tokenId.Int64()) },
	})
}

// IsTokenLiked get like status of the token for given user
func (sdb *SharedMongoDbBridge) IsTokenLiked(user *common.Address, contract *common.Address, tokenId *big.Int) (bool, error) {
	col := sdb.client.Database(sdb.dbName).Collection(coTokenLikes)
	count, err := col.CountDocuments(context.Background(), bson.D{
		{Key: fiTokenLikeUser, Value: strings.ToLower(user.String())},
		{Key: fiTokenLikeContract, Value: strings.ToLower(contract.String()) },
		{Key: fiTokenLikeToken, Value: int32(tokenId.Int64()) },
	})
	return count != 0, err
}

func (sdb *SharedMongoDbBridge) ListUserTokenLikes(user *common.Address, cursor types.Cursor, count int, backward bool) (out *types.TokenLikeList, err error) {
	filter := bson.D{ {Key: fiTokenLikeUser, Value: strings.ToLower(user.String())} }
	return sdb.listTokenLikes(filter, cursor, count, backward)
}

func (sdb *SharedMongoDbBridge) listTokenLikes(filter bson.D, cursor types.Cursor, count int, backward bool) (out *types.TokenLikeList, err error) {
	db := (*MongoDbBridge)(sdb)
	var list types.TokenLikeList
	col := db.client.Database(db.dbName).Collection(coTokenLikes)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.TokenLikeSortingNone, backward)
	if err != nil {
		log.Errorf("error loading token likes list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			log.Errorf("error closing token like list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.TokenLike
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the token like in list; %s", err.Error())
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
