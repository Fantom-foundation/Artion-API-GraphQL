// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/core/types"
)

// Repository defines interface used to interact with the persistent storage
// and the blockchain node.
type Repository interface {
	// Close terminates the repository instance.
	Close()

	// BlockLogs provides list of event logs for the given block and list of topics.
	BlockLogs(*common.Hash, []common.Hash) ([]eth.Log, error)

	// NotifyLastObservedBlock stores information about last seen block into persistent storage
	// so the API server can start where it left off thr last time.
	NotifyLastObservedBlock(*eth.Header)

	// ListTokenEvents provides a list of events filtered by NFT contract, specific token and owner's account.
	ListTokenEvents(nft *common.Address, tokenId *hexutil.Big, account *common.Address, cursor types.Cursor, count int, backward bool) (list *types.TokenEventList, err error)
}
