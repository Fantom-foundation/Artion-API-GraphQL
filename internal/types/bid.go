// Package types provides high level structures for the API server.
package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
)

// AuctionBid represents a single auction bid towards an auction.
type AuctionBid struct {
	Contract common.Address `bson:"contract"`
	TokenId  hexutil.Big    `bson:"token"`
	Bidder   common.Address `bson:"bidder"`
	Placed   Time           `bson:"placed"`
	Amount   hexutil.Big    `bson:"amount"`
}

// AuctionBidID generates unique auction bid ID for the given contract, token, and owner.
func AuctionBidID(contract *common.Address, tokenID *big.Int, bidder *common.Address) primitive.ObjectID {
	hash := sha256.New()
	hash.Write(contract.Bytes())
	hash.Write(tokenID.Bytes())
	hash.Write(bidder.Bytes())

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}

// ID generates a unique identifier of the Artion auction bid.
func (bid *AuctionBid) ID() primitive.ObjectID {
	return AuctionBidID(&bid.Contract, (*big.Int)(&bid.TokenId), &bid.Bidder)
}
