package types

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TokenEventType int8

const (
	// CoTokenEvents is the name of database collection.
	CoTokenEvents = "TokenEvents"

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
	Id            primitive.ObjectID
	EventTime     Time
	EventType     TokenEventType

	// subject of trade
	Nft          common.Address
	TokenId      hexutil.Big
	Quantity     *hexutil.Big

	// parties
	Seller       *common.Address // owner before event
	Buyer        *common.Address // or offer creator (future buyer)

	// money for the subject
	PayToken     *common.Address
	PricePerItem *hexutil.Big

	StartTime    *Time // for postponed actions
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

type tokenEventBson struct {
	Id            primitive.ObjectID `bson:"_id"`
	EventTime     time.Time          `bson:"evtTime"`
	EventType     int32              `bson:"type"`
	Nft          string              `bson:"nft"`
	TokenId      string              `bson:"tokenId"`
	Quantity     *string             `bson:"quantity"`
	Seller       *string             `bson:"seller"`
	Buyer        *string             `bson:"buyer"`
	PayToken     *string             `bson:"payToken"`
	PricePerItem *string             `bson:"pricePerItem"`
	StartTime    *time.Time          `bson:"startTime"`
}

// MarshalBSON prepares data to be stored in MongoDB.
func (ev *TokenEvent) MarshalBSON() ([]byte, error) {
	row := tokenEventBson {
		Id: ev.Id,
		EventTime: time.Time(ev.EventTime),
		EventType: int32(ev.EventType),
		Nft: ev.Nft.String(),
		TokenId: (&ev.TokenId).String(),
	}
	if ev.Quantity != nil {
		quantity := ev.Quantity.String()
		row.Quantity = &quantity
	}
	if ev.Seller != nil {
		seller := ev.Seller.String()
		row.Seller = &seller
	}
	if ev.Buyer != nil {
		buyer := ev.Buyer.String()
		row.Buyer = &buyer
	}
	if ev.PayToken != nil {
		payToken := ev.PayToken.String()
		row.PayToken = &payToken
	}
	if ev.PricePerItem != nil {
		price := ev.PricePerItem.String()
		row.PricePerItem = &price
	}
	row.StartTime = (*time.Time)(ev.StartTime)
	return bson.Marshal(row)
}

// UnmarshalBSON parses data from MongoDB.
func (ev *TokenEvent) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode TokenEvent; %s", err.Error())
		}
	}()

	// try to decode the BSON data
	var row tokenEventBson
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	ev.Id = row.Id
	ev.EventTime = Time(row.EventTime)
	ev.EventType = TokenEventType(row.EventType)
	ev.Nft = common.HexToAddress(row.Nft)
	ev.TokenId = (hexutil.Big)(*hexutil.MustDecodeBig(row.TokenId))
	if row.Quantity != nil {
		quantity := (hexutil.Big)(*hexutil.MustDecodeBig(*row.Quantity))
		ev.Quantity = &quantity
	}
	if row.Seller != nil {
		seller := common.HexToAddress(*row.Seller)
		ev.Seller = &seller
	}
	if row.Buyer != nil {
		buyer := common.HexToAddress(*row.Buyer)
		ev.Buyer = &buyer
	}
	if row.PayToken != nil {
		payToken := common.HexToAddress(*row.PayToken)
		ev.PayToken = &payToken
	}
	if row.PricePerItem != nil {
		price := (hexutil.Big)(*hexutil.MustDecodeBig(*row.PricePerItem))
		ev.PricePerItem = &price
	}
	ev.StartTime = (*Time)(row.StartTime)
	return nil
}
