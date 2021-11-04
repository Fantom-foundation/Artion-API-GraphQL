package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// teenToFifteen is helper variable to initialize weiPriceDecimalsCorrection.
var teenToFifteen = big.NewInt(1_000_000_000_000_000)

// weiPriceDecimalsCorrection represents the value used to reduce 1 wei price to stored fixed 6 decimals - 10^30.
// Price come in 18 decimals, product has 36 decimals, to preserve 6 decimals we remove 30 decimals.
var weiPriceDecimalsCorrection = new(big.Int).Mul(teenToFifteen, teenToFifteen)

// tokenPriceDecimalsCorrection represents the value used to reduce 1 token price to stored fixed 6 decimals - 10^12.
// Price come in 18 decimals, to preserve 6 decimals we remove 12 decimals.
var tokenPriceDecimalsCorrection = big.NewInt(1_000_000_000_000)

// GetUnifiedPriceAt converts token price in pay-tokens to value in unified units for storing in database.
func (p *Proxy) GetUnifiedPriceAt(marketplace *common.Address, payToken *common.Address, block *big.Int, val *big.Int) int64 {
	// get price of 1 wei of payToken in USD(?) in 18-decimals fixed point
	unit, err := p.rpc.GetPayTokenPrice(marketplace, payToken, block)
	if err != nil {
		log.Warningf("unable to get price of pay token %s; %s", payToken.String(), err.Error())
		return 0
	}

	// calculate price for val wei in USD(?) in 6-decimals fixed point
	// val and unit is in 18-decimals, product is in 36 decimals - we need to remove 30 decimals to get 6-decimals
	return new(big.Int).Div(new(big.Int).Mul(val, unit), weiPriceDecimalsCorrection).Int64()
}

// GetUnifiedUnitPrice converts token price in pay-tokens to value in unified units for storing in database.
func (p *Proxy) GetUnifiedUnitPrice(payToken *common.Address, decimals int32) (uint64, error) {
	price, err, _ := p.callGroup.Do("GetUnifiedUnitPrice" + payToken.String(), func() (interface{}, error) {
		// get price of 1 wei of payToken in USD(?) in 18-decimals fixed point
		unit, err := p.rpc.GetPayTokenPrice(nil, payToken, nil)
		if err != nil {
			return nil, err
		}

		// calculate price for 1 whole token in USD(?) in 6-decimals fixed point
		// unit is in 18-decimals - we need to remove 12 decimals to get 6 decimals
		return new(big.Int).Div(unit, tokenPriceDecimalsCorrection).Uint64(), nil
	})
	return price.(uint64), err
}

// ListPayTokens provides list of tokens allowed for market payments
func (p *Proxy) ListPayTokens() ([]types.PayToken, error) {
	tokens, err, _ := p.callGroup.Do("ListPayTokens", func() (interface{}, error) {
		return p.cache.ListPayTokens(p.rpc.ListPayTokens)
	})
	return tokens.([]types.PayToken), err
}
