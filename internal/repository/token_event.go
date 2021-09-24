package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *proxy) StoreTokenEvent(event *types.TokenEvent) error {
	return p.db.StoreTokenEvent(event)
}

func (p * proxy) ListTokenEvents(nftAddr common.Address, tokenId hexutil.Big, cursor string, count int) (out *types.TokenEventList, err error) {
	filter := bson.D{
		{Key: types.FiTokenEventNft, Value: nftAddr.String() },
		{Key: types.FiTokenEventTokenId, Value: tokenId.String() },
	}
	return p.db.ListTokenEvents(&filter, cursor, count)
}
