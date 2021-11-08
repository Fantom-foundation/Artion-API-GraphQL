package resolvers

import (
	"artion-api-graphql/internal/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type MintFeeGas struct {
	PlatformFee *hexutil.Big
	Gas *hexutil.Big
}

func (rs *RootResolver) EstimateMintFeeGas(args struct{
	User common.Address
	Contract common.Address
	TokenUri string
}) (*MintFeeGas, error) {
	platformFee, gas, err := repository.R().EstimateMintGas(args.User, args.Contract, args.TokenUri)
	if err != nil {
		return nil, err
	}
	return &MintFeeGas{
		PlatformFee: (*hexutil.Big)(platformFee),
		Gas:         (*hexutil.Big)(new(big.Int).SetUint64(gas)),
	}, nil
}
