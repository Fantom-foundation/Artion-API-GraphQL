package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (o *Opera) GetBalance(address common.Address) (*big.Int, error) {
	return o.ftm.BalanceAt(context.Background(), address, nil)
}
