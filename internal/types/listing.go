package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
)

// Listing represents offer for anybody to buy given token from the owner.
type Listing struct {
	Owner        common.Address `bson:"owner"`
	Contract     common.Address `bson:"contract"`
	TokenId      hexutil.Big    `bson:"token"`
	Quantity     hexutil.Big    `bson:"quantity"`
	PayToken     common.Address `bson:"pay_token"`
	UnitPrice    hexutil.Big    `bson:"price"`
	Created      Time           `bson:"created"`
	StartTime    Time           `bson:"start"`
	LastUpdate   *Time          `bson:"updated"`
	Closed       *Time          `bson:"closed"`
	OrdinalIndex int64          `bson:"index"`
}

// ListingsIndexes provides list of indexes expected on listings.
func ListingsIndexes() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixToken}}

	ixOwner := "ix_owner"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "owner", Value: 1}}, Options: &options.IndexOptions{Name: &ixOwner}}
	return ix
}

// ListingID generates unique listing ID for the given contract, token, and owner.
func ListingID(contract *common.Address, tokenID *big.Int, owner *common.Address) primitive.ObjectID {
	hash := sha256.New()
	hash.Write(contract.Bytes())
	hash.Write(tokenID.Bytes())
	hash.Write(owner.Bytes())

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}

// ID generates a unique identifier of the Artion Marketplace listing.
func (l *Listing) ID() primitive.ObjectID {
	return ListingID(&l.Contract, (*big.Int)(&l.TokenId), &l.Owner)
}
