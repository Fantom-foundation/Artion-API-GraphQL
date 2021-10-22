// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// RandomTrade defines resolvable RandomTrade structure.
type RandomTrade types.RandomTrade

// RandomTrade resolves root query for random trade structure by contract address.
func (rs *RootResolver) RandomTrade(args struct {
	Contract common.Address
}) (*RandomTrade, error) {
	rt, err := repository.R().RandomTrade(&args.Contract)
	if err != nil {
		return nil, err
	}
	return (*RandomTrade)(rt), nil
}

// TokensAvailable resolves the number of NFTs available for trading.
func (rt *RandomTrade) TokensAvailable() (hexutil.Big, error) {
	return repository.R().RandomTradeNFTAvailable(&rt.Contract)
}

// TotalTokens resolves the total number of NFTs in the pool.
func (rt *RandomTrade) TotalTokens() (hexutil.Big, error) {
	return repository.R().RandomTradeNFTCount(&rt.Contract)
}

// PayTokens resolves a list of ERC20 tokens allowed to be used for purchase.
func (rt *RandomTrade) PayTokens() ([]common.Address, error) {
	return repository.R().RandomTradePayTokens(&rt.Contract)
}

// Price resolves the price of a token in the trade for the given denomination.
func (rt *RandomTrade) Price(args struct {
	Token common.Address
}) (hexutil.Big, error) {
	return repository.R().RandomTradePrice(&rt.Contract, &args.Token)
}
