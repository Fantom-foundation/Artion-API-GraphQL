package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (p *Proxy) GetUnlockable(contract common.Address, tokenId big.Int) (unlockable *types.Unlockable, err error) {
	return p.shared.GetUnlockable(contract, tokenId)
}

func (p *Proxy) InsertUnlockable(unlockable *types.Unlockable) error {
	return p.shared.InsertUnlockable(unlockable)
}
