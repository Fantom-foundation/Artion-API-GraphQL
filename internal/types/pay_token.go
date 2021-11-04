package types

import "github.com/ethereum/go-ethereum/common"

// PayToken represents token supported for payments on the marketplace
type PayToken struct {
	Contract common.Address
	Name     string
	Symbol string
	Decimals int32
}
