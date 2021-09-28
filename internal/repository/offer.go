package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p *proxy) StoreOffer(event *types.Offer) error {
	return p.db.StoreOffer(event)
}

func (p * proxy) ListOffers(nft *common.Address, tokenId *hexutil.Big, creator *common.Address, cursor types.Cursor, count int, backward bool) (out *types.OfferList, err error) {
	return p.db.ListOffers(nft, tokenId, creator, cursor, count, backward)
}
