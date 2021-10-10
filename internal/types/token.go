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
	Uri             string         `bson:"uri"`
	Name            string         `bson:"name"`
	Description     string         `bson:"desc"`
	ImageURI        string         `bson:"image"`
	OrdinalIndex    int64          `bson:"index"`
	Created         Time           `bson:"created"`
	IsListed        bool           `bson:"is_listed"`
	IsAuction       bool           `bson:"is_auction"`
	HasOffers       bool           `bson:"has_offers"`
	HasBids         bool           `bson:"has_bids"`
	AmountLastTrade hexutil.Big    `bson:"amo_trade"`
	AmountLastOffer hexutil.Big    `bson:"amo_offer"`
	AmountLastBid   hexutil.Big    `bson:"amo_bid"`
	AmountLastList  hexutil.Big    `bson:"amo_list"`
	LastTrade       Time           `bson:"last_trade"`
	LastOffer       Time           `bson:"last_offer"`
	LastBid         Time           `bson:"last_bid"`
	LastList        Time           `bson:"last_list"`
	Price           hexutil.Big    `bson:"price"`

	// metadata refresh helpers
	MetaUpdate   Time  `bson:"meta_update"`
	MetaFailures int32 `bson:"meta_failures"`
}

// TokensIndexes provides a list of indexes expected to exist on tokens' collection.
func TokensIndexes() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixContractToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixContractToken}}

	ixOrdinal := "ix_ordinal"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "index", Value: -1}}, Options: &options.IndexOptions{Name: &ixOrdinal}}
	return ix
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
		Uri:          uri,
		Created:      Time(time.Unix(ts, 0)),
		OrdinalIndex: OrdinalIndex(int64(block), int64(index)),
		MetaUpdate:   Time(time.Now().Add(TokenDefaultMetadataUpdateDelay)),
	}
}

// TokenIdFromAddress generates unique token ID from an NFT contract address and token ID.
// Collision approx. for p(n)=1e-10: n=4.000.000.000 tokens indexed
// Collision approx. for p(n)=1e-12: n=500.000.000 tokens indexed
func TokenIdFromAddress(adr *common.Address, tokenId *big.Int) primitive.ObjectID {
	hash := sha256.New()
	hash.Write(adr.Bytes())
	hash.Write(tokenId.Bytes())

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}

// ID generates unique identifier for the NFT owner record.
func (t *Token) ID() primitive.ObjectID {
	return TokenIdFromAddress(&t.Contract, (*big.Int)(&t.TokenId))
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
