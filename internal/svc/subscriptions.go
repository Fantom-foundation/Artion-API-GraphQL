package svc

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"sync"
)

// singleton subscription manager instance
var instance *SubscriptionsManager
var oneInstance sync.Once

// eventListenerSet represents set of listeners waiting for Events
type eventListenerSet map[types.EventListener]bool

// SubscriptionsManager provides messaging from services layer to graphql resolvers
type SubscriptionsManager struct {
	auctionListeners map[tokenMapKey]eventListenerSet
	auctionBidMutex  sync.RWMutex
	userEventListeners map[common.Address]eventListenerSet
	userEventMutex     sync.RWMutex
}

// GetSubscriptionsManager provides singleton instance of SubscriptionsManager
func GetSubscriptionsManager() *SubscriptionsManager {
	oneInstance.Do(func() {
		instance = new(SubscriptionsManager)
		instance.auctionListeners = make(map[tokenMapKey]eventListenerSet)
		instance.userEventListeners = make(map[common.Address]eventListenerSet)
	})
	return instance
}

// tokenMapKey identifies tokens (address + tokenId) in maps
type tokenMapKey [52]byte

func getTokenMapKey(contract common.Address, tokenId hexutil.Big) (out tokenMapKey) {
	copy(out[0:20], contract.Bytes())
	copy(out[20:52], (*big.Int)(&tokenId).Bytes())
	return out
}

// SubscribeAuctionEvent adds auction event listener
func (sm *SubscriptionsManager) SubscribeAuctionEvent(contract common.Address, tokenId hexutil.Big, listener types.EventListener) {
	key := getTokenMapKey(contract, tokenId)
	sm.auctionBidMutex.Lock()
	defer sm.auctionBidMutex.Unlock()

	if sm.auctionListeners[key] == nil {
		sm.auctionListeners[key] = make(eventListenerSet)
	}
	sm.auctionListeners[key][listener] = true
}

// PublishAuctionEvent sends event to all listeners watching given auction
func (sm *SubscriptionsManager) PublishAuctionEvent(event types.Event) {
	if event.Auction == nil {
		log.Errorf("auction is nil")
		return
	}
	key := getTokenMapKey(event.Auction.Contract, event.Auction.TokenId)
	sm.auctionBidMutex.Lock()
	defer sm.auctionBidMutex.Unlock()

	for listener, _ := range sm.auctionListeners[key] {
		select { // try receive
		case <-listener.StopChan: // listener context terminated
			delete(sm.auctionListeners[key], listener)
			continue // skip sending
		default:
		}
		select { // non-blocking send
		case listener.EventsChan <- event:
		default:
		}
	}
}

// SubscribeUserEvent adds user event listener
func (sm *SubscriptionsManager) SubscribeUserEvent(user common.Address, listener types.EventListener) {
	sm.userEventMutex.Lock()
	defer sm.userEventMutex.Unlock()

	if sm.userEventListeners[user] == nil {
		sm.userEventListeners[user] = make(eventListenerSet)
	}
	sm.userEventListeners[user][listener] = true
}

// PublishUserEvent sends event to all listeners watching given user
func (sm *SubscriptionsManager) PublishUserEvent(user common.Address, event types.Event) {
	sm.userEventMutex.Lock()
	defer sm.userEventMutex.Unlock()

	for listener, _ := range sm.userEventListeners[user] {
		select { // try receive
		case <-listener.StopChan: // listener context terminated
			delete(sm.userEventListeners[user], listener)
			continue // skip sending
		default:
		}
		select { // non-blocking send
		case listener.EventsChan <- event:
		default:
		}
	}
}
