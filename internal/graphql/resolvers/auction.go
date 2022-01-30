package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/svc"
	"artion-api-graphql/internal/types"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"time"
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

// Props provides properties of auctions running on the auction contract
func (au *Auction) Props() (*types.AuctionProps, error) {
	props, err := repository.R().GetAuctionProps(au.AuctionHall)
	if err != nil {
		log.Errorf("unable to get auction props; %s", err)
		return nil, err
	}
	return props, nil
}

// ReservePriceExceeded tells whether is the last bid greater or equal than the reserve price
func (au *Auction) ReservePriceExceeded() bool {
	if au.LastBid == nil {
		return false
	}
	return au.ReservePrice.ToInt().Cmp(au.LastBid.ToInt()) <= 0
}

// WithdrawSince tells since when can the bidder withdraw the bid
func (au *Auction) WithdrawSince() (*types.Time, error) {
	if au.LastBid == nil || au.Resolved != nil {
		return nil, nil
	}
	props, err := repository.R().GetAuctionProps(au.AuctionHall)
	if err != nil {
		return nil, err
	}
	since := time.Time(au.EndTime).Add(time.Second * 43200) // 12 hours after end
	if props.HasResultFailed && !au.ReservePriceExceeded() {
		// can be withdrawn immediately after end when bid is less than reserve
		since = time.Time(au.EndTime)
	}
	if props.Withdraw2MonthsAfterStart {
		sinceStart := time.Time(au.StartTime).Add(time.Second * 5184000) // 2 months after start
		if since.After(sinceStart) {
			since = sinceStart
		}
	}
	sinceTime := types.Time(since)
	return &sinceTime, nil
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
