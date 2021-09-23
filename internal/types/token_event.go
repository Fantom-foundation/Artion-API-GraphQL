package types

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenEventType int8

const (
	// CoTokenEvents is the name of database collection.
	CoTokenEvents = "token_events"

	// BSON attributes names used in the database collection.
	FiTokenEventId      = "_id"
	FiTokenEventNft     = "nft"
	FiTokenEventTokenId = "tokenId"
	FiTokenEventSeller  = "seller"
	FiTokenEventBuyer   = "buyer"
	FiTokenEventTime    = "evtTime"
)

const (
	EvtTpUnknown TokenEventType = iota
	EvtTpItemListed
	EvtTpItemUpdated
	EvtTpItemCanceled
	EvtTpItemSold
	EvtTpOfferCreated
	EvtTpOfferCanceled
)

// TokenEvent represents marketplace related events on tokens - when they are sold etc.
type TokenEvent struct {
	Id            primitive.ObjectID `bson:"_id"`
	EventTime     Time               `bson:"evtTime"`
	EventType     TokenEventType     `bson:"type"`

	// subject of trade
	Nft          common.Address      `bson:"nft"`
	TokenId      BigInt              `bson:"tokenId"`
	Quantity     *BigInt             `bson:"quantity"`

	// parties
	Seller       *common.Address     `bson:"seller"` // owner before event
	Buyer        *common.Address     `bson:"buyer"` // or offer creator (future buyer)

	// money for the subject
	PayToken     *common.Address     `bson:"payToken"`
	PricePerItem *BigInt             `bson:"pricePerItem"`

	StartTime    *Time               `bson:"startTime"` // for postponed actions
}

func (e *TokenEvent) GenerateId(EventTime uint32, BlockNumber uint64, LogIndex uint) {
	var b primitive.ObjectID
	logIndex := make([]byte, 4)
	binary.BigEndian.PutUint32(b[0:4], EventTime) // from Mongo ObjectID impl
	binary.BigEndian.PutUint64(b[4:12], BlockNumber)
	binary.BigEndian.PutUint32(logIndex, uint32(LogIndex))
	b[4] = b[4] ^ logIndex[3]
	b[5] = b[5] ^ logIndex[2]
	b[6] = b[6] ^ logIndex[1]
	b[7] = b[7] ^ logIndex[0]
	e.Id = b
}
