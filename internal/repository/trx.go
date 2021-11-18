package repository

import "github.com/ethereum/go-ethereum/common"

// MustTransactionSender provides sender of the given transaction, if available, or empty address in other cases.
func (p *Proxy) MustTransactionSender(block common.Hash, txIndex uint) common.Address {
	return p.rpc.MustTransactionSender(block, txIndex)
}

// MustTransactionData provides sender, recipient and call data
// of the given transaction, if available, or empty slice in other cases.
func (p *Proxy) MustTransactionData(block common.Hash, txIndex uint) (common.Address, common.Address, []byte) {
	return p.rpc.MustTransactionData(block, txIndex)
}
