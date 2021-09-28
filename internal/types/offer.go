package types

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type OfferType int8

const (
	// CoOffers is the name of database collection.
	CoOffers = "Offers"

	// BSON attributes names used in the database collection.
	FiOfferNft     = "nft"
	FiOfferTokenId = "tokenId"
	FiOfferCreator   = "creator"
)

// Offer represents offer to buy given token from any current owner.
type Offer struct {
	Id           []byte
	Creator      common.Address

	// subject of trade
	Nft          common.Address
	TokenId      hexutil.Big
	Quantity     hexutil.Big

	// money for the subject
	PayToken     common.Address
	PricePerItem hexutil.Big

	StartTime    Time
	Deadline     Time
}

// GenerateId generates unique Offer ID from OfferCreated event attributes
func (e *Offer) GenerateId(EventTime uint64, BlockNumber uint64, LogIndex uint) {
	id := make([]byte, 20, 20)
	binary.BigEndian.PutUint64(id[0:8], EventTime) // time prefix for sorting
	binary.BigEndian.PutUint64(id[8:16], BlockNumber)
	binary.BigEndian.PutUint32(id[16:20], uint32(LogIndex))
	e.Id = id
}

type OfferBson struct {
	Id           []byte    `bson:"_id"`
	Creator      string    `bson:"creator"`
	Nft          string    `bson:"nft"`
	TokenId      string    `bson:"tokenId"`
	Quantity     string    `bson:"quantity"`
	PayToken     string    `bson:"payToken"`
	PricePerItem string    `bson:"pricePerItem"`
	StartTime    time.Time `bson:"startTime"`
	Deadline     time.Time `bson:"deadline"`
}

// MarshalBSON prepares data to be stored in MongoDB.
func (ev *Offer) MarshalBSON() ([]byte, error) {
	row := OfferBson {
		Id: ev.Id,
		Creator: ev.Creator.String(),
		Nft: ev.Nft.String(),
		TokenId: ev.TokenId.String(),
		Quantity: ev.Quantity.String(),
		PayToken: ev.PayToken.String(),
		PricePerItem: ev.PricePerItem.String(),
		StartTime: time.Time(ev.StartTime),
		Deadline: time.Time(ev.Deadline),
	}
	return bson.Marshal(row)
}

// UnmarshalBSON parses data from MongoDB.
func (ev *Offer) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode Offer; %s", err.Error())
		}
	}()

	// try to decode the BSON data
	var row OfferBson
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	ev.Id = row.Id
	ev.Creator = common.HexToAddress(row.Creator)
	ev.Nft = common.HexToAddress(row.Nft)
	ev.TokenId = (hexutil.Big)(*hexutil.MustDecodeBig(row.TokenId))
	ev.Quantity = (hexutil.Big)(*hexutil.MustDecodeBig(row.Quantity))
	ev.PayToken = common.HexToAddress(row.PayToken)
	ev.PricePerItem = (hexutil.Big)(*hexutil.MustDecodeBig(row.PricePerItem))
	ev.StartTime = Time(row.StartTime)
	ev.Deadline = Time(row.Deadline)
	return nil
}
