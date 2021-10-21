// Package repository implements persistent data access and processing.
package repository

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// IsPendingRandomNumberRequest checks if the given random number request is pending.
func (p *Proxy) IsPendingRandomNumberRequest(reqID *common.Hash, seed *common.Hash) bool {
	return p.rpc.IsPendingRandomNumberRequest(reqID, seed)
}

// FulfillRandomNumberRequest send the given random INT number into the RNG feed oracle
// as a response to the detected request.
func (p *Proxy) FulfillRandomNumberRequest(reqID *common.Hash, rnd *big.Int) error {
	return p.rpc.FulfillRandomNumberRequest(reqID, rnd)
}
