// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

// MustTransactionSender provides sender of the given transaction, if available, or empty address in other cases.
func (o *Opera) MustTransactionSender(tx common.Hash) common.Address {
	from, _, _ := o.MustTransactionData(tx)
	return from
}

// MustTransactionData provides call data of the given transaction, if available, or empty slice in other cases.
func (o *Opera) MustTransactionData(tx common.Hash) (common.Address, common.Address, []byte) {
	trx, pending, err := o.ftm.TransactionByHash(context.Background(), tx)
	if err != nil {
		log.Errorf("transaction %s detail unknown; %s", tx.String(), err.Error())
		return common.Address{}, common.Address{}, make([]byte, 0)
	}

	if pending {
		log.Errorf("transaction %s is pending", tx.String())
		return common.Address{}, common.Address{}, make([]byte, 0)
	}

	rec, err := o.ftm.TransactionReceipt(context.Background(), tx)
	if err != nil {
		log.Errorf("transaction receipt not available; %s", err.Error())
		return common.Address{}, common.Address{}, make([]byte, 0)
	}

	from, err := o.ftm.TransactionSender(context.Background(), trx, rec.BlockHash, rec.TransactionIndex)
	if err != nil {
		log.Errorf("transaction sender unknown; %s", err.Error())
		return common.Address{}, common.Address{}, make([]byte, 0)
	}

	log.Noticef("loaded trx %s / %d [%s] from %s to %s; with %d input bytes", rec.BlockHash.String(), rec.TransactionIndex, trx.Hash().String(), from.String(), trx.To().String(), len(trx.Data()))
	return from, *trx.To(), trx.Data()
}
