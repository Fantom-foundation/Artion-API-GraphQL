// Package db provides access to the persistent storage.
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
)

const (
	// coNotifications is the name of the collection keeping event notifications.
	coNotifications = "notifications"

	// coNotificationTemplates is the name of the collection keeping event notification templates.
	coNotificationTemplates = "notification_tpl"
)

// StoreNotification stores the given notification in persistent storage.
// The function returns true if the notification was unique and didn't exist before.
func (db *MongoDbBridge) StoreNotification(no *types.Notification) (bool, error) {
	if no == nil {
		return false, fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coNotifications)

	// try to do the insert
	id := types.NotificationID(no)
	rs, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: fieldId, Value: id}},
		bson.D{
			{Key: "$set", Value: no},
			{Key: "$setOnInsert", Value: bson.D{{
				Key: fieldId, Value: id,
			}}},
		},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Errorf("can not store notification; %s", err.Error())
		return false, err
	}

	if rs.UpsertedCount > 0 {
		log.Infof("notification #%d on %s/%s to %s stored", no.Type, no.Contract.String(), no.TokenId.String(), no.Recipient.String())
	}
	return rs.UpsertedCount > 0, nil
}

// NotificationTemplates pulls a list of notification templates applicable to the
func (db *MongoDbBridge) NotificationTemplates(nt int32, contract *common.Address, tokenID *hexutil.Big) []types.NotificationTemplate {
	// get the collection
	col := db.client.Database(db.dbName).Collection(coNotificationTemplates)

	list := loadNotificationTemplates(col, notificationTemplateFilter(nt, contract, tokenID))
	if list != nil && len(list) > 0 {
		return list
	}

	// is there the token ID? try generic contract wide list
	if contract != nil && tokenID != nil {
		return db.NotificationTemplates(nt, contract, nil)
	}

	// is there the token ID but no token ID? try generic type only level list
	if contract != nil && tokenID == nil {
		return db.NotificationTemplates(nt, nil, nil)
	}
	return make([]types.NotificationTemplate, 0)
}

// notificationTemplateFilter generates BSON filter for the given set of params.
func notificationTemplateFilter(nt int32, contract *common.Address, tokenID *hexutil.Big) bson.D {
	// base filter for notification type
	filter := bson.D{{Key: "type", Value: nt}}

	// filter contract address set or nil
	if contract != nil {
		filter = append(filter, bson.E{Key: "contract", Value: *contract})
	} else {
		filter = append(filter, bson.E{Key: "contract", Value: bson.D{{Key: "$type", Value: 10}}})
	}

	// filter token ID set or nil
	if tokenID != nil {
		filter = append(filter, bson.E{Key: "token", Value: *tokenID})
	} else {
		filter = append(filter, bson.E{Key: "token", Value: bson.D{{Key: "$type", Value: 10}}})
	}

	log.Infof("loading notification templates for filter %#v", filter)
	return filter
}

// loadNotificationTemplates loads a list of notification templates
// form given collection based on provided filter.
func loadNotificationTemplates(col *mongo.Collection, filter bson.D) []types.NotificationTemplate {
	cur, err := col.Find(context.Background(), filter, options.Find().SetSort(bson.D{
		{Key: "provider", Value: 1},
		{Key: "template", Value: 1},
	}))
	if err != nil {
		log.Errorf("could not query templates; %s", err.Error())
		return nil
	}

	defer func(c *mongo.Cursor) {
		if err := c.Close(context.Background()); err != nil {
			log.Errorf("could not close search cursor; %s", err.Error())
		}
	}(cur)

	// collect templates
	templates := make([]types.NotificationTemplate, 0)

	// loop the data
	for cur.Next(context.Background()) {
		var row types.NotificationTemplate
		if err := cur.Decode(&row); err != nil {
			log.Errorf("could not decode notification template; %s", err.Error())
			return nil
		}

		templates = append(templates, row)
	}

	return templates
}
