package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// Token represents item list-able in the marketplace.
type Token struct {
	Id          []byte         `bson:"_id"`
	Nft         common.Address `bson:"nft"`
	TokenId     hexutil.Big    `bson:"tokenId"`
	Uri         string         `bson:"uri"`
	Name        string         `bson:"name"`
	Description string         `bson:"desc"`
	ImageURI    string         `bson:"image"`
	MetadataAge Time           `bson:"metadata_age"`
}

// TokenIdFromAddress generates unique token ID from an NFT contract address and token ID.
func TokenIdFromAddress(Nft *common.Address, TokenId *big.Int) []byte {
	hash := sha256.New()
	hash.Write(Nft.Bytes())
	hash.Write(TokenId.Bytes())
	return hash.Sum(nil)
}

func (t *Token) GenerateId() {
	t.Id = TokenIdFromAddress(&t.Nft, (*big.Int)(&t.TokenId))
}
