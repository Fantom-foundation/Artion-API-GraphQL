// Package types provides high level structures for the API server.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents user account/profile.
type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Address  common.Address `bson:"address"`
	Username *string        `bson:"alias"`
	Email    *string        `bson:"email"`
	Bio      *string        `bson:"bio"`
	Avatar   *string        `bson:"imageHash"`
	Banner   *string        `bson:"bannerHash"`
	Bundles  []string       `bson:"bundleIDs"`
	Created  Time           `bson:"createdAt"`
	Updated  Time           `bson:"updatedAt"`
}
