// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BlockLogs provides list of event logs for the given block number and list of topics.
func (o *Opera) BlockLogs(blk *common.Hash, topics []common.Hash) ([]types.Log, error) {
	return o.ftm.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: blk,
		Topics:    [][]common.Hash{topics},
	})
}

// CurrentHead provides the ID of the latest known block.
func (o *Opera) CurrentHead() (uint64, error) {
	return o.ftm.BlockNumber(context.Background())
}