package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
)

// Token represents item list-able in the marketplace.
type Token struct {
	Id          []byte         `bson:"_id"`
	Contract    common.Address `bson:"nft"`
	TokenId     hexutil.Big    `bson:"tokenId"`
	Uri         string         `bson:"uri"`
	Name        string         `bson:"name"`
	Description string         `bson:"desc"`
	ImageURI    string         `bson:"image"`
	MetadataAge Time           `bson:"metadata_age"`
}

// TokenIdFromAddress generates unique token ID from an NFT contract address and token ID.
func TokenIdFromAddress(adr *common.Address, TokenId *big.Int) []byte {
	hash := sha256.New()
	hash.Write(adr.Bytes())
	hash.Write(TokenId.Bytes())
	return hash.Sum(nil)
}

func (t *Token) GenerateId() {
	t.Id = TokenIdFromAddress(&t.Contract, (*big.Int)(&t.TokenId))
}

// ID generates unique identifier for the NFT owner record.
// Collision approx. for p(n)=1e-10: n=4.000.000.000 tokens indexed
// Collision approx. for p(n)=1e-12: n=500.000.000 tokens indexed
func (t *Token) ID() primitive.ObjectID {
	hash := sha256.New()
	hash.Write(t.Contract.Bytes())
	hash.Write(t.TokenId.ToInt().Bytes())

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}
