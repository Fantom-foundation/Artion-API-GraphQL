package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
)

// PayToken represents token supported for payments on the marketplace
type PayToken struct {
	Contract common.Address
	Name     string
	Symbol string
	Decimals int32
	UnitPrice *big.Int // 18 decimals USD
}

type TokenPrice struct {
	Usd      int64          `bson:"usd"`
	Amount   hexutil.Big    `bson:"amount"`
	PayToken common.Address `bson:"token"`
}

func (tp *TokenPrice) UsdPrice() string {
	return strconv.FormatInt(tp.Usd, 10)
}

type PriceCalcFunc func(payToken common.Address, amount hexutil.Big) (TokenPrice, error)
