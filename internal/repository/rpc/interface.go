// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
)

// Blockchain represents the interface defining set of high level function
// implementing interaction with the Fantom Opera blockchain.
type Blockchain interface {
	// Close closes the blockchain node connection
	Close()

	// BlockLogs provides list of event logs for the given block number and list of topics.
	BlockLogs(*common.Hash, []common.Hash) ([]eth.Log, error)

	// HeaderObserver provides a channel receiving new observed blockchain headers.
	HeaderObserver() chan *eth.Header
}
