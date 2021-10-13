package types

import (
	"github.com/ethereum/go-ethereum/common"
)

type TokenFilter struct {
	Search     *string
	HasListing *bool
	HasAuction *bool
	HasOffer   *bool
	HasBids    *bool
	Collections *[]common.Address
	CreatedBy  *common.Address
}
