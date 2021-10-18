// Package types provides high level structures for the API server.
package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
	"time"
)

const (
	// TokenDefaultMetadataUpdateDelay is the minimal time delay between subsequent
	// metadata update attempts of new tokens.
	TokenDefaultMetadataUpdateDelay = 5 * time.Minute

	// TokenSuccessMetadataUpdateDelay is the minimal time delay before
	// a new metadata update attempt after a successful update.
	TokenSuccessMetadataUpdateDelay = 7 * 24 * time.Hour

	// MetadataRefreshSetSize is the max size of metadata refresh set pulled at once.
	MetadataRefreshSetSize = 50
)

// Token represents item list-able in the marketplace.
type Token struct {
	Contract        common.Address `bson:"contract"`
	TokenId         hexutil.Big    `bson:"token"`
	IsActive        bool           `bson:"is_active"`
	Uri             string         `bson:"uri"`
	Name            string         `bson:"name"`
	Description     string         `bson:"desc"`
	ImageURI        string         `bson:"image"`
	OrdinalIndex    int64          `bson:"index"`
	Created         Time           `bson:"created"`
	CreatedBy       common.Address `bson:"created_by"`
	HasListingSince *Time          `bson:"listed_since"`
	HasAuctionSince *Time          `bson:"auction_since"`
	HasAuctionUntil *Time          `bson:"auction_until"`
	HasOfferUntil   *Time          `bson:"offer_until"`
	HasBids         bool           `bson:"has_bid"`
	LastTrade       *Time          `bson:"last_trade"`
	LastListing     *Time          `bson:"last_list"`
	LastOffer       *Time          `bson:"last_offer"`
	LastAuction     *Time          `bson:"last_auction"`
	LastBid         *Time          `bson:"last_bid"`
	AmountLastTrade int64          `bson:"amo_trade"`
	AmountLastOffer int64          `bson:"amo_offer"`
	AmountLastBid   int64          `bson:"amo_bid"`
	AmountLastList  int64          `bson:"amo_list"`
	Price           int64          `bson:"price"`
	Categories      []int          `bson:"categories"`

	// metadata refresh helpers
	MetaUpdate   Time  `bson:"meta_update"`
	MetaFailures int32 `bson:"meta_failures"`
}

// OrdinalIndex generates numeric ordinal index from block number and log record index.
func OrdinalIndex(blk int64, index int64) int64 {
	return (blk<<12)&0x7FFFFFFFFFFFFFFF | (index & 0x3fff)
}

// NewToken creates a new instance of the Token structure.
func NewToken(con *common.Address, tokenId *big.Int, uri string, ts int64, block uint64, index uint) *Token {
	return &Token{
		Contract:     *con,
		TokenId:      hexutil.Big(*tokenId),
		IsActive:     false,
		Uri:          uri,
		Created:      Time(time.Unix(ts, 0)),
		OrdinalIndex: OrdinalIndex(int64(block), int64(index)),
		MetaUpdate:   Time(time.Now().Add(TokenDefaultMetadataUpdateDelay)),
	}
}

// TokenID generates unique token ID from an NFT contract address and token ID.
// Collision approx. for p(n)=1e-10: n=4.000.000.000 tokens indexed
// Collision approx. for p(n)=1e-12: n=500.000.000 tokens indexed
func TokenID(adr *common.Address, tokenId *big.Int) primitive.ObjectID {
	hash := sha256.New()
	hash.Write(adr.Bytes())
	hash.Write(tokenId.Bytes())

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}

// ID generates unique identifier for the NFT owner record.
func (t *Token) ID() primitive.ObjectID {
	return TokenID(&t.Contract, (*big.Int)(&t.TokenId))
}

// ScheduleMetaUpdateOnFailure sets new metadata update time after failed attempt.
// Every failure makes the next delay longer since we expect the failure to happen again.
func (t *Token) ScheduleMetaUpdateOnFailure() {
	t.MetaUpdate = Time(time.Now().Add(time.Duration(2*t.MetaFailures+1) * TokenDefaultMetadataUpdateDelay))
	t.MetaFailures++
}

// ScheduleMetaUpdateOnSuccess sets new metadata update time successful metadata update.
func (t *Token) ScheduleMetaUpdateOnSuccess() {
	t.MetaUpdate = Time(time.Now().Add(TokenSuccessMetadataUpdateDelay))
	t.MetaFailures = 0
}
