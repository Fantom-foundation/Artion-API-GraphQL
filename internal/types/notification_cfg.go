package types

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

// NotificationSettings represents settings of notifications for one user.
type NotificationSettings struct {
	// address omitted here intentionally; bool flags only

	SNotification          bool `bson:"sNotification"`          // Your Activity Notifications
	SBundleBuy             bool `bson:"sBundleBuy"`             // You have purchased a bundle.
	SBundleSell            bool `bson:"sBundleSell"`            // Your bundle is sold.
	SBundleOffer           bool `bson:"sBundleOffer"`           // You get an offer for your bundle.
	SBundleOfferCancel     bool `bson:"sBundleOfferCancel"`     // An offer to your bundle is canceled.
	SNftAuctionPrice       bool `bson:"sNftAuctionPrice"`       // The bid price to your auction is updated.
	SNftBidToAuction       bool `bson:"sNftBidToAuction"`       // You get a bid to your auction.
	SNftBidToAuctionCancel bool `bson:"sNftBidToAuctionCancel"` // A bid to your auction is canceled.
	SAuctionWin            bool `bson:"sAuctionWin"`            // You purchased a nft in auction.
	SAuctionOfBidCancel    bool `bson:"sAuctionOfBidCancel"`    // An auction you made a bid is canceled.
	SNftSell               bool `bson:"sNftSell"`               // Your nft item is sold.
	SNftBuy                bool `bson:"sNftBuy"`                // You purchased a new nft item.
	SNftOffer              bool `bson:"sNftOffer"`              // You get an offer to your nft item.
	SNftOfferCancel        bool `bson:"sNftOfferCancel"`        // An offer to your nft item is canceled.

	FNotification    bool `bson:"fNotification"`    // Follower Activity Notifications
	FBundleCreation  bool `bson:"fBundleCreation"`  // New bundle creation by follower
	FBundleList      bool `bson:"fBundleList"`      // Bundle Listing by follower
	FBundlePrice     bool `bson:"fBundlePrice"`     // Bundle AmountPrice Update by follower
	FNftAuctionPrice bool `bson:"fNftAuctionPrice"` // NFT Auction AmountPrice update by follower
	FNftList         bool `bson:"fNftList"`         // NFT Listing by follower
	FNftAuction      bool `bson:"fNftAuction"`      // New NFT Auction
	FNftPrice        bool `bson:"fNftPrice"`        // NFT AmountPrice Update by follower
}

// notificationSettingsBitMap defines map between notification setting field and its bit mask.
var notificationSettingsBitMap = map[string]uint64{
	"SNotification":          0b_0000_0000_0000_0000_0000_0000_0000_0001,
	"SBundleBuy":             0b_0000_0000_0000_0000_0000_0000_0000_0010,
	"SBundleSell":            0b_0000_0000_0000_0000_0000_0000_0000_0100,
	"SBundleOffer":           0b_0000_0000_0000_0000_0000_0000_0000_1000,
	"SBundleOfferCancel":     0b_0000_0000_0000_0000_0000_0000_0001_0000,
	"SNftAuctionPrice":       0b_0000_0000_0000_0000_0000_0000_0010_0000,
	"SNftBidToAuction":       0b_0000_0000_0000_0000_0000_0000_0100_0000,
	"SNftBidToAuctionCancel": 0b_0000_0000_0000_0000_0000_0000_1000_0000,
	"SAuctionWin":            0b_0000_0000_0000_0000_0000_0001_0000_0000,
	"SAuctionOfBidCancel":    0b_0000_0000_0000_0000_0000_0010_0000_0000,
	"SNftSell":               0b_0000_0000_0000_0000_0000_0100_0000_0000,
	"SNftBuy":                0b_0000_0000_0000_0000_0000_1000_0000_0000,
	"SNftOffer":              0b_0000_0000_0000_0000_0001_0000_0000_0000,
	"SNftOfferCancel":        0b_0000_0000_0000_0000_0010_0000_0000_0000,
	"FNotification":          0b_0000_0000_0000_0000_0100_0000_0000_0000,
	"FBundleCreation":        0b_0000_0000_0000_0000_1000_0000_0000_0000,
	"FBundleList":            0b_0000_0000_0000_0001_0000_0000_0000_0000,
	"FBundlePrice":           0b_0000_0000_0000_0010_0000_0000_0000_0000,
	"FNftAuctionPrice":       0b_0000_0000_0000_0100_0000_0000_0000_0000,
	"FNftList":               0b_0000_0000_0000_1000_0000_0000_0000_0000,
	"FNftAuction":            0b_0000_0000_0001_0000_0000_0000_0000_0000,
	"FNftPrice":              0b_0000_0000_0010_0000_0000_0000_0000_0000,
}

// Marshal converts the notification setting to a byte slice.
func (ns *NotificationSettings) Marshal() (result []byte) {
	defer func() {
		if r := recover(); r != nil {
			result = nil
		}
	}()

	var out uint64
	v := reflect.ValueOf(*ns)
	tv := v.Type()

	for i := 0; i < v.NumField(); i++ {
		// is the flag set?
		if !v.Field(i).Bool() {
			continue
		}

		key := tv.Field(i).Name
		mask, ok := notificationSettingsBitMap[key]
		if ok {
			out |= mask
		}
	}

	result = make([]byte, 8)
	binary.BigEndian.PutUint64(result, out)
	return result
}

// Unmarshal decodes slice of bytes set into the structure.
func (ns *NotificationSettings) Unmarshal(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("not enough data provided")
	}

	in := binary.BigEndian.Uint64(data[:8])

	nns := new(NotificationSettings)
	ty := reflect.TypeOf(nns).Elem()
	v := reflect.ValueOf(nns).Elem()

	num := ty.NumField()
	for i := 0; i < num; i++ {
		key := ty.Field(i).Name
		mask, ok := notificationSettingsBitMap[key]
		if ok {
			v.Field(i).SetBool((in & mask) > 0)
		}
	}

	*ns = *nns
	return nil
}

// IsTypeEnabled checks if the given notification type is enabled on the setting.
func (ns *NotificationSettings) IsTypeEnabled(nt int32) (bool, error) {
	cbt, ok := notificationCheckByType[nt]
	if !ok {
		return false, fmt.Errorf("unknown type #%d", nt)
	}
	return cbt(ns), nil
}
