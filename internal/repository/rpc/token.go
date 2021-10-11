// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// GetUnitPriceAt extracts token price using Marketplace contract.
func (o *Opera) GetUnitPriceAt(contract *common.Address, token *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.abiMarketplace.Pack("getPrice", *token)
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}
