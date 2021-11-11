package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type ActivityType int8

const (
	EvtUnknown ActivityType = iota // 0
	EvtListingCreated // 1
	EvtListingUpdated // 2
	EvtListingCancelled // 3
	EvtListingSold // 4
	EvtOfferCreated // 5
	EvtOfferCancelled // 6
	EvtOfferSold // 7
	EvtAuctionCreated // 8
	EvtAuctionBid // 9
	EvtAuctionBidWithdrawn // 10
	EvtAuctionCancelled // 11
	EvtAuctionResolved // 12
	EvtAuctionUpdated // 13
)

// Activity represents marketplace related events on tokens - when they are sold etc.
type Activity struct {
	OrdinalIndex int64           `bson:"index"`
	Time         Time            `bson:"time"`
	ActType      ActivityType    `bson:"type"`

	// subject of trade
	Contract     common.Address  `bson:"contract"`
	TokenId      hexutil.Big     `bson:"token"`
	Quantity     *hexutil.Big    `bson:"quantity"`

	// parties
	From         common.Address  `bson:"from"`
	To           *common.Address `bson:"to"`

	// money for the token
	PayToken     *common.Address `bson:"payToken"`
	UnitPrice    *hexutil.Big    `bson:"price"`
	UnifiedPrice int64           `bson:"uprice"`

	StartTime    *Time           `bson:"startTime"`
	EndTime      *Time           `bson:"endTime"`
}
