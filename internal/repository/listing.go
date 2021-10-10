package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// ListingGet provides an NFT listing stored in the database, if available.
func (p *Proxy) ListingGet(contract *common.Address, tokenID *big.Int, owner *common.Address) (*types.Listing, error) {
	return p.db.ListingGet(contract, tokenID, owner)
}

// ListingStore stores the given token listing into the persistent storage.
func (p *Proxy) ListingStore(lst *types.Listing) error {
	return p.db.ListingStore(lst)
}

func (p *Proxy) ListListings(nft *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (out *types.ListingList, err error) {
	return p.db.ListListings(nft, tokenId, owner, cursor, count, backward)
}
