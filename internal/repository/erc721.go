// Package repository implements persistent data access and processing.
package repository

import "github.com/ethereum/go-ethereum/common"

const (
	// Erc721BaseInterfaceID represents an identifier of the base ERC-721 interface in ERC-165 encoding.
	Erc721BaseInterfaceID = "0x80ac58cd"

	// Erc721MetadataInterfaceID represents an identifier of the ERC-721 Metadata interface in ERC-165 encoding.
	Erc721MetadataInterfaceID = "0x5b5e139f"
)

// IsErc721Contract checks if the given address is a contract with ERC-721 interface support.
func (p *Proxy) IsErc721Contract(adr *common.Address) bool {
	return p.rpc.SupportsInterface(adr, Erc721BaseInterfaceID) && p.rpc.SupportsInterface(adr, Erc721MetadataInterfaceID)
}

// Erc721StartingBlockNumber provides the first important block number for the ERC-721 contract.
func (p *Proxy) Erc721StartingBlockNumber(adr *common.Address) (uint64, error) {
	return p.rpc.Erc721StartingBlockNumber(adr)
}
