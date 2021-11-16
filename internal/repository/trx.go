package repository

import "github.com/ethereum/go-ethereum/common"

// MustTransactionSender provides sender of the given transaction, if available, or empty address in other cases.
func (p *Proxy) MustTransactionSender(block common.Hash, txIndex uint) common.Address {
	return p.rpc.MustTransactionSender(block, txIndex)
}
