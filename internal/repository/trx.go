package repository

import "github.com/ethereum/go-ethereum/common"

// MustTransactionSender provides sender of the given transaction, if available, or empty address in other cases.
func (p *Proxy) MustTransactionSender(tx common.Hash) common.Address {
	return p.rpc.MustTransactionSender(tx)
}

// MustTransactionData provides sender, recipient and call data
// of the given transaction, if available, or empty slice in other cases.
func (p *Proxy) MustTransactionData(tx common.Hash) (common.Address, common.Address, []byte) {
	return p.rpc.MustTransactionData(tx)
}
