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

func GetSubscriptionsManager() *SubscriptionsManager {
	oneInstance.Do(func() {
		instance = new(SubscriptionsManager)
		instance.auctionBidSubscribers = make(map[tokenMapKey]auctionListenerSet)
		instance.userEventSubscribers = make(map[common.Address]userEventListenerSet)
	})
	return instance
}

type SubscriptionsManager struct {
	auctionBidSubscribers map[tokenMapKey]auctionListenerSet
	auctionBidMutex sync.RWMutex
	userEventSubscribers map[common.Address]userEventListenerSet
	userEventMutex sync.RWMutex
}

// tokenMapKey identifies tokens (address + tokenId) in maps
type tokenMapKey [52]byte

type AuctionListener struct {
	StopChan   <-chan struct{}
	EventsChan chan *types.Auction
}

type EventListener struct {
	StopChan   <-chan struct{}
	EventsChan chan types.Event
}

// auctionListenerSet represents set of channels waiting for Auction update
type auctionListenerSet map[AuctionListener]bool

// auctionListenerSet represents set of channels waiting for Auction update
type userEventListenerSet map[EventListener]bool

func (sm *SubscriptionsManager) SubscribeAuctionBid(contract common.Address, tokenId hexutil.Big, subscriber AuctionListener) {
	key := getTokenKeyMap(contract, tokenId)
	sm.auctionBidMutex.Lock()
	defer sm.auctionBidMutex.Unlock()

	if sm.auctionBidSubscribers[key] == nil {
		sm.auctionBidSubscribers[key] = make(auctionListenerSet)
	}
	sm.auctionBidSubscribers[key][subscriber] = true
}

func (sm *SubscriptionsManager) unsubscribeAuctionBid(contract common.Address, tokenId hexutil.Big, subscriber AuctionListener) {
	key := getTokenKeyMap(contract, tokenId)
	delete(sm.auctionBidSubscribers[key], subscriber)
}

func (sm *SubscriptionsManager) PublishAuctionBid(auction *types.Auction) {
	key := getTokenKeyMap(auction.Contract, auction.TokenId)
	sm.auctionBidMutex.Lock()
	defer sm.auctionBidMutex.Unlock()

	for subscriber, _ := range sm.auctionBidSubscribers[key] {
		select { // try receive
		case <-subscriber.StopChan: // subscriber context terminated
			sm.unsubscribeAuctionBid(auction.Contract, auction.TokenId, subscriber)
			continue
		default:
		}
		select { // non-blocking send
		case subscriber.EventsChan <- auction:
		default:
		}
	}
}

func getTokenKeyMap(contract common.Address, tokenId hexutil.Big) (out tokenMapKey) {
	copy(out[0:20], contract.Bytes())
	copy(out[20:52], (*big.Int)(&tokenId).Bytes())
	return out
}

func (sm *SubscriptionsManager) SubscribeUserEvent(user common.Address, subscriber EventListener) {
	sm.userEventMutex.Lock()
	defer sm.userEventMutex.Unlock()

	if sm.userEventSubscribers[user] == nil {
		sm.userEventSubscribers[user] = make(userEventListenerSet)
	}
	sm.userEventSubscribers[user][subscriber] = true
}

func (sm *SubscriptionsManager) unsubscribeUserEvent(user common.Address, subscriber EventListener) {
	delete(sm.userEventSubscribers[user], subscriber)
}

func (sm *SubscriptionsManager) PublishUserEvent(user common.Address, event types.Event) {
	sm.userEventMutex.Lock()
	defer sm.userEventMutex.Unlock()

	for subscriber, _ := range sm.userEventSubscribers[user] {
		select { // try receive
		case <-subscriber.StopChan: // subscriber context terminated
			sm.unsubscribeUserEvent(user, subscriber)
			continue
		default:
		}
		select { // non-blocking send
		case subscriber.EventsChan <- event:
		default:
		}
	}
}
