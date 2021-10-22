// Package types provides high level structures for the API server.
package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RandomTrade represents random trade being conducted on a set of ERC-721 tokens.
type RandomTrade struct {
	Contract     common.Address `bson:"contract"`
	TradeManager common.Address `bson:"manager"`
	Name         string         `bson:"name"`
	TradeStarts  Time           `bson:"starts"`
	TradeEnds    Time           `bson:"ends"`
}

// RandomTradeID generates unique ID for the given random trade.
func RandomTradeID(contract *common.Address) primitive.ObjectID {
	hash := sha256.New()
	hash.Write(contract.Bytes())

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}

// ID generates a unique identifier of the Random Trade auction.
func (rt *RandomTrade) ID() primitive.ObjectID {
	return RandomTradeID(&rt.Contract)
}
