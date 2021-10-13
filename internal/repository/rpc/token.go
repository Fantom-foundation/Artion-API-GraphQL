// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// GetUnitPriceAt extracts token price using Marketplace contract.
func (o *Opera) GetUnitPriceAt(contract *common.Address, token *common.Address, block *big.Int) (*big.Int, error) {
	// try on block
	val, err := o.marketplace.GetPrice(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: block,
		Context:     context.Background(),
	}, *token)
	if err != nil {
		log.Warningf("get %s price failed for block #%d; %s", token.String(), block.Uint64(), err.Error())
		return o.marketplace.GetPrice(nil, *token)
	}
	return val, nil
}
