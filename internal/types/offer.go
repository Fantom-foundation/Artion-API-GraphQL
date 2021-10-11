// Package types provides high level structures for the API server.
package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
)

// Offer represents offer to buy given token from any current owner.
type Offer struct {
	Contract     common.Address `bson:"contract"`
	TokenId      hexutil.Big    `bson:"token"`
	ProposedBy   common.Address `bson:"proposer"`
	Quantity     hexutil.Big    `bson:"qty"`
	PayToken     common.Address `bson:"pay_token"`
	UnitPrice    hexutil.Big    `bson:"price"`
	Created      Time           `bson:"created"`
	Deadline     Time           `bson:"deadline"`
	Closed       *Time          `bson:"closed"`
	OrdinalIndex int64          `bson:"index"`
}

// OfferID generates unique offer ID for the given contract, token, and owner.
func OfferID(contract *common.Address, tokenID *big.Int, offeredBy *common.Address) primitive.ObjectID {
	hash := sha256.New()
	hash.Write(contract.Bytes())
	hash.Write(tokenID.Bytes())
	hash.Write(offeredBy.Bytes())

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}

// ID generates a unique identifier of the Artion Marketplace direct offer.
func (o *Offer) ID() primitive.ObjectID {
	return OfferID(&o.Contract, (*big.Int)(&o.TokenId), &o.ProposedBy)
}
