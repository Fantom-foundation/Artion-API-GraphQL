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

// CalculateUnifiedPrice calculates the unified price for the given params.
func (p *Proxy) CalculateUnifiedPrice(payToken *common.Address, amount *big.Int, unitPrice *big.Int) types.TokenPrice {
	// prep the struct
	out := types.TokenPrice{
		Usd:      0,
		Amount:   hexutil.Big(*amount),
		PayToken: *payToken,
	}

	// get amount of pay-token decimals
	decimals, err := p.getPayTokenDecimals(payToken)
	if err != nil {
		log.Warningf("unable to get decimals of pay token %s; %s", payToken.String(), err.Error())
		return out
	}

	// calculate price for val wei in USD in 6-decimals fixed point
	// val and unit is in 18-decimals, product is in 36 decimals - we need to remove 30 decimals to get 6-decimals
	// val is D-decimals, unit 18-decimals, product is D+18 decimals - to get 6-decimals we need to remove D+12
	out.Usd = mulTeenPower(new(big.Int).Mul(amount, unitPrice), -(decimals + 12)).Int64()
	if out.Usd < 0 {
		log.Warningf("GetUnifiedPriceAt overflow - val: %s unit: %s decimals: %d", amount.String(), unitPrice.String(), decimals)
		out.Usd = 0
	}
	return out
}

// GetUnifiedPriceAt converts token price in pay-tokens to value in unified units for storing in database.
func (p *Proxy) GetUnifiedPriceAt(payToken *common.Address, block *big.Int, amount *big.Int) types.TokenPrice {
	// get price of 1 whole payToken in USD in 18-decimals fixed point
	unit, err := p.rpc.GetPayTokenPrice(payToken, block)
	if err != nil {
		log.Warningf("unable to get price of pay token %s; %s", payToken.String(), err.Error())
		return types.TokenPrice{
			Usd:      0,
			Amount:   hexutil.Big(*amount),
			PayToken: *payToken,
		}
	}

	return p.CalculateUnifiedPrice(payToken, amount, unit)
}

// GetUnifiedUnitPrice obtains price of pay-token in unified units (USD with 6 decimals) for storing in database.
func (p *Proxy) GetUnifiedUnitPrice(payToken *common.Address) (uint64, error) {
	price, err, _ := p.callGroup.Do("GetUnifiedUnitPrice"+payToken.String(), func() (interface{}, error) {
		// get price of 1 whole payToken in USD in 18-decimals fixed point
		unit, err := p.rpc.GetPayTokenPrice(payToken, nil)
		if err != nil {
			return nil, err
		}

		// calculate price for 1 whole token in USD in 6-decimals fixed point
		// unit is in 18-decimals - we need to remove 12 decimals to get 6 decimals
		return new(big.Int).Div(unit, tokenPriceDecimalsCorrection).Uint64(), nil
	})
	return price.(uint64), err
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

// getPayTokenDecimals provides the number of decimals of the given pay token.
func (p *Proxy) getPayTokenDecimals(address *common.Address) (int32, error) {
	list, err := p.ListPayTokens() // cached
	if err != nil {
		return 0, err
	}
	for _, payToken := range list {
		if 0 == bytes.Compare(payToken.Contract.Bytes(), address.Bytes()) {
			return payToken.Decimals, nil
		}
	}
	return 0, fmt.Errorf("unknown pay token %s", address)
}
