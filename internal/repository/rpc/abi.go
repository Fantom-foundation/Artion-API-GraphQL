// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import "github.com/ethereum/go-ethereum/accounts/abi"

// Erc721Abi provides access to decoded ABI of Fantom ERC-721 contract.
func (o *Opera) Erc721Abi() *abi.ABI {
	return o.abiFantom721
}
