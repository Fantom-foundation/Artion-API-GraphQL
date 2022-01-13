package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

const (
	// coBannedNfts is the name of database collection.
	coBannedNfts = "bannednfts"

	fiBannedNftContract = "contractAddress"

	fiBannedNftTokenId = "tokenID"

	fiBannedNftIsBanned = "banned"

	fiBannedNftUpdated = "updatedAt"
)

// BanNft inserts NFT into list of banned NFTs.
func (sdb *SharedMongoDbBridge) BanNft(contract *common.Address, tokenId *hexutil.Big) error {
	col := sdb.client.Database(sdb.dbName).Collection(coBannedNfts)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: fiBannedNftContract, Value: strings.ToLower(contract.String())},
			{Key: fiBannedNftTokenId, Value: tokenId.String()},
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: fiBannedNftIsBanned, Value: true},
				{Key: fiBannedNftUpdated, Value: time.Now()},
			}},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fiBannedNftContract, Value: strings.ToLower(contract.String())},
				{Key: fiBannedNftTokenId, Value: tokenId.String()},
			}},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		log.Errorf("can not ban NFT; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) UnbanNft(contract *common.Address, tokenId *hexutil.Big) error {
	col := sdb.client.Database(sdb.dbName).Collection(coBannedNfts)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: fiBannedNftContract, Value: strings.ToLower(contract.String())},
			{Key: fiBannedNftTokenId, Value: tokenId.String()},
			{Key: fiBannedNftIsBanned, Value: true},
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: fiBannedNftIsBanned, Value: false},
				{Key: fiBannedNftUpdated, Value: time.Now()},
			}},
		},
	); err != nil {
		log.Errorf("can not unban NFT; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) ListBannedNftsAfter(after time.Time, maxAmount int64) (out []*types.BannedNft, err error) {
	db := (*MongoDbBridge)(sdb)
	list := make([]*types.BannedNft, maxAmount)
	col := db.client.Database(db.dbName).Collection(coBannedNfts)

	cur, err := col.Find(
		context.Background(),
		bson.D{{Key: fiBannedNftUpdated, Value: bson.D{{"$gte", after}}}},
		options.Find().SetSort(bson.D{{Key: fiBannedNftUpdated, Value: 1}}).SetLimit(maxAmount),
	)
	if err != nil {
		log.Errorf("can not list of newly banned nfts; %s", err.Error())
		return nil, err
	}
	defer closeFindCursor("BannedNfts", cur)

	var i int
	for cur.Next(context.Background()) {
		var row types.BannedNft
		if err := cur.Decode(&row); err != nil {
			log.Errorf("can not decode BannedNft; %s", err.Error())
			return nil, err
		}
		list[i] = &row
		i++
	}
	return list[:i], nil
}

func (sdb *SharedMongoDbBridge) ListBannedNfts(cursor types.Cursor, count int, backward bool) (out *types.BannedNftList, err error) {
	db := (*MongoDbBridge)(sdb)
	var list types.BannedNftList
	col := sdb.client.Database(sdb.dbName).Collection(coBannedNfts)
	ctx := context.Background()

	filter := bson.D{
		{Key: fiBannedNftIsBanned, Value: true},
	}

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return &list, nil // interested in TotalCount only
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.BannedNftSortingUpdated, !backward)
	if err != nil {
		log.Errorf("error loading BannedNft list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer closeFindCursor("BannedNft", ld)
	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.BannedNft
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the BannedNft in list; %s", err.Error())
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
