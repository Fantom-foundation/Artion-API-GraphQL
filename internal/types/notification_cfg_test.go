package types

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestNotificationSettings_Marshal(t *testing.T) {
	ns := &NotificationSettings{
		SNotification:          false, // 1
		SBundleBuy:             true,  // 2
		SBundleSell:            false, // 4
		SBundleOffer:           true,  // 8
		SBundleOfferCancel:     false, // 16
		SNftAuctionPrice:       true,  // 32
		SNftBidToAuction:       false, // 64
		SNftBidToAuctionCancel: false, // 128
		SAuctionWin:            true,  // 256
		SAuctionOfBidCancel:    false, // 512
		SNftSell:               true,  // 1 024
		SNftBuy:                false, // 2 048
		SNftOffer:              true,  // 4 096
		SNftOfferCancel:        false, // 8 192
		FNotification:          false, // 16 384
		FBundleCreation:        true,  // 32 768
		FBundleList:            false, // 65 536
		FBundlePrice:           true,  // 131 072
		FNftAuctionPrice:       false, // 262 144
		FNftList:               true,  // 524 288
		FNftAuction:            false, // 1 048 576
		FNftPrice:              true,  // 2 097 152
	}

	// what to expect
	value := uint64(2) + 8 + 32 + 256 + 1024 + 4096 + 32768 + 131072 + 524288 + 2097152
	testData := make([]byte, 8)
	binary.BigEndian.PutUint64(testData, value)

	// calculate
	data := ns.Marshal()

	// compare
	if 0 != bytes.Compare(testData, data) {
		t.Errorf("expected %+v, received %+v", testData, data)
	}
}

func TestNotificationSettings_Unmarshal(t *testing.T) {
	// make the data set
	value := uint64(2) + 8 + 32 + 256 + 1024 + 4096 + 32768 + 131072 + 524288 + 2097152
	testData := make([]byte, 8)
	binary.BigEndian.PutUint64(testData, value)

	ns := new(NotificationSettings)
	err := ns.Unmarshal(testData)
	if err != nil {
		t.Errorf("could not unmarshal; %s", err.Error())
	}

	out := ns.Marshal()
	if 0 != bytes.Compare(testData, out) {
		t.Errorf("expected %+v, received %+v", testData, out)
	}
}

func TestNotificationSettings_IsTypeEnabled(t *testing.T) {
	value := uint64(2) + 8 + 32 + 256 + 1024 + 4096 + 32768 + 131072 + 524288 + 2097152
	testData := make([]byte, 8)
	binary.BigEndian.PutUint64(testData, value)

	ns := new(NotificationSettings)
	err := ns.Unmarshal(testData)
	if err != nil {
		t.Errorf("could not unmarshal; %s", err.Error())
	}

	toCheck := map[int32]bool{
		NotifyNFTBurned:            false,
		NotifyNFTCreated:           false,
		NotifyNFTSold:              true,
		NotifyNFTPurchased:         false,
		NotifyAuctionCreated:       false,
		NotifyListingCreated:       false,
		NotifyListingCanceled:      false,
		NotifyAuctionBidAdded:      false,
		NotifyAuctionBidCanceled:   false,
		NotifyAuctionPurchased:     true,
		NotifyAuctionCanceled:      false,
		NotifyOfferAdded:           true,
		NotifyOfferCanceled:        false,
		NotifyFollowerListingAdded: true,
		NotifyFollowerAuctionAdded: false,
	}

	for k, o := range toCheck {
		test, err := ns.IsTypeEnabled(k)
		if err != nil {
			t.Errorf("could not check notification %d; %s", k, err.Error())
		}
		if test != o {
			t.Errorf("#%d expected %+v; received %+v", k, o, test)
		}
	}
}
