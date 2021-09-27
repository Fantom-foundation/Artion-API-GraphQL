// Package repository implements persistent data access and processing.
package repository

import (
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
)

// NewHeaders provides a channel receiving new blockchain headers.
func (p *Proxy) NewHeaders() chan *eth.Header {
	return p.rpc.NewHeaders()
}

// CurrentHead provides the ID of the latest known block.
func (p *Proxy) CurrentHead() (uint64, error) {
	return p.rpc.CurrentHead()
}

// BlockLogs provides list of event logs for the given block number and list of topics.
func (p *Proxy) BlockLogs(blk *common.Hash, topics []common.Hash) ([]eth.Log, error) {
	return p.rpc.BlockLogs(blk, topics)
}

// NotifyLastObservedBlock stores information about last seen block into persistent storage
// so the API server can start where it left off thr last time.
func (p *Proxy) NotifyLastObservedBlock(blk *eth.Header) {
	p.db.UpdateLastSeenBlock(blk)
}

// LastSeenBlockNumber pulls the ID of the last known block from the persistent storage.
func (p *Proxy) LastSeenBlockNumber() (uint64, error) {
	return p.db.LastSeenBlockNumber()
}
