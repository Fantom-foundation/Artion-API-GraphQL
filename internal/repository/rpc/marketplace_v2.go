package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// MarketplaceContractV2 is IMarketplaceContract implementation for contracts in 2.0.0 version
type MarketplaceContractV2 struct {
	marketplace *contracts.FantomMarketplaceV2
}

// GetTokenRoyalty provides fee for token minter when the token is sold and its recipient (royalty has 2 decimals)
func (mc *MarketplaceContractV2) GetTokenRoyalty(contract common.Address, tokenId *big.Int) (royalty uint16, recipient common.Address, err error) {
	var zeroAddress = common.Address{}

	// use token-level royalty if defined
	royalty, err = mc.marketplace.Royalties(nil, contract, tokenId)
	if err != nil {
		return 0, common.Address{}, err
	}
	if royalty != 0 {
		recipient, err = mc.marketplace.Minters(nil, contract, tokenId)
		if !bytes.Equal(zeroAddress.Bytes(), recipient.Bytes()) {
			return
		}
	}

	// use collection-level royalty otherwise
	collectionRoyalty, err := mc.marketplace.CollectionRoyalties(nil, contract)
	if err != nil {
		return 0, common.Address{}, err
	}
	if collectionRoyalty.Royalty != 0 && !bytes.Equal(zeroAddress.Bytes(), collectionRoyalty.FeeRecipient.Bytes()) {
		royalty = collectionRoyalty.Royalty
		recipient = collectionRoyalty.FeeRecipient
		return
	}

	// no royalty defined
	return 0, zeroAddress, nil
}

// GetPayTokenPrice extracts price of 1 whole pay token in USD in 6-decimals fixed point using Marketplace contract.
func (mc *MarketplaceContractV2) GetPayTokenPrice(token *common.Address, _ *big.Int) (*big.Int, error) {
	return mc.marketplace.GetPrice(nil, *token)
}
