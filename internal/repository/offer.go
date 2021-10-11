package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// GetOffer provides the token offer stored in the database, if available.
func (p *Proxy) GetOffer(contract *common.Address, tokenID *big.Int, proposer *common.Address) (*types.Offer, error) {
	return p.db.GetOffer(contract, tokenID, proposer)
}

// StoreOffer adds the provided offer into the database.
func (p *Proxy) StoreOffer(offer *types.Offer) error {
	return p.db.StoreOffer(offer)
}

func (p *Proxy) ListOffers(nft *common.Address, tokenId *hexutil.Big, creator *common.Address, cursor types.Cursor, count int, backward bool) (out *types.OfferList, err error) {
	return p.db.ListOffers(nft, tokenId, creator, cursor, count, backward)
}
