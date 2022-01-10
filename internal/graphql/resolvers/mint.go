package resolvers

import (
	"artion-api-graphql/internal/repository"
	"errors"
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
	Royalty int32
}) (*MintFeeGas, error) {
	if args.Royalty < 0 || args.Royalty > 10000 {
		return nil, errors.New("invalid royalty value - needs to be in interval from 0 to 10000")
	}
	platformFee, gas, err := repository.R().EstimateMintGas(args.User, args.Contract, args.TokenUri, uint16(args.Royalty))
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
