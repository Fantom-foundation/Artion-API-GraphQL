// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// MustTransactionSender provides sender of the given transaction, if available, or empty address in other cases.
func (o *Opera) MustTransactionSender(block common.Hash, txIndex uint) common.Address {
	var meta struct {
		Hash common.Hash
		From common.Address
	}
	if err := o.rpc.CallContext(context.Background(), &meta, "eth_getTransactionByBlockHashAndIndex", block, hexutil.Uint64(txIndex)); err != nil {
		log.Errorf("transaction sender unknown; %s", err.Error())
		return common.Address{}
	}

	// empty transaction hash? invalid transaction requested
	if meta.Hash == (common.Hash{}) {
		log.Errorf("wrong inclusion %s / #%d", block.String(), txIndex)
		return common.Address{}
	}
	return meta.From
}

// MustTransactionData provides call data of the given transaction, if available, or empty slice in other cases.
func (o *Opera) MustTransactionData(block common.Hash, txIndex uint) (common.Address, common.Address, []byte) {
	trx, err := o.ftm.TransactionInBlock(context.Background(), block, txIndex)
	if err != nil {
		log.Errorf("transaction detail unknown; %s", err.Error())
		return common.Address{}, common.Address{}, make([]byte, 0)
	}

	from, err := o.ftm.TransactionSender(context.Background(), trx, block, txIndex)
	if err != nil {
		log.Errorf("transaction sender unknown; %s", err.Error())
		return common.Address{}, common.Address{}, make([]byte, 0)
	}

	log.Noticef("loaded trx %s / %d [%s] from %s to %s; with %d input bytes", block.String(), txIndex, trx.Hash().String(), from.String(), trx.To().String(), len(trx.Data()))
	return from, *trx.To(), trx.Data()
}
