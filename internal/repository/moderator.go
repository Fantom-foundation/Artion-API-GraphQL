package repository

import (
	"github.com/ethereum/go-ethereum/common"
)

func (p *Proxy) IsModerator(address common.Address) (isMod bool, err error) {
	key := "IsModerator-" + address.String()
	isModerator, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.shared.IsModerator(address)
	})
	return isModerator.(bool), err
}
