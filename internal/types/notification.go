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
	notifyBundle    = 0b_0001_0000_0000_0000
	notifyFollower  = 0b_1000_0000_0000_0000

	NotifyNFTBurned             = /* d-626685e374d3441fa8e7a2d63090739b */ notifyBase
	NotifyNFTCreated            = /* d-2fd097843ab9480884ce8dfebe58f015 */ notifyBase | notifyAdded
	NotifyAuctionCreated        = /* d-07fb0b0e3e9b46fba27a9bbaccb12bc5 */ notifyAuction | notifyAdded
	NotifyListingCreated        = /* d-40798e32c4a64c01bf6ed8f9f44ff640 */ notifyListing | notifyAdded
	NotifyListingCanceled       = /* d-08ed9be9205d460d99e6c581e2bddb7b */ notifyListing | notifyCanceled
	NotifyBundlePurchased       = /* d-49a8868cb94a4631a05f7177af05bb44 */ notifyBundle | notifyFinished | notifyPurchased
	NotifyBundleSold            = /* d-cfd8ca0c1a904f269d1a6af3de227ca8 */ notifyBundle | notifyFinished | notifySold
	NotifyOfferOnBundleAdded    = /* d-c79239f83c7c42afb30bd80854303661 */ notifyBundle | notifyOffer | notifyAdded
	NotifyOfferOnBundleCanceled = /* d-46c783ce27924a6d8f1fae46600544c8 */ notifyBundle | notifyOffer | notifyCanceled
	NotifyAuctionBidAdded       = /* d-a2143c02992c48018325101da545a0da */ notifyAuction | notifyBid | notifyAdded
	NotifyAuctionBidCanceled    = /* d-402c6fc9a24e4d0484d5a7a1310f653c */ notifyAuction | notifyBid | notifyCanceled
	NotifyAuctionPurchased      = /* d-d6f14efa4d324b62a9c7a0dc378412af */ notifyAuction | notifyFinished
	NotifyAuctionCanceled       = /* d-6517723fea4d475090f051f20aa8575c */ notifyAuction | notifyCanceled
	NotifyOfferAdded            = /* d-8d7bfd27fb414570996e961b8db4ff23 */ notifyOffer | notifyAdded
	NotifyOfferCanceled         = /* d-d801345093f04fcda6c08012b713e9e3 */ notifyOffer | notifyCanceled
	NotifyFollowerBundleListed  = /* d-4fe28af290f74fc0ac934427518c5bb6 */ notifyFollower | notifyBundle | notifyListing | notifyAdded
	NotifyFollowerListingAdded  = /* d-10c7b86506874975aaac47efdfd19ee1 */ notifyFollower | notifyListing | notifyAdded
	NotifyFollowerAuctionAdded  = /* d-8f15781aa8ed41feae05791dd28f5de8 */ notifyFollower | notifyAuction | notifyAdded

	/*
		notifyUpdated   = 0b_0000_0000_0000_0100
		notifyPrice     = 0b_0010_0000_0000_0000

		NotifyAuctionPriceUpdated         = notifyAuction | notifyPrice | notifyUpdated
		NotifyNFTSold                     = notifySold
		NotifyNFTPurchased                = notifyPurchased
		NotifyFollowerActivity            = notifyFollower
		NotifyFollowerBundleCreated       = notifyFollower | notifyBundle | notifyAdded
		NotifyFollowerBundlePriceUpdated  = notifyFollower | notifyBundle | notifyPrice | notifyUpdated
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

	hash.Write(([]byte)(time.Time(no.TimeStamp).String()))
	return hexutils.BytesToHex(hash.Sum(nil))
}
