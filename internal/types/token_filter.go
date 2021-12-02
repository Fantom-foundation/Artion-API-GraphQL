package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// TokenFilter represents a filter for token lists.
type TokenFilter struct {
	Search          *string
	HasListing      *bool
	HasAuction      *bool
	HasOffer        *bool
	HasBids         *bool
	IncludeInactive *bool
	Collections     *[]common.Address
	Categories      *[]int32
	CreatedBy       *common.Address
	PriceMin        *hexutil.Big
	PriceMax        *hexutil.Big
	ListPriceMin    *hexutil.Big
	ListPriceMax    *hexutil.Big
	OfferPriceMin   *hexutil.Big
	OfferPriceMax   *hexutil.Big
}
