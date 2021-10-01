package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"time"
)

const (
	// CoOffers is the name of database collection.
	CoOffers = "Offers"

	// FiOfferNft is the name of the DB column of the NFT contract address.
	FiOfferNft = "nft"

	// FiOfferTokenId is the name of the DB column of the token ID.
	FiOfferTokenId = "tokenId"

	// FiOfferCreator is the name of the DB column of the offer creator address.
	FiOfferCreator = "creator"
)

// Offer represents offer to buy given token from any current owner.
type Offer struct {
	Id           []byte         `bson:"_id"`
	Creator      common.Address `bson:"creator"`
	Nft          common.Address `bson:"nft"`
	TokenId      hexutil.Big    `bson:"tokenId"`
	Quantity     hexutil.Big    `bson:"quantity"`
	PayToken     common.Address `bson:"payToken"`
	PricePerItem hexutil.Big    `bson:"pricePerItem"`
	StartTime    time.Time      `bson:"startTime"`
	Deadline     time.Time      `bson:"deadline"`
}

// GenerateId generates unique Offer ID
// @todo Make sure the ID is unique.
func (o *Offer) GenerateId() {
	hash := sha256.New()
	hash.Write(o.Nft.Bytes())
	hash.Write(o.TokenId.ToInt().Bytes())
	hash.Write(o.Creator.Bytes())
	o.Id = hash.Sum(nil)
}
