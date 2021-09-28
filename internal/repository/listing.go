package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p *proxy) AddListing(listing *types.Listing) error {
	return p.db.AddListing(listing)
}

func (p *proxy) UpdateListing(listing *types.Listing) error {
	return p.db.UpdateListing(listing)
}

func (p *proxy) RemoveListing(owner common.Address, nft common.Address, tokenId hexutil.Big) error {
	return p.db.RemoveListing(owner, nft, tokenId)
}

func (p * proxy) ListListings(nft *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (out *types.ListingList, err error) {
	return p.db.ListListings(nft, tokenId, owner, cursor, count, backward)
}
