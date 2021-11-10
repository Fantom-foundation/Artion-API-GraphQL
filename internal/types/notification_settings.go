package types

// NotificationSettings represents settings of notifications for one user.
type NotificationSettings struct {
	// address omitted here intentionally

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
