package rpc

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// zeroAddress represents an empty address.
var zeroAddress = common.Address{}

// GetTokenRoyalty provides fee for token minter when the token is sold and its recipient (royalty has 2 decimals)
func (o *Opera) GetTokenRoyalty(contract common.Address, tokenId *big.Int) (royalty uint16, recipient common.Address, err error) {

	// use token-level royalty if defined
	royalty, err = o.marketplace.Royalties(nil, contract, tokenId)
	if err != nil {
		return 0, common.Address{}, err
	}
	if royalty != 0 {
		recipient, err = o.marketplace.Minters(nil, contract, tokenId)
		if !bytes.Equal(zeroAddress.Bytes(), recipient.Bytes()) {
			return
		}
	}

	// use collection-level royalty otherwise
	collectionRoyalty, err := o.marketplace.CollectionRoyalties(nil, contract)
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
