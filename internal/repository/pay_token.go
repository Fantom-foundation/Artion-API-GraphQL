package repository

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// tokenPriceDecimalsCorrection represents the value used to reduce 1 token price to stored fixed 6 decimals - 10^12.
// Price come in 18 decimals, to preserve 6 decimals we remove 12 decimals.
var tokenPriceDecimalsCorrection = big.NewInt(1_000_000_000_000)

var zeroAddress = common.Address{}
var wFtmAddress = common.HexToAddress("0x21be370d5312f44cb42ce377bc9b8a0cef1a4c83")

// GetUnifiedPrice converts token price in pay-tokens to value in unified units for storing in database.
func (p *Proxy) GetUnifiedPrice(address common.Address, amount hexutil.Big) (out types.TokenPrice, err error) {
	out = types.TokenPrice{
		Usd:      0,
		Amount:   amount,
		PayToken: address,
	}

	// load pay token price and decimals from cache
	payToken, err := p.getPayToken(&address)
	if err != nil {
		return types.TokenPrice{}, err
	}

	// calculate price for val wei in USD in 6-decimals fixed point
	// val and unit is in 18-decimals, product is in 36 decimals - we need to remove 30 decimals to get 6-decimals
	// val is D-decimals, unit 18-decimals, product is D+18 decimals - to get 6-decimals we need to remove D+12
	out.Usd = mulTeenPower(new(big.Int).Mul(amount.ToInt(), payToken.UnitPrice), -(payToken.Decimals + 12)).Int64()
	if out.Usd < 0 {
		return out, fmt.Errorf("GetUnifiedPriceAt overflow - val: %s unit: %s decimals: %d", amount.String(), payToken.UnitPrice.String(), payToken.Decimals)
	}
	return out, nil
}

// mulTeenPower adjusts the given number by the given number of decimals.
func mulTeenPower(num *big.Int, decimals int32) *big.Int {
	if 0 == decimals {
		return num
	}

	if decimals > 0 {
		teen := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
		num.Mul(num, teen)
	} else {
		teen := new(big.Int).Exp(big.NewInt(10), big.NewInt(-int64(decimals)), nil)
		num.Div(num, teen)
	}
	return num
}

// ListPayTokens provides list of tokens allowed for market payments.
func (p *Proxy) ListPayTokens() ([]types.PayToken, error) {
	tokens, err, _ := p.callGroup.Do("ListPayTokens", func() (interface{}, error) {
		return p.cache.ListPayTokens(p.rpc.ListPayTokens)
	})
	return tokens.([]types.PayToken), err
}

func (p *Proxy) getPayToken(address *common.Address) (*types.PayToken, error) {
	list, err := p.ListPayTokens() // cached
	if err != nil {
		return nil, err
	}
	// replace zero address (native tokens) by wFTM token
	if bytes.Equal(address.Bytes(), zeroAddress.Bytes()) {
		address = &wFtmAddress
	}
	for _, payToken := range list {
		if 0 == bytes.Compare(payToken.Contract.Bytes(), address.Bytes()) {
			return &payToken, nil
		}
	}
	return nil, fmt.Errorf("unknown pay token %s", address)
}
