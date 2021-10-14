package types

import (
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TokenLike represents "like" given by a user to a token (aka adding to favourites).
type TokenLike struct {
	Id        primitive.ObjectID `bson:"_id"`
	User      common.Address     `bson:"follower"`
	Contract  common.Address     `bson:"contractAddress"`
	TokenId32 int32              `bson:"tokenID"`
}
