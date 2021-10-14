package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type Auction types.Auction

// MinBidAmount is the mount required to be placed as a bid to outbid the current highest bidder.
func (au *Auction) MinBidAmount() (hexutil.Big, error) {
	val, err := repository.R().AuctionGetMinBid(&au.Contract, (*big.Int)(&au.TokenId))
	if err != nil {
		return hexutil.Big{}, err
	}
	return (hexutil.Big)(*val), nil
}
