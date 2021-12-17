// Package types provides high level structures for the API server.
package types

import (
	"crypto/sha256"
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/status-im/keycard-go/hexutils"
	"time"
)

const (
	notifyBase      = 0b_0000_0000_0000_0000
	notifyAdded     = 0b_0000_0000_0000_0001
	notifyFinished  = 0b_0000_0000_0000_1000
	notifySold      = 0b_0000_0000_0001_0000
	notifyPurchased = 0b_0000_0000_0010_0000
	notifyCanceled  = 0b_0000_0000_1000_0000
	notifyAuction   = 0b_0000_0001_0000_0000
	notifyBid       = 0b_0000_0011_0000_0000
	notifyOffer     = 0b_0000_0100_0000_0000
	notifyListing   = 0b_0000_1000_0000_0000
	notifyFollower  = 0b_1000_0000_0000_0000

	NotifyNFTBurned            = /* #0     +ok d-626685e374d3441fa8e7a2d63090739b */ notifyBase
	NotifyNFTCreated           = /* #1     +ok d-2fd097843ab9480884ce8dfebe58f015 */ notifyBase | notifyAdded
	NotifyNFTSold              = /* #16    +ok +offer +list d-f941e62f2cae4c20a64116868d531510 */ notifySold
	NotifyNFTPurchased         = /* #32    +ok +offer +list d-c41cf896b892445094171ebb9950537d */ notifyPurchased
	NotifyAuctionCreated       = /* #257   +ok d-07fb0b0e3e9b46fba27a9bbaccb12bc5 */ notifyAuction | notifyAdded
	NotifyListingCreated       = /* #2049  +ok d-40798e32c4a64c01bf6ed8f9f44ff640 */ notifyListing | notifyAdded
	NotifyListingCanceled      = /* #2176  +ok d-08ed9be9205d460d99e6c581e2bddb7b */ notifyListing | notifyCanceled
	NotifyAuctionBidAdded      = /* #769   +ok d-a2143c02992c48018325101da545a0da */ notifyAuction | notifyBid | notifyAdded
	NotifyAuctionBidCanceled   = /* #896   +ok d-402c6fc9a24e4d0484d5a7a1310f653c */ notifyAuction | notifyBid | notifyCanceled
	NotifyAuctionPurchased     = /* #264   +ok d-d6f14efa4d324b62a9c7a0dc378412af */ notifyAuction | notifyFinished
	NotifyAuctionCanceled      = /* #387   +ok d-6517723fea4d475090f051f20aa8575c */ notifyAuction | notifyCanceled
	NotifyOfferAdded           = /* #1025  +ok d-8d7bfd27fb414570996e961b8db4ff23 */ notifyOffer | notifyAdded
	NotifyOfferCanceled        = /* #1152  +ok d-d801345093f04fcda6c08012b713e9e3 */ notifyOffer | notifyCanceled
	NotifyFollowerListingAdded = /* #34817 +ok d-10c7b86506874975aaac47efdfd19ee1 */ notifyFollower | notifyListing | notifyAdded
	NotifyFollowerAuctionAdded = /* #33025 +ok d-8f15781aa8ed41feae05791dd28f5de8 */ notifyFollower | notifyAuction | notifyAdded

	/*
		notifyUpdated   = 0b_0000_0000_0000_0100
		notifyPrice     = 0b_0010_0000_0000_0000
		notifyBundle    = 0b_0001_0000_0000_0000

		NotifyAuctionPriceUpdated         = notifyAuction | notifyPrice | notifyUpdated
		NotifyBundlePurchased       	  = notifyBundle | notifyFinished | notifyPurchased
		NotifyBundleSold            	  = notifyBundle | notifyFinished | notifySold
		NotifyOfferOnBundleAdded    	  = notifyBundle | notifyOffer | notifyAdded
		NotifyOfferOnBundleCanceled 	  = notifyBundle | notifyOffer | notifyCanceled
		NotifyFollowerActivity            = notifyFollower
		NotifyFollowerBundleCreated       = notifyFollower | notifyBundle | notifyAdded
		NotifyFollowerBundlePriceUpdated  = notifyFollower | notifyBundle | notifyPrice | notifyUpdated
		NotifyFollowerBundleListed 		  = notifyFollower | notifyBundle | notifyListing | notifyAdded
		NotifyFollowerAuctionPriceUpdated = notifyFollower | notifyAuction | notifyPrice | notifyUpdated
		NotifyFollowerPriceUpdated        = notifyFollower | notifyPrice | notifyUpdated
	*/
)

// Notification represents an event notification structure.
type Notification struct {
	Type       int32           `bson:"type" json:"type"`
	Contract   *common.Address `bson:"contract" json:"contract"`
	TokenId    *hexutil.Big    `bson:"token" json:"token"`
	TimeStamp  Time            `bson:"fired" json:"fired"`
	Recipient  common.Address  `bson:"to" json:"to"`
	Originator *common.Address `bson:"from" json:"from"`
}

// notificationCheckByType defines a map of callbacks for notification settings validation by notification type
var notificationCheckByType = map[int32]func(*NotificationSettings) bool{
	NotifyNFTBurned: func(ns *NotificationSettings) bool {
		return ns.SNotification
	},
	NotifyNFTCreated: func(ns *NotificationSettings) bool {
		return ns.SNotification
	},
	NotifyNFTSold: func(ns *NotificationSettings) bool {
		return ns.SNftSell
	},
	NotifyNFTPurchased: func(ns *NotificationSettings) bool {
		return ns.SNftBuy
	},
	NotifyAuctionCreated: func(ns *NotificationSettings) bool {
		return ns.SNotification
	},
	NotifyListingCreated: func(ns *NotificationSettings) bool {
		return ns.SNotification
	},
	NotifyListingCanceled: func(ns *NotificationSettings) bool {
		return ns.SNotification
	},
	NotifyAuctionBidAdded: func(ns *NotificationSettings) bool {
		return ns.SNftBidToAuction
	},
	NotifyAuctionBidCanceled: func(ns *NotificationSettings) bool {
		return ns.SNftBidToAuctionCancel
	},
	NotifyAuctionPurchased: func(ns *NotificationSettings) bool {
		return ns.SAuctionWin
	},
	NotifyAuctionCanceled: func(ns *NotificationSettings) bool {
		return ns.SAuctionOfBidCancel
	},
	NotifyOfferAdded: func(ns *NotificationSettings) bool {
		return ns.SNftOffer
	},
	NotifyOfferCanceled: func(ns *NotificationSettings) bool {
		return ns.SNftOfferCancel
	},
	NotifyFollowerListingAdded: func(ns *NotificationSettings) bool {
		return ns.FNftList
	},
	NotifyFollowerAuctionAdded: func(ns *NotificationSettings) bool {
		return ns.FNftAuction
	},
}

// NotificationID generates unique identifier for the notification record.
func NotificationID(no *Notification) string {
	hash := sha256.New()

	var id [8]byte
	binary.BigEndian.PutUint32(id[:4], uint32(no.Type))
	hash.Write(id[:4])

	hash.Write(no.Recipient.Bytes())

	if no.Contract != nil {
		hash.Write(no.Contract.Bytes())
	}

	if no.TokenId != nil {
		hash.Write(no.TokenId.ToInt().Bytes())
	}

	if nil != no.Originator {
		hash.Write(no.Originator.Bytes())
	}

	binary.BigEndian.PutUint64(id[:], uint64(time.Time(no.TimeStamp).Unix()))
	hash.Write(id[:])

	return hexutils.BytesToHex(hash.Sum(nil))
}
