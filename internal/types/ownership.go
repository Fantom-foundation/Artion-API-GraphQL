// Package types provides high level structures for the API server.
package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Ownership represents an ownership reference between user address and owned NFT.
type Ownership struct {
	Contract common.Address `bson:"contract"`
	TokenId  hexutil.Big    `bson:"tokenId"`
	Owner    common.Address `bson:"owner"`
	Qty      hexutil.Big    `bson:"qty"`
	Updated  Time           `bson:"updated"`
}

// ID generates unique identifier for the NFT owner record.
// Collision approx for p(n)=1e-12: n = sqrt(2 x 2^96 x 2^-39) = 536.870.912 documents
func (no *Ownership) ID() primitive.ObjectID {
	hash := sha256.New()
	hash.Write(no.Contract.Bytes())
	hash.Write(no.TokenId.ToInt().Bytes())
	hash.Write(no.Owner.Bytes())

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}
