package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// MarketplaceContractV1 is IMarketplaceContract implementation for contracts in 1.0.0 version
type MarketplaceContractV1 struct {
	marketplace *contracts.FantomMarketplace
}

// GetPayTokenPrice extracts price of 1 whole pay token in USD in 6-decimals fixed point using Marketplace contract.
func (mc *MarketplaceContractV1) GetPayTokenPrice(token *common.Address, _ *big.Int) (*big.Int, error) {
	return mc.marketplace.GetPrice(nil, *token)
}
