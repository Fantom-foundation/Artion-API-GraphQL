// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// GetListing provides an NFT listing stored in the database, if available, err otherwise.
func (p *Proxy) GetListing(contract *common.Address, tokenID *big.Int, owner *common.Address, marketplace *common.Address) (*types.Listing, error) {
	return p.db.GetListing(contract, tokenID, owner, marketplace)
}

// StoreListing stores the given token listing into the persistent storage.
func (p *Proxy) StoreListing(lst *types.Listing) error {
	return p.db.StoreListing(lst)
}

// SetListingActive sets IsActive state of listing when token ownership changes.
func (p *Proxy) SetListingActive(contract *common.Address, tokenID *big.Int, owner *common.Address, isActive bool) error {
	return p.db.SetListingActive(contract, tokenID, owner, isActive)
}

func (p *Proxy) ListListings(nft *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (out *types.ListingList, err error) {
	return p.db.ListListings(nft, tokenId, owner, cursor, count, backward)
}
