package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

const (
	// coCollection is the name of the collection of NFT contracts in shared database.
	coLegacyCollection = "collections"

	// fiCollectionAddress is the name of the field keeping the NFT contract address.
	fiLegacyCollectionAddress = "erc721Address"

	fiLegacyCollectionName = "collectionName"
	fiLegacyCollectionDescription = "description"
	fiLegacyCollectionCategoriesStr = "categories"
	fiLegacyCollectionImage = "logoImageHash"
	fiLegacyCollectionOwner = "owner"
	fiLegacyCollectionFeeRecipient = "feeRecipient"
	fiLegacyCollectionRoyaltyValue = "royalty"
	fiLegacyCollectionDiscord = "discord"
	fiLegacyCollectionEmail = "email"
	fiLegacyCollectionTelegram = "telegram"
	fiLegacyCollectionSiteUrl = "siteUrl"
	fiLegacyCollectionMediumHandle = "mediumHandle"
	fiLegacyCollectionTwitterHandle = "twitterHandle"
	fiLegacyCollectionIsAppropriate = "isAppropriate"
	fiLegacyCollectionIsInternal  = "isInternal"
	fiLegacyCollectionIsOwnerOnly = "isOwnerble"
	fiLegacyCollectionIsVerified  = "isVerified"
	fiLegacyCollectionIsReviewed = "status"
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

// InsertLegacyCollection inserts collection record.
func (sdb *SharedMongoDbBridge) InsertLegacyCollection(c *types.LegacyCollection) error {
	if c == nil {
		return fmt.Errorf("no value to store")
	}
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)

	if _, err := col.InsertOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyCollectionAddress, Value: strings.ToLower(c.Address.String())},
			{Key: fiLegacyCollectionName, Value: c.Name},
			{Key: fiLegacyCollectionDescription, Value: c.Description},
			{Key: fiLegacyCollectionCategoriesStr, Value: c.CategoriesStr},
			{Key: fiLegacyCollectionImage, Value: c.Image},
			{Key: fiLegacyCollectionOwner, Value: strings.ToLower(c.Owner.String())},
			{Key: fiLegacyCollectionFeeRecipient, Value: strings.ToLower(c.FeeRecipient.String())},
			{Key: fiLegacyCollectionRoyaltyValue, Value: c.RoyaltyValue},
			{Key: fiLegacyCollectionDiscord, Value: c.Discord},
			{Key: fiLegacyCollectionEmail, Value: c.Email},
			{Key: fiLegacyCollectionTelegram, Value: c.Telegram},
			{Key: fiLegacyCollectionSiteUrl, Value: c.SiteUrl},
			{Key: fiLegacyCollectionMediumHandle, Value: c.MediumHandle},
			{Key: fiLegacyCollectionTwitterHandle, Value: c.TwitterHandle},
			{Key: fiLegacyCollectionIsAppropriate, Value: c.IsAppropriate},
			{Key: fiLegacyCollectionIsInternal, Value: c.IsInternal},
			{Key: fiLegacyCollectionIsOwnerOnly, Value: c.IsOwnerOnly},
			{Key: fiLegacyCollectionIsVerified, Value: c.IsVerified},
			{Key: fiLegacyCollectionIsReviewed, Value: c.IsReviewed},
		},
	); err != nil {
		log.Errorf("can not insert LegacyCollection; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) ListLegacyCollections(search *string, mintableBy *common.Address, cursor types.Cursor, count int, backward bool) (out *types.LegacyCollectionList, err error) {
	db := (*MongoDbBridge)(sdb)
	var list types.LegacyCollectionList
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)
	ctx := context.Background()

	filter := bson.D{
		{Key: fiLegacyCollectionIsAppropriate, Value: true},
	}
	if search != nil {
		filter = append(filter, bson.E{Key: fiLegacyCollectionName, Value: primitive.Regex{
			Pattern: *search,
			Options: "i",
		}})
	}
	if mintableBy != nil {
		filter = append(filter, bson.E{Key: fiLegacyCollectionIsInternal, Value: true})
		filter = append(filter, bson.E{Key: "$or", Value: bson.A{
			bson.D{{Key: fiLegacyCollectionIsOwnerOnly, Value: false }},
			bson.D{{Key: fiLegacyCollectionOwner, Value: *mintableBy }},
		}})
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
	defer closeFindCursor("LegacyCollections", ld)
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
