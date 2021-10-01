package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

const (
	// CoTokens is the name of database collection.
	CoTokens = "Tokens"

	// FiTokenNft is the name of DB column storing NFT contract address.
	FiTokenNft = "nft"

	// FiTokenTokenId is the name of DB column storing NFT token ID.
	FiTokenTokenId = "tokenId"

	// FiTokenName is the column storing the name of the NFT token.
	FiTokenName = "name"
)

// Token represents item list-able in the marketplace.
type Token struct {
	Id      []byte         `bson:"_id"`
	Nft     common.Address `bson:"nft"`
	TokenId hexutil.Big    `bson:"tokenId"`
	Uri     string         `bson:"uri"`
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
