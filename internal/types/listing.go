package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"time"
)

const (
	// CoListings is the name of database collection.
	CoListings = "Listings"

	// FiListingNft is the name of the DB column storing NFT contract address.
	FiListingNft     = "nft"

	// FiListingTokenId represents the name of the DB column storing NFT token ID.
	FiListingTokenId = "tokenId"

	// FiListingOwner represents the name of the DB column storing token owner.
	FiListingOwner   = "owner"
)

// Listing represents offer for anybody to buy given token from the owner.
type Listing struct {
	Id           []byte         `bson:"_id"`
	Owner        common.Address `bson:"owner"`
	Nft          common.Address `bson:"nft"`
	TokenId      hexutil.Big    `bson:"tokenId"`
	Quantity     hexutil.Big    `bson:"quantity"`
	PayToken     common.Address `bson:"payToken"`
	PricePerItem hexutil.Big    `bson:"pricePerItem"`
	StartTime    Time           `bson:"startTime"`
}

// GenerateId generates unique listing ID
// One owner can have only one listing of one token.
func (l *Listing) GenerateId() {
	hash := sha256.New()
	hash.Write(l.Nft.Bytes())
	hash.Write(l.TokenId.ToInt().Bytes())
	hash.Write(l.Owner.Bytes())
	hash.Write(([]byte)(time.Time(l.StartTime).String()))
	l.Id = hash.Sum(nil)
}
