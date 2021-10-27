// Package types provides high level structures for the API server.
package types

import (
	"crypto/sha256"
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	// NotifyBurnedNFTToken represents a notification type for NFT burn event.
	NotifyBurnedNFTToken = iota
)

// Notification represents an event notification structure.
type Notification struct {
	Type       int32           `bson:"type" json:"type"`
	Contract   common.Address  `bson:"contract" json:"contract"`
	TokenId    hexutil.Big     `bson:"token" json:"token"`
	TimeStamp  Time            `bson:"fired" json:"fired"`
	Recipient  common.Address  `bson:"to" json:"to"`
	Originator *common.Address `bson:"from" json:"from"`
}

// NotificationID generates unique identifier for the notification record.
func NotificationID(no *Notification) primitive.ObjectID {
	hash := sha256.New()

	var id [12]byte
	binary.BigEndian.PutUint32(id[:4], uint32(no.Type))
	hash.Write(id[:4])

	hash.Write(no.Contract.Bytes())
	hash.Write(no.TokenId.ToInt().Bytes())
	hash.Write(no.Recipient.Bytes())

	if nil != no.Originator {
		hash.Write(no.Originator.Bytes())
	}

	hash.Write(([]byte)(time.Time(no.TimeStamp).String()))

	copy(id[:], hash.Sum(nil))
	return id
}
