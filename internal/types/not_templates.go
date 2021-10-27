// Package types provides high level structures for the API server.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// NotificationTemplate represents a notification template reference on the remote notification API
// to be used to build and deliver a notification of the given type and on the given collection.
type NotificationTemplate struct {
	Type           int32           `bson:"type"`
	Contract       *common.Address `bson:"contract"`
	TokenID        *hexutil.Big    `bson:"token"`
	Provider       string          `bson:"provider"`
	TemplateID     string          `bson:"template"`
	SenderName     string          `bson:"sender"`
	SenderID       string          `bson:"sender_id"`
	Subject        string          `bson:"subject"`
	Recipient      *string         `bson:"recipient"`
	ExtendedParams *string         `bson:"params"`
}
