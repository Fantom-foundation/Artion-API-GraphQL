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
func (o *Opera) GetPayTokenPrice(marketplace *common.Address, token *common.Address, block *big.Int) (*big.Int, error) {
	marketplaceContract := o.defaultMarketplaceContract
	if marketplace != nil {
		mc, exists := o.marketplaceContracts[*marketplace]
		if !exists {
			log.Errorf("unable to get pay token %s price - unknown marketplace contract %s",
				token.String(), marketplace.String())
		} else {
			marketplaceContract = mc
		}
	}
	return marketplaceContract.GetPayTokenPrice(token, block)
}
