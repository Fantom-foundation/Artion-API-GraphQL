// Package types provides high level structures for the API server.
package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// User represents user account/profile.
type User struct {
	Address  common.Address     `bson:"address"`
	Username string             `bson:"alias"`
	Email    string             `bson:"email"`
	Bio      string             `bson:"bio"`
	Avatar   string             `bson:"imageHash"`
	Banner   string             `bson:"bannerHash"`
	Bundles  []string           `bson:"bundleIDs"`
	Created  Time               `bson:"createdAt"`
	Updated  Time               `bson:"updatedAt"`
}
