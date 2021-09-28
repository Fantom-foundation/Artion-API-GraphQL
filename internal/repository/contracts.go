// Package repository implements persistent data access and processing.
package repository

import "github.com/ethereum/go-ethereum/common"

// ObservedContractsAddressList provides list of addresses of all observed contracts
// stored in the persistent storage.
func (p *Proxy) ObservedContractsAddressList() []common.Address {
	return p.db.ObservedContractsAddressList()
}

// NFTContractsTypeMap provides a map of observed contract addresses to corresponding
// contract type for ERC721 and ERC1155 contracts including their factory.
// In case of a factory contract, we need the deployed NFT type for processing.
func (p *Proxy) NFTContractsTypeMap() map[common.Address]string {
	return p.db.NFTContractsTypeMap()
}
