package types

import (
	"crypto/md5"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

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

// GenerateId generates unique Offer ID
func (o *Offer) GenerateId() {
	hash := md5.New()
	hash.Write(o.Nft.Bytes())
	hash.Write(o.TokenId.ToInt().Bytes())
	hash.Write(o.Creator.Bytes())
	o.Id = hash.Sum(nil)
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
func (o *Offer) MarshalBSON() ([]byte, error) {
	row := OfferBson {
		Id:           o.Id,
		Creator:      o.Creator.String(),
		Nft:          o.Nft.String(),
		TokenId:      o.TokenId.String(),
		Quantity:     o.Quantity.String(),
		PayToken:     o.PayToken.String(),
		PricePerItem: o.PricePerItem.String(),
		StartTime:    time.Time(o.StartTime),
		Deadline:     time.Time(o.Deadline),
	}
	return bson.Marshal(row)
}

// UnmarshalBSON parses data from MongoDB.
func (o *Offer) UnmarshalBSON(data []byte) (err error) {
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

	o.Id = row.Id
	o.Creator = common.HexToAddress(row.Creator)
	o.Nft = common.HexToAddress(row.Nft)
	o.TokenId = (hexutil.Big)(*hexutil.MustDecodeBig(row.TokenId))
	o.Quantity = (hexutil.Big)(*hexutil.MustDecodeBig(row.Quantity))
	o.PayToken = common.HexToAddress(row.PayToken)
	o.PricePerItem = (hexutil.Big)(*hexutil.MustDecodeBig(row.PricePerItem))
	o.StartTime = Time(row.StartTime)
	o.Deadline = Time(row.Deadline)
	return nil
}
