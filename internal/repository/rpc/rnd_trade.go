// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// RandomTrade provides the random trade structure by address, if available.
func (o *Opera) RandomTrade(adr *common.Address) (*types.RandomTrade, error) {
	// load the contract
	con, err := contracts.NewRandomTrade(*adr, o.ftm)
	if err != nil {
		return nil, err
	}

	// prep the trade
	trade := types.RandomTrade{
		Contract:    *adr,
		Name:        getString(con.GetName),
		TradeStarts: getTimeStamp(con.GetTradeStarts),
		TradeEnds:   getTimeStamp(con.GetTradeEnds),
	}

	// pull the manager address and make sure the call succeeds
	trade.TradeManager, err = con.GetOwner(nil)
	if err != nil {
		return nil, err
	}
	return &trade, nil
}

// RandomTradeNFTCount provides the number of tokens left in the random trading pool.
func (o *Opera) RandomTradeNFTCount(adr *common.Address) (hexutil.Big, error) {
	// load the contract
	con, err := contracts.NewRandomTrade(*adr, o.ftm)
	if err != nil {
		return hexutil.Big{}, err
	}
	return getBigInt(con.GetNFTCount), nil
}

// RandomTradeNFTAvailable provides the number of tokens available for purchase in the random trading pool.
func (o *Opera) RandomTradeNFTAvailable(adr *common.Address) (hexutil.Big, error) {
	// load the contract
	con, err := contracts.NewRandomTrade(*adr, o.ftm)
	if err != nil {
		return hexutil.Big{}, err
	}
	return getBigInt(con.GetNFTAvailable), nil
}

// RandomTradePrice provides the price of the trade in given ERC20 denomination.
func (o *Opera) RandomTradePrice(trade *common.Address, pay *common.Address) (hexutil.Big, error) {
	con, err := contracts.NewRandomTrade(*trade, o.ftm)
	if err != nil {
		return hexutil.Big{}, err
	}

	val, err := con.GetPrice(nil, *pay)
	if err != nil {
		return hexutil.Big{}, err
	}

	return hexutil.Big(*val), nil
}

// getTimeStamp uses callback to extract time stamp from a contract.
func getTimeStamp(call func(*bind.CallOpts) (*big.Int, error)) types.Time {
	ts, err := call(nil)
	if err != nil {
		return types.Time{}
	}
	return (types.Time)(time.Unix(ts.Int64(), 0))
}

// getString uses callback to extract string from a contract.
func getString(call func(*bind.CallOpts) (string, error)) string {
	str, err := call(nil)
	if err != nil {
		return ""
	}
	return str
}

// getBigInt uses callback to extract string from a contract.
func getBigInt(call func(*bind.CallOpts) (*big.Int, error)) hexutil.Big {
	i, err := call(nil)
	if err != nil {
		return hexutil.Big{}
	}
	return (hexutil.Big)(*i)
}
