package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type PayToken types.PayToken

// PayTokens provides list of tokens supported for payments on the marketplace
func (rs *RootResolver) PayTokens() (out []PayToken, err error) {
	list, err := repository.R().ListPayTokens()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(list); i++ {
		out = append(out, PayToken(list[i]))
	}
	return out, nil
}

// Price returns current price of the token in 6-decimals USD
func (pt PayToken) Price() (hexutil.Uint64, error) {
	return hexutil.Uint64(new(big.Int).Div(pt.UnitPrice, big.NewInt(1_000_000_000_000)).Uint64()), nil
}
