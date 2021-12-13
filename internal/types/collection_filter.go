package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// CollectionFilter represents a filter for collections lists.
type CollectionFilter struct {
	Search *string
	MintableBy *common.Address
	InReview bool
	Banned bool
}

