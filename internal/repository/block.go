// Package repository implements persistent data access and processing.
package repository

import (
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
)

// BlockLogs provides list of event logs for the given block number and list of topics.
func (p *proxy) BlockLogs(blk *common.Hash, topics []common.Hash) ([]eth.Log, error) {
	return p.rpc.BlockLogs(blk, topics)
}

// NotifyLastObservedBlock stores information about last seen block into persistent storage
// so the API server can start where it left off thr last time.
func (p *proxy) NotifyLastObservedBlock(blk *eth.Header) {
	p.db.UpdateLastSeenBlock(blk)
}
