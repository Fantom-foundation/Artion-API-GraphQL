package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/svc"
	"artion-api-graphql/internal/types"
	"context"
	"github.com/ethereum/go-ethereum/common"
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

func (rs *RootResolver) WatchAuction(ctx context.Context, args struct {
	Contract common.Address
	TokenId  hexutil.Big
}) <-chan *Auction {
	subscriber := svc.AuctionListener{
		StopChan: ctx.Done(),
		EventsChan: make(chan *types.Auction),
	}
	mgr := svc.GetSubscriptionsManager()
	mgr.SubscribeAuctionBid(args.Contract, args.TokenId, subscriber)

	// convert channel of types.Auction to channel of resolvers.Auction
	outChan := make(chan *Auction)
	go func() {
		for {
			event, more := <-subscriber.EventsChan
			if more {
				outChan <- (*Auction)(event)
			} else {
				close(outChan)
				return
			}
		}
	}()
	return outChan
}
