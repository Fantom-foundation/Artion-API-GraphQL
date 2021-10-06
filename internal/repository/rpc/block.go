// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// CurrentHead provides the ID of the latest known block.
func (o *Opera) CurrentHead() (uint64, error) {
	return o.ftm.BlockNumber(context.Background())
}

// GetHeader pulls given block header by the block number.
func (o *Opera) GetHeader(id uint64) (*eth.Header, error) {
	return o.ftm.HeaderByNumber(context.Background(), new(big.Int).SetUint64(id))
}

// BlockLogs provides list of event logs for the given block number and list of topics.
func (o *Opera) BlockLogs(blk *big.Int, topics [][]common.Hash) ([]eth.Log, error) {
	return o.ftm.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: blk,
		ToBlock:   blk,
		Topics:    topics,
	})
}
