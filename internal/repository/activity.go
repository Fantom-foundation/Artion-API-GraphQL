package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"time"
)

func (p *Proxy) StoreActivity(event *types.Activity) error {
	return p.db.StoreActivity(event)
}

func (p *Proxy) ListActivities(contract *common.Address, tokenId *hexutil.Big, user *common.Address, actTypes []types.ActivityType, cursor types.Cursor, count int, backward bool) (out *types.ActivityList, err error) {
	return p.db.ListActivities(contract, tokenId, user, actTypes, cursor, count, backward)
}

// TokenPriceHistory provides aggregation of trading prices of the token in time.
func (p *Proxy) TokenPriceHistory(contract *common.Address, tokenId *hexutil.Big, from time.Time, to time.Time) ([]types.PriceHistory, error) {
	return p.db.TokenPriceHistory(contract, tokenId, from, to)
}
