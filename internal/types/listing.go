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
	// CoListings is the name of database collection.
	CoListings = "Listings"

	// BSON attributes names used in the database collection.
	FiListingNft     = "nft"
	FiListingTokenId = "tokenId"
	FiListingOwner   = "owner"
)

// Listing represents offer for anybody to buy given token from the owner.
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

// GenerateId generates unique listing ID
func (l *Listing) GenerateId() {
	hash := md5.New()
	hash.Write(l.Nft.Bytes())
	hash.Write(l.TokenId.ToInt().Bytes())
	hash.Write(l.Owner.Bytes())
	l.Id = hash.Sum(nil)
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
func (l *Listing) MarshalBSON() ([]byte, error) {
	row := ListingBson {
		Id:           l.Id,
		Owner:        l.Owner.String(),
		Nft:          l.Nft.String(),
		TokenId:      l.TokenId.String(),
		Quantity:     l.Quantity.String(),
		PayToken:     l.PayToken.String(),
		PricePerItem: l.PricePerItem.String(),
		StartTime:    time.Time(l.StartTime),
	}
	return bson.Marshal(row)
}

// UnmarshalBSON parses data from MongoDB.
func (l *Listing) UnmarshalBSON(data []byte) (err error) {
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

	l.Id = row.Id
	l.Owner = common.HexToAddress(row.Owner)
	l.Nft = common.HexToAddress(row.Nft)
	l.TokenId = (hexutil.Big)(*hexutil.MustDecodeBig(row.TokenId))
	l.Quantity = (hexutil.Big)(*hexutil.MustDecodeBig(row.Quantity))
	l.PayToken = common.HexToAddress(row.PayToken)
	l.PricePerItem = (hexutil.Big)(*hexutil.MustDecodeBig(row.PricePerItem))
	l.StartTime = Time(row.StartTime)
	return nil
}
