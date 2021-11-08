package resolvers

import (
	"artion-api-graphql/internal/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
)

type MintFeeGas struct {
	Error *string
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
		log.Warningf("failed to estimate mint fee/gas for %s: %s", args.Contract, err)
		errStr := err.Error()
		errStr = strings.TrimPrefix(errStr, "execution reverted: ")
		return &MintFeeGas{
			Error:       &errStr,
			PlatformFee: nil,
			Gas:         nil,
		}, nil
	}
	return &MintFeeGas{
		Error:       nil,
		PlatformFee: (*hexutil.Big)(platformFee),
		Gas:         (*hexutil.Big)(new(big.Int).SetUint64(gas)),
	}, nil
}
