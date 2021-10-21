package resolvers

import (
	"artion-api-graphql/internal/svc"
	"artion-api-graphql/internal/types"
	"context"
	"github.com/ethereum/go-ethereum/common"
)

type Event struct {
	types.Event
}

func (e Event) Auction() (*Auction, error) {
	return (*Auction)(e.Event.Auction), nil
}

func (e Event) Offer() (*Offer, error) {
	return (*Offer)(e.Event.Offer), nil
}

func (rs *RootResolver) WatchUserEvents(ctx context.Context, args struct {
	User common.Address
}) <-chan Event {
	listener := types.EventListener{
		StopChan: ctx.Done(),
		EventsChan: make(chan types.Event),
	}
	mgr := svc.GetSubscriptionsManager()
	mgr.SubscribeUserEvent(args.User, listener)

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
