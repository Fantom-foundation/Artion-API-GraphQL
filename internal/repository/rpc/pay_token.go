package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"
	"artion-api-graphql/internal/types"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// ListPayTokens obtains list of tokens allowed for market payments from TokenRegistry contract
func (o *Opera) ListPayTokens() (payTokens []types.PayToken, err error) {
	filterOps := bind.FilterOpts{
		Context: context.Background(),
		Start: 0,
		End: nil,
	}
	itr, err := o.tokenRegistryContract.FilterTokenAdded(&filterOps)
	if err != nil {
		return nil, err
	}
	for itr.Next() {
		address := itr.Event.Token
		payToken, err := o.getPayToken(address)
		if err != nil {
			return nil, err
		}
		payTokens = append(payTokens, payToken)
	}
	return
}

func (o *Opera) getPayToken(address common.Address) (payToken types.PayToken, err error) {
	token, err := contracts.NewErc20(address, o.ftm)
	if err != nil {
		return
	}
	payToken.Contract = address
	payToken.Name, err = token.Name(nil)
	if err != nil {
		return
	}
	payToken.Symbol, err = token.Symbol(nil)
	if err != nil {
		return
	}
	decimals, err := token.Decimals(nil)
	if err != nil {
		return
	}
	payToken.Decimals = int32(decimals)
	return
}
