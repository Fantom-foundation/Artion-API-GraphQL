// Package types provides high level structures for the API server.
package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
)

// Auction represents an auction being conducted on a given ERC-721 token.
type Auction struct {
	Contract      common.Address  `bson:"contract"`
	TokenId       hexutil.Big     `bson:"token"`
	Owner         common.Address  `bson:"owner"`
	Quantity      hexutil.Big     `bson:"qty"`
	PayToken      common.Address  `bson:"pay_token"`
	ReservePrice  hexutil.Big     `bson:"price"`
	Created       Time            `bson:"created"`
	StartTime     Time            `bson:"start"`
	EndTime       Time            `bson:"end"`
	Closed        *Time           `bson:"closed"`
	LastBid       *hexutil.Big    `bson:"last_bid"`
	LastBidPlaced *Time           `bson:"last_bid_ts"`
	LastBidder    *common.Address `bson:"last_bidder"`
	Winner        *common.Address `bson:"winner"`
	WinningBid    *hexutil.Big    `bson:"win_bid"`
	Resolved      *Time           `bson:"resolved"`
	OrdinalIndex  int64           `bson:"index"`
}

// AuctionID generates unique auction ID for the given contract, token, and owner.
func AuctionID(contract *common.Address, tokenID *big.Int) primitive.ObjectID {
	hash := sha256.New()
	hash.Write(contract.Bytes())
	hash.Write(tokenID.Bytes())

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}

// ID generates a unique identifier of the Artion auction.
func (au *Auction) ID() primitive.ObjectID {
	return AuctionID(&au.Contract, (*big.Int)(&au.TokenId))
}
