package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// LegacyToken represents token record in legacy database.
type LegacyToken struct {
	Contract  common.Address `bson:"contractAddress"`
	TokenId   int32          `bson:"tokenID"`
	Viewed    int32          `bson:"viewed"`
	Created   Time           `bson:"createdAt"`
	Updated   Time           `bson:"updatedAt"`
	IsActive  bool           `bson:"isAppropriate"`
}
