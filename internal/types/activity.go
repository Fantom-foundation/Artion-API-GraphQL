package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
)

type ActivityType int8

const (
	EvtUnknown             ActivityType = iota // 0
	EvtListingCreated                          // 1
	EvtListingUpdated                          // 2
	EvtListingCancelled                        // 3
	EvtListingSold                             // 4
	EvtOfferCreated                            // 5
	EvtOfferCancelled                          // 6
	EvtOfferSold                               // 7
	EvtAuctionCreated                          // 8
	EvtAuctionBid                              // 9
	EvtAuctionBidWithdrawn                     // 10
	EvtAuctionCancelled                        // 11
	EvtAuctionResolved                         // 12
	EvtAuctionUpdated                          // 13
)

// Activity represents marketplace related events on tokens - when they are sold etc.
type Activity struct {
	Transaction  common.Hash  `bson:"trx"`
	OrdinalIndex int64        `bson:"index"`
	Time         Time         `bson:"time"`
	ActType      ActivityType `bson:"type"`

	// subject of trade
	Contract common.Address `bson:"contract"`
	TokenId  hexutil.Big    `bson:"token"`
	Quantity *hexutil.Big   `bson:"quantity"`

	// parties
	From common.Address  `bson:"from"`
	To   *common.Address `bson:"to"`

	// money for the token
	PayToken     *common.Address `bson:"payToken"`
	UnitPrice    *hexutil.Big    `bson:"price"`
	UnifiedPrice int64           `bson:"uprice"`

	// date/time boundaries
	StartTime *Time `bson:"startTime"`
	EndTime   *Time `bson:"endTime"`
}

// PriceHistory aggregates price history from Activity collection
type PriceHistory struct {
	Time         Time        `bson:"time"`
	UnifiedPrice json.Number `bson:"uprice"`
}

func (ph PriceHistory) Price() (hexutil.Uint64, error) {
	priceStr := ph.UnifiedPrice
	dotPos := strings.IndexByte(string(priceStr), '.')
	if dotPos != -1 {
		priceStr = priceStr[0:dotPos]
	}
	pr, err := priceStr.Int64()
	if err != nil {
		return hexutil.Uint64(0), err
	}
	return hexutil.Uint64(uint64(pr)), nil
}
