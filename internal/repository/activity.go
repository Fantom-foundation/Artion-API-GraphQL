package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p *Proxy) StoreActivity(event *types.Activity) error {
	return p.db.StoreActivity(event)
}

func (p *Proxy) ListActivities(contract *common.Address, tokenId *hexutil.Big, user *common.Address, cursor types.Cursor, count int, backward bool) (out *types.ActivityList, err error) {
	return p.db.ListActivities(contract, tokenId, user, cursor, count, backward)
}
