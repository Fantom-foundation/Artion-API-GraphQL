package types

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
	"time"
)

type TokenEventType int8

const (
	EvtTpItemListed TokenEventType = iota
	EvtTpItemUpdated
	EvtTpItemCanceled
	EvtTpItemSold
	EvtTpOfferCreated
	EvtTpOfferCanceled
)

type TokenEvent struct {
	Id            primitive.ObjectID `bson:"_id"`
	EventTime     time.Time          `bson:"evtTime"`
	Type          TokenEventType     `bson:"type"`

	// subject of trade
	Nft          common.Address      `bson:"nft"`
	TokenId      common.Address      `bson:"tokenId"`
	Quantity     big.Int             `bson:"quantity"`

	// parties
	Seller       common.Address      `bson:"seller"`
	Buyer        common.Address      `bson:"buyer"` // or offer creator - who buys the NFT

	// money for the subject
	PayToken     common.Address      `bson:"payToken"`
	PricePerItem big.Int             `bson:"pricePerItem"`

	StartTime time.Time              `bson:"startTime"` // for postponed actions
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
