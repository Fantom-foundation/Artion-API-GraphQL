package rpc

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// IMarketplaceContract defines single interface for all marketplace contract versions.
type IMarketplaceContract interface {
	GetPayTokenPrice(token *common.Address, block *big.Int) (*big.Int, error)
}

// GetPayTokenPrice extracts price of 1 whole pay token in USD in 6-decimals fixed point using Marketplace contract.
func (o *Opera) GetPayTokenPrice(token *common.Address, block *big.Int) (*big.Int, error) {
	return o.payTokenPriceContract.GetPayTokenPrice(token, block)
}
