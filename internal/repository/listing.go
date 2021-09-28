package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p *proxy) StoreListing(event *types.Listing) error {
	return p.db.StoreListing(event)
}

func (p * proxy) ListListings(nft *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (out *types.ListingList, err error) {
	return p.db.ListListings(nft, tokenId, owner, cursor, count, backward)
}
