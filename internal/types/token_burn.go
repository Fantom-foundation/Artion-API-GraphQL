// Package types provides high level structures for the API server.
package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NFTBurn represents an NFT burn.
type NFTBurn struct {
	Contract common.Address `bson:"contract"`
	TokenId  hexutil.Big    `bson:"token"`
	Owner    common.Address `bson:"owner"`
	Qty      hexutil.Big    `bson:"qty"`
	Burned   Time           `bson:"burned"`
}

// ID generates unique identifier for the NFT burn record.
// Collision approx for p(n)=1e-12: n = sqrt(2 x 2^96 x 2^-39) = 536.870.912 documents
func (bu *NFTBurn) ID() primitive.ObjectID {
	hash := sha256.New()
	hash.Write(bu.Contract.Bytes())
	hash.Write(bu.TokenId.ToInt().Bytes())
	hash.Write(bu.Owner.Bytes())

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}
