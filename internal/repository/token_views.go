package repository

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

func (p *Proxy) GetTokenViews(contract common.Address, tokenId big.Int) (*big.Int, error) {
	var key strings.Builder
	key.WriteString("TokenViews")
	key.WriteString(contract.String())
	key.WriteString(tokenId.String())

	views, err, _ := p.callGroup.Do(key.String(), func() (interface{}, error) {
		return p.shared.GetTokenViews(contract, tokenId)
	})
	return views.(*big.Int), err
}

func (p *Proxy) IncrementTokenViews(contract common.Address, tokenId big.Int) (err error) {
	return p.shared.IncrementTokenViews(contract, tokenId)
}
