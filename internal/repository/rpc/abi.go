// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import "github.com/ethereum/go-ethereum/accounts/abi"

// Erc721Abi provides access to decoded ABI of Fantom ERC-721 contract.
func (o *Opera) Erc721Abi() *abi.ABI {
	return o.abiFantom721
}

// Erc1155Abi provides access to decoded ABI of Fantom ERC-1155 contract.
func (o *Opera) Erc1155Abi() *abi.ABI {
	return o.abiFantom1155
}
