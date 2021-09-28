// Package repository implements persistent data access and processing.
package repository

import "github.com/ethereum/go-ethereum/common"

// ObservedContractsAddressList provides list of addresses of all observed contracts
// stored in the persistent storage.
func (p *Proxy) ObservedContractsAddressList() []common.Address {
	return p.db.ObservedContractsAddressList()
}
