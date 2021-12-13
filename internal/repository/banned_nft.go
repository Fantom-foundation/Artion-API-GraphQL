package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p *Proxy) BanNft(contract *common.Address, tokenId *hexutil.Big) error {
	return p.shared.BanNft(contract, tokenId)
}

func (p *Proxy) UnbanNft(contract *common.Address, tokenId *hexutil.Big) error {
	return p.shared.UnbanNft(contract, tokenId)
}

func (p *Proxy) ListBannedNfts(cursor types.Cursor, count int, backward bool) (out *types.BannedNftList, err error) {
	return p.shared.ListBannedNfts(cursor, count, backward)
}
