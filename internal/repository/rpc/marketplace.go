package rpc

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// IMarketplaceContract defines single interface for all marketplace contract versions.
type IMarketplaceContract interface {
	GetTokenRoyalty(contract common.Address, tokenId *big.Int) (royalty uint16, recipient common.Address, err error)
	GetPayTokenPrice(token *common.Address, block *big.Int) (*big.Int, error)
}

// GetTokenRoyalty provides fee for token minter when the token is sold and its recipient (royalty has 2 decimals)
func (o *Opera) GetTokenRoyalty(contract common.Address, tokenId *big.Int) (royalty uint16, recipient common.Address, err error) {
	return o.defaultMarketplaceContract.GetTokenRoyalty(contract, tokenId)
}

// GetPayTokenPrice extracts price of 1 whole pay token in USD in 6-decimals fixed point using Marketplace contract.
func (o *Opera) GetPayTokenPrice(token *common.Address, block *big.Int) (*big.Int, error) {
	return o.defaultMarketplaceContract.GetPayTokenPrice(token, block)
}
