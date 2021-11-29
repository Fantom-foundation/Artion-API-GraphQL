package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/svc"
	"artion-api-graphql/internal/types"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Auction types.Auction

// MinBidAmount is the mount required to be placed as a bid to outbid the current highest bidder.
func (au *Auction) MinBidAmount() (hexutil.Big, error) {
	val, err := repository.R().AuctionGetMinBid((*types.Auction)(au))
	if err != nil {
		return hexutil.Big{}, err
	}
	return (hexutil.Big)(*val), nil
}

// WatchAuction creates a client subscription for auction events.
func (rs *RootResolver) WatchAuction(ctx context.Context, args struct {
	Contract common.Address
	TokenId  hexutil.Big
}) <-chan Event {
	listener := types.EventListener{
		StopChan:   ctx.Done(),
		EventsChan: make(chan types.Event),
	}
	mgr := svc.GetSubscriptionsManager()
	mgr.SubscribeAuctionEvent(args.Contract, args.TokenId, listener)

	// convert channel of types.Event to channel of resolvers.Event
	outChan := make(chan Event)
	go func() {
		for {
			event, more := <-listener.EventsChan
			if more {
				outChan <- Event{event}
			} else {
				close(outChan)
				return
			}
		}
	}()
	return outChan
}
