package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
)

func (p *proxy) StoreTokenEvent(event *types.TokenEvent) error {
	return p.db.StoreTokenEvent(event)
}

func (p * proxy) ListTokenEvents(nftAddr common.Address, tokenId big.Int, cursor string, count int) (out *types.TokenEventList, err error) {
	filter := bson.D{
		{Key: types.FiTokenEventNft, Value: nftAddr},
		{Key: types.FiTokenEventTokenId, Value: tokenId},
	}
	return p.db.ListTokenEvents(&filter, cursor, count)
}
