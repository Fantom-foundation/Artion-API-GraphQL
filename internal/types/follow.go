package types

import (
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Follow represents "following" of one user by another user.
type Follow struct {
	Id       primitive.ObjectID `bson:"_id"`
	Follower common.Address     `bson:"from"`
	Followed common.Address     `bson:"to"`
}
