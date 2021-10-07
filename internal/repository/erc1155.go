package repository

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Erc1155BalanceOf extracts balance of a NFT for an owner.
func (p *Proxy) Erc1155BalanceOf(contract *common.Address, tokenId *big.Int, owner *common.Address, block *big.Int) (*big.Int, error) {
	return p.rpc.Erc1155BalanceOf(contract, tokenId, owner, block)
}

// Erc1155TokenUri gets a token specific URI address from ERC-1155 contract using uri() call.
func (p *Proxy) Erc1155TokenUri(contract *common.Address, tokenId *big.Int) (string, error) {
	return p.rpc.Erc1155TokenUri(contract, tokenId)
}
