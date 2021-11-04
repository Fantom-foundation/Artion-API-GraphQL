package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

func (pt PayToken) Price() (hexutil.Uint64, error) {
	price, err := repository.R().GetUnifiedUnitPrice(&pt.Contract, pt.Decimals)
	if err != nil {
		return hexutil.Uint64(0), err
	}
	return hexutil.Uint64(price), nil
}
