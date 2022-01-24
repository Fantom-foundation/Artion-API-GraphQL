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

func (cf CollectionFilter) IsUsed() bool {
	return (cf.Search != nil && *cf.Search != "") ||
		cf.MintableBy != nil ||
		cf.InReview ||
		cf.Banned
}
