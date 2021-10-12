package types

import (
	"crypto/md5"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
)

// TokenLike represents "like" given by a user to a token (aka adding to favourites).
type TokenLike struct {
	User      common.Address `bson:"user"`
	Contract common.Address `bson:"contract"`
	TokenId  hexutil.Big    `bson:"token"`
	Created  Time           `bson:"created"`
}

func GetTokenLikeId(user *common.Address, contract *common.Address, tokenId *big.Int) primitive.ObjectID {
	hash := md5.New() // provides 16 bytes - more the sufficient to get 12-bytes id
	hash.Write(contract.Bytes())
	hash.Write(tokenId.Bytes())
	hash.Write(user.Bytes())

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}

// ID generates unique identifier for the NFT owner record.
func (t *TokenLike) ID() primitive.ObjectID {
	return GetTokenLikeId(&t.User, &t.Contract, (*big.Int)(&t.TokenId))
}
