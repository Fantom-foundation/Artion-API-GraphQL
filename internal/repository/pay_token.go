package repository

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// tokenPriceDecimalsCorrection represents the value used to reduce 1 token price to stored fixed 6 decimals - 10^12.
// Price come in 18 decimals, to preserve 6 decimals we remove 12 decimals.
var tokenPriceDecimalsCorrection = big.NewInt(1_000_000_000_000)

// GetUnifiedPriceAt converts token price in pay-tokens to value in unified units for storing in database.
func (p *Proxy) GetUnifiedPriceAt(marketplace *common.Address, payToken *common.Address, block *big.Int, val *big.Int) int64 {
	// get price of 1 wei of payToken in USD in 18-decimals fixed point
	unit, err := p.rpc.GetPayTokenPrice(marketplace, payToken, block)
	if err != nil {
		log.Warningf("unable to get price of pay token %s; %s", payToken.String(), err.Error())
		return 0
	}

	// get amount of pay-token decimals
	decimals, err := p.getPayTokenDecimals(payToken)
	if err != nil {
		log.Warningf("unable to get decimals of pay token %s; %s", payToken.String(), err.Error())
		return 0
	}

	// calculate price for val wei in USD in 6-decimals fixed point
	// val and unit is in 18-decimals, product is in 36 decimals - we need to remove 30 decimals to get 6-decimals
	// val is D-decimals, unit 18-decimals, product is D+18 decimals - to get 6-decimals we need to remove D+12
	result := mulTeenPower(new(big.Int).Mul(val, unit), -(decimals + 12)).Int64()
	log.Warningf("converting val %s to %s (%s)", val.String(), result, payToken.String())
	return result
}

// GetUnifiedUnitPrice converts token price in pay-tokens to value in unified units for storing in database.
func (p *Proxy) GetUnifiedUnitPrice(payToken *common.Address) (uint64, error) {
	price, err, _ := p.callGroup.Do("GetUnifiedUnitPrice" + payToken.String(), func() (interface{}, error) {
		// get price of 1 whole payToken in USD in 18-decimals fixed point
		unit, err := p.rpc.GetPayTokenPrice(nil, payToken, nil)
		if err != nil {
			return nil, err
		}

		// calculate price for 1 whole token in USD in 6-decimals fixed point
		// unit is in 18-decimals - we need to remove 12 decimals to get 6 decimals
		return new(big.Int).Div(unit, tokenPriceDecimalsCorrection).Uint64(), nil
	})
	return price.(uint64), err
}

func mulTeenPower(num *big.Int, decimals int32) *big.Int {
	teen := big.NewInt(10)
	if decimals > 0 {
		for i := int32(0); i < decimals; i++ {
			num.Mul(num, teen)
		}
	} else {
		for i := decimals; i < int32(0); i++ {
			num.Div(num, teen)
		}
	}
	return num
}

// ListPayTokens provides list of tokens allowed for market payments
func (p *Proxy) ListPayTokens() ([]types.PayToken, error) {
	tokens, err, _ := p.callGroup.Do("ListPayTokens", func() (interface{}, error) {
		return p.cache.ListPayTokens(p.rpc.ListPayTokens)
	})
	return tokens.([]types.PayToken), err
}

func (p *Proxy) getPayTokenDecimals(address *common.Address) (int32, error) {
	list, err := p.ListPayTokens() // cached
	if err != nil {
		return 0, err
	}
	for _, payToken := range list {
		if bytes.Equal(payToken.Contract.Bytes(), address.Bytes()) {
			return payToken.Decimals, nil
		}
	}
	return 0, fmt.Errorf("unknown pay token %s", address)
}
