package rpc

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// GetTokenRoyalty provides fee for token minter when the token is sold and its recipient (royalty has 2 decimals)
func (o *Opera) GetTokenRoyalty(contract common.Address, tokenId *big.Int) (royalty uint16, recipient common.Address, err error) {

	roy, err := o.royaltyRegistryContract.RoyaltyInfo(nil, contract, tokenId, big.NewInt(10000))
	if err != nil {
		return 0, common.Address{}, err
	}

	return uint16(roy.RoyaltyAmount.Uint64()), roy.Receiver, nil
}
