package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p *Proxy) AddOffer(offer *types.Offer) error {
	return p.db.AddOffer(offer)
}

func (p *Proxy) UpdateOffer(offer *types.Offer) error {
	return p.db.UpdateOffer(offer)
}

func (p *Proxy) RemoveOffer(creator common.Address, nft common.Address, tokenId hexutil.Big) error {
	return p.db.RemoveOffer(creator, nft, tokenId)
}

func (p * Proxy) ListOffers(nft *common.Address, tokenId *hexutil.Big, creator *common.Address, cursor types.Cursor, count int, backward bool) (out *types.OfferList, err error) {
	return p.db.ListOffers(nft, tokenId, creator, cursor, count, backward)
}
