package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	StartTime    Time           `bson:"startTime"`
	Deadline     Time           `bson:"deadline"`
}

// GenerateId generates unique Offer ID
// One creator can have only one offer for one token.
func (o *Offer) GenerateId() {
	hash := sha256.New()
	hash.Write(o.Nft.Bytes())
	hash.Write(o.TokenId.ToInt().Bytes())
	hash.Write(o.Creator.Bytes())
	o.Id = hash.Sum(nil)
}
