package types

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type ListingType int8

const (
	// CoListings is the name of database collection.
	CoListings = "Listings"

	// BSON attributes names used in the database collection.
	FiListingNft     = "nft"
	FiListingTokenId = "tokenId"
	FiListingOwner   = "owner"
)

// Listing represents offer to buy given token from given owner.
type Listing struct {
	Id           []byte
	Owner        common.Address

	// subject of trade
	Nft          common.Address
	TokenId      hexutil.Big
	Quantity     hexutil.Big

	// money for the subject
	PayToken     common.Address
	PricePerItem hexutil.Big

	StartTime    Time
}

// GenerateId generates unique listing ID from ItemListed event attributes
func (e *Listing) GenerateId(EventTime uint64, BlockNumber uint64, LogIndex uint) {
	id := make([]byte, 20, 20)
	binary.BigEndian.PutUint64(id[0:8], EventTime) // time prefix for sorting
	binary.BigEndian.PutUint64(id[8:16], BlockNumber)
	binary.BigEndian.PutUint32(id[16:20], uint32(LogIndex))
	e.Id = id
}

type ListingBson struct {
	Id           []byte    `bson:"_id"`
	Owner        string    `bson:"owner"`
	Nft          string    `bson:"nft"`
	TokenId      string    `bson:"tokenId"`
	Quantity     string    `bson:"quantity"`
	PayToken     string    `bson:"payToken"`
	PricePerItem string    `bson:"pricePerItem"`
	StartTime    time.Time `bson:"startTime"`
}

// MarshalBSON prepares data to be stored in MongoDB.
func (ev *Listing) MarshalBSON() ([]byte, error) {
	row := ListingBson {
		Id: ev.Id,
		Owner: ev.Owner.String(),
		Nft: ev.Nft.String(),
		TokenId: ev.TokenId.String(),
		Quantity: ev.Quantity.String(),
		PayToken: ev.PayToken.String(),
		PricePerItem: ev.PricePerItem.String(),
		StartTime: time.Time(ev.StartTime),
	}
	return bson.Marshal(row)
}

// UnmarshalBSON parses data from MongoDB.
func (ev *Listing) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode Listing; %s", err.Error())
		}
	}()

	// try to decode the BSON data
	var row ListingBson
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	ev.Id = row.Id
	ev.Owner = common.HexToAddress(row.Owner)
	ev.Nft = common.HexToAddress(row.Nft)
	ev.TokenId = (hexutil.Big)(*hexutil.MustDecodeBig(row.TokenId))
	ev.Quantity = (hexutil.Big)(*hexutil.MustDecodeBig(row.Quantity))
	ev.PayToken = common.HexToAddress(row.PayToken)
	ev.PricePerItem = (hexutil.Big)(*hexutil.MustDecodeBig(row.PricePerItem))
	ev.StartTime = Time(row.StartTime)
	return nil
}
