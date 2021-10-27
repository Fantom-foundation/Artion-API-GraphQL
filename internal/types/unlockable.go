package types

import "github.com/ethereum/go-ethereum/common"

// Unlockable represents unlockable content attached to an NFT token.
type Unlockable struct {
	Contract common.Address `bson:"contractAddress"`
	TokenId  int32          `bson:"tokenID"`
	Content  string         `bson:"content"`
}
