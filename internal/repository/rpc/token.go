// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// GetUnitPriceAt extracts token price using Marketplace contract.
func (o *Opera) GetUnitPriceAt(_ *common.Address, token *common.Address, _ *big.Int) (*big.Int, error) {
	return o.marketplace.GetPrice(nil, *token)
}
