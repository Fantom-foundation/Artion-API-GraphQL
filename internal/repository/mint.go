package repository

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// EstimateMintGas provides platform fee and gas estimation for new token minting
func (p *Proxy) EstimateMintGas(from common.Address, contract common.Address, tokenUri string, royalty uint16) (platformFee *big.Int, gas uint64, err error) {
	return p.rpc.EstimateMintGas(from, contract, tokenUri, royalty)
}
