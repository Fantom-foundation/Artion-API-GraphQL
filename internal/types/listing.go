// Package types provides high level structures for the API server.
package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	UnifiedPrice int64          `bson:"uprice"`
	Created      Time           `bson:"created"`
	StartTime    Time           `bson:"start"`
	LastUpdate   *Time          `bson:"updated"`
	Closed       *Time          `bson:"closed"`
	OrdinalIndex int64          `bson:"index"`
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
