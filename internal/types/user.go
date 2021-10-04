package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// User represents user account/profile.
type User struct {
	Address  common.Address `bson:"_id"`
	Username string         `bson:"username"`
	Bio      string         `bson:"bio"`
	Email    string         `bson:"email"`
	Avatar   string         `bson:"avatar"`
}
