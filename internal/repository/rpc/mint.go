package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// EstimateMintGas provides platform fee and gas estimation for new token minting
func (o *Opera) EstimateMintGas(user common.Address, contract common.Address, tokenUri string) (platformFee *big.Int, gas uint64, err error) {
	contr, err := contracts.NewFantomNFTTradable(contract, o.ftm)
	if err != nil {
		return nil, 0, err
	}

	// load minting fee user the collection
	platformFee, err = contr.PlatformFee(&bind.CallOpts{
		From:        user,
	})
	if err != nil {
		log.Debugf("contract %s does not define platformFee: %s", contract, err)
		platformFee = big.NewInt(0)
	}

	// construct minting transaction
	abi, err := contracts.FantomNFTTradableMetaData.GetAbi()
	if err != nil {
		return nil, 0, err
	}
	data, err := abi.Pack("mint", user, tokenUri)
	if err != nil {
		return nil, 0, err
	}

	// estimate minting gas
	gas, err = o.ftm.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  user,
		To:    &contract,
		Value: platformFee,
		Data:  data,
	})
	return platformFee, gas, err
}
