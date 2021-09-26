// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"time"
)

const (
	// reSubscribeAttemptDelay represents the delay between consecutive re-subscribe attempts
	reSubscribeAttemptDelay = 30 * time.Second
)

// observe executes the headers' observer proxy through subscription.
func (o *opera) observe() {
	var sub ethereum.Subscription
	defer func() {
		if sub != nil {
			sub.Unsubscribe()
		}
		log.Notice("block observer closed")
		o.wg.Done()
	}()

	sub = o.subscribe()
	for {
		if sub == nil {
			select {
			case <-o.sigClose:
				return
			case <-time.After(reSubscribeAttemptDelay):
				sub = o.subscribe()
				continue
			}
		}

		select {
		case <-o.sigClose:
			return
		case err := <-sub.Err():
			log.Errorf("block subscription lost; %s", err.Error())
			sub = nil
		}
	}
}

// subscribe creates a new subscription for the observer.
func (o *opera) subscribe() ethereum.Subscription {
	sub, err := o.rpc.EthSubscribe(context.Background(), o.headers, "newHeads")
	if err != nil {
		log.Criticalf("can not observe new blocks; %s", err.Error())
		return nil
	}
	return sub
}
