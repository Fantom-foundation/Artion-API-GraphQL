// Package repository implements persistent data access and processing.
package repository

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const (
	// Erc1155BaseInterfaceID represents an identifier of the base ERC-1155 interface in ERC-165 encoding.
	Erc1155BaseInterfaceID = "0xd9b67a26"

	// Erc1155MetadataInterfaceID represents an identifier of the ERC-1155 Metadata interface in ERC-165 encoding.
	Erc1155MetadataInterfaceID = "0x0e89341c"
)

// Erc1155BalanceOf extracts balance of a NFT for an owner.
func (p *Proxy) Erc1155BalanceOf(contract *common.Address, tokenId *big.Int, owner *common.Address, block *big.Int) (*big.Int, error) {
	return p.rpc.Erc1155BalanceOf(contract, tokenId, owner, block)
}

// Erc1155TokenUri gets a token specific URI address from ERC-1155 contract using uri() call.
func (p *Proxy) Erc1155TokenUri(contract *common.Address, tokenId *big.Int) (string, error) {
	return p.rpc.Erc1155TokenUri(contract, tokenId)
}

// IsErc1155Contract checks if the given address is a contract with ERC-1155 interface support.
func (p *Proxy) IsErc1155Contract(adr *common.Address) bool {
	return p.rpc.SupportsInterface(adr, Erc1155BaseInterfaceID) && p.rpc.SupportsInterface(adr, Erc1155MetadataInterfaceID)
}

// Erc1155StartingBlockNumber provides the first important block number for the ERC-1155 contract.
func (p *Proxy) Erc1155StartingBlockNumber(adr *common.Address) (uint64, error) {
	return p.rpc.Erc1155StartingBlockNumber(adr)
}
