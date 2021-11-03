package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

const (
	// coNotificationSettings is the name of database collection.
	coNotificationSettings = "notificationsettings"

	// fiNotifSettUser is the column storing the address of the user.
	fiNotifSettUser = "address"

	fiNotifSettSNotification = "sNotification"
	fiNotifSettSBundleBuy = "sBundleBuy"
	fiNotifSettSBundleSell = "sBundleSell"
	fiNotifSettSBundleOffer = "sBundleOffer"
	fiNotifSettSBundleOfferCancel = "sBundleOfferCancel"
	fiNotifSettSNftAuctionPrice = "sNftAuctionPrice"
	fiNotifSettSNftBidToAuction = "sNftBidToAuction"
	fiNotifSettSNftBidToAuctionCancel = "sNftBidToAuctionCancel"
	fiNotifSettSAuctionWin = "sAuctionWin"
	fiNotifSettSAuctionOfBidCancel = "sAuctionOfBidCancel"
	fiNotifSettSNftSell = "sNftSell"
	fiNotifSettSNftBuy = "sNftBuy"
	fiNotifSettSNftOffer = "sNftOffer"
	fiNotifSettSNftOfferCancel = "sNftOfferCancel"

	fiNotifSettFNotification = "fNotification"
	fiNotifSettFBundleCreation = "fBundleCreation"
	fiNotifSettFBundleList = "fBundleList"
	fiNotifSettFBundlePrice = "fBundlePrice"
	fiNotifSettFNftAuctionPrice = "fNftAuctionPrice"
	fiNotifSettFNftList = "fNftList"
	fiNotifSettFNftAuction = "fNftAuction"
	fiNotifSettFNftPrice = "fNftPrice"
)

// GetNotificationSettings provides notification settings of the user, if available.
func (sdb *SharedMongoDbBridge) GetNotificationSettings(user common.Address) (settings *types.NotificationSettings, err error) {
	col := sdb.client.Database(sdb.dbName).Collection(coNotificationSettings)

	filter := bson.D{{Key: fiNotifSettUser, Value: strings.ToLower(user.String())}}
	result := col.FindOne(context.Background(), filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var row types.NotificationSettings
	if err = result.Decode(&row); err != nil {
		log.Errorf("can not decode notification settings; %s", err.Error())
		return nil, err
	}

	return &row, err
}

// StoreNotificationSettings inserts or updates existing notification settings of the user.
func (sdb *SharedMongoDbBridge) StoreNotificationSettings(user common.Address, settings types.NotificationSettings) error {
	col := sdb.client.Database(sdb.dbName).Collection(coNotificationSettings)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: fiNotifSettUser, Value: strings.ToLower(user.String())}},
		bson.D{
			{Key: "$set", Value: settings},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fiNotifSettUser, Value: strings.ToLower(user.String())},
			}},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		log.Errorf("can not update notification settings; %s", err)
		return err
	}
	return nil
}
