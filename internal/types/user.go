package types

import (
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User represents user account/profile.
type User struct {
	Address  common.Address `bson:"_id"`
	Username string         `bson:"username"`
	Bio      string         `bson:"bio"`
	Email    string         `bson:"email"`
	Avatar   string         `bson:"avatar"`
}

// UsersIndexes provides a list of indexes expected to exist on users' collection.
func UsersIndexes() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixUser := "ix_username"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "username", Value: 1}}, Options: &options.IndexOptions{Name: &ixUser}}

	ixEmail := "ixEmail"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "email", Value: 1}}, Options: &options.IndexOptions{Name: &ixEmail}}
	return ix
}
