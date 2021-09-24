package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p *proxy) StoreTokenEvent(event *types.TokenEvent) error {
	return p.db.StoreTokenEvent(event)
}

func (p * proxy) ListTokenEvents(nft *common.Address, tokenId *hexutil.Big, account *common.Address, cursor types.Cursor, count int, backward bool) (out *types.TokenEventList, err error) {
	return p.db.ListTokenEvents(nft, tokenId, account, cursor, count, backward)
}
