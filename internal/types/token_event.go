package types

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type TokenEventType int8

const (
	// CoTokenEvents is the name of database collection.
	CoTokenEvents = "TokenEvents"

	// BSON attributes names used in the database collection.
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
	Id            []byte
	EventTime     time.Time
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

	StartTime    *time.Time // for postponed actions
}

func (e *TokenEvent) GenerateId(EventTime uint64, BlockNumber uint64, LogIndex uint) {
	id := make([]byte, 20, 20)
	binary.BigEndian.PutUint64(id[0:8], EventTime) // time prefix for sorting
	binary.BigEndian.PutUint64(id[8:16], BlockNumber)
	binary.BigEndian.PutUint32(id[16:20], uint32(LogIndex))
	e.Id = id
}

type tokenEventBson struct {
	Id            []byte             `bson:"_id"`
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
	ev.EventTime = time.Time(row.EventTime)
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
	ev.StartTime = (*time.Time)(row.StartTime)
	return nil
}
