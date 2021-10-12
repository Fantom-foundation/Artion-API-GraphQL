package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TokenLike represents "like" given by a user to a token (aka adding to favourites).
type TokenLike struct {
	Id       primitive.ObjectID `bson:"_id"`
	User     common.Address     `bson:"follower"`
	Contract common.Address     `bson:"contractAddress"`
	TokenId  hexutil.Big        `bson:"tokenID"`
	Created  Time               `bson:"created"`
}

// GenerateId generates unique identifier for the record.
func (t *TokenLike) GenerateId() {
	t.Id = primitive.NewObjectID()
}
