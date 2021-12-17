package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"time"
)

func (p *Proxy) BanNft(contract *common.Address, tokenId *hexutil.Big) error {
	return p.shared.BanNft(contract, tokenId)
}

func (p *Proxy) UnbanNft(contract *common.Address, tokenId *hexutil.Big) error {
	return p.shared.UnbanNft(contract, tokenId)
}

func (p *Proxy) ListBannedNftsAfter(after time.Time, maxAmount int64) (out []*types.BannedNft, err error) {
	return p.shared.ListBannedNftsAfter(after, maxAmount)
}

func (p *Proxy) ListBannedNfts(cursor types.Cursor, count int, backward bool) (out *types.BannedNftList, err error) {
	return p.shared.ListBannedNfts(cursor, count, backward)
}

func (p *Proxy) UpdateLastBanUpdate(time time.Time) {
	p.db.UpdateLastBanUpdate(time)
}

func (p *Proxy) LastBanUpdate() (*time.Time, error) {
	return p.db.LastBanUpdate()
}
