// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// RandomTrade provides the random trade structure by address, if available.
func (p *Proxy) RandomTrade(adr *common.Address) (*types.RandomTrade, error) {
	// try the cache first
	rt := p.cache.PullRandomTrade(adr)
	if nil != rt {
		return rt, nil
	}

	rt, err := p.rpc.RandomTrade(adr)
	if err != nil {
		return nil, err
	}

	p.cache.PushRandomTrade(rt)
	return rt, err
}

// RandomTradeNFTCount provides the number of tokens left in the random trading pool.
func (p *Proxy) RandomTradeNFTCount(adr *common.Address) (hexutil.Big, error) {
	return p.rpc.RandomTradeNFTCount(adr)
}

// RandomTradeNFTAvailable provides the number of tokens available for purchase in the random trading pool.
func (p *Proxy) RandomTradeNFTAvailable(adr *common.Address) (hexutil.Big, error) {
	return p.rpc.RandomTradeNFTAvailable(adr)
}

// RandomTradePrice provides the price of the trade in given ERC20 denomination.
func (p *Proxy) RandomTradePrice(trade *common.Address, pay *common.Address) (hexutil.Big, error) {
	return p.rpc.RandomTradePrice(trade, pay)
}

// RandomTradePayTokens provides a list of ERC20 tokens allowed to be used for trading.
func (p *Proxy) RandomTradePayTokens(trade *common.Address) ([]common.Address, error) {
	return p.rpc.RandomTradePayTokens(trade)
}
