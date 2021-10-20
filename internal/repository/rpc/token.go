// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// GetPayTokenPrice extracts price of pay token using Marketplace contract.
func (o *Opera) GetPayTokenPrice(_ *common.Address, token *common.Address, _ *big.Int) (*big.Int, error) {
	return o.marketplace.GetPrice(nil, *token)
}
