package repository

import "github.com/ethereum/go-ethereum/accounts/abi"

// Erc721Abi provides access to decoded ABI of Fantom ERC-721 contract.
func (p *Proxy) Erc721Abi() *abi.ABI {
	return p.rpc.Erc721Abi()
}
