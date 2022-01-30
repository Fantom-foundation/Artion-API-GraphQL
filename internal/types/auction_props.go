package types

type AuctionProps struct {
	WinnerCanResult            bool // can the winner result the auction instead of the seller?
	SellerCanResultUnderpriced bool // can seller result auction when topBid is less than reserve price?
	CanCancelSuccessful        bool // can seller cancel auction when topBid is more than reserve price?
	Withdraw2MonthsAfterStart  bool // can bidder withdraw the bid 2 months after the auction START?
	HasResultFailed            bool // has the contract "resultFailedAuction" method?
}

var AuctionV1Props = AuctionProps{
	WinnerCanResult:            false,
	SellerCanResultUnderpriced: false,
	CanCancelSuccessful:        true,
	Withdraw2MonthsAfterStart:  true,
	HasResultFailed:            false,
}

var AuctionV2Props = AuctionProps{
	WinnerCanResult:            true,
	SellerCanResultUnderpriced: false,
	CanCancelSuccessful:        false,
	Withdraw2MonthsAfterStart:  false,
	HasResultFailed:            true,
}

var AuctionV3Props = AuctionProps{
	WinnerCanResult:            true,
	SellerCanResultUnderpriced: true,
	CanCancelSuccessful:        false,
	Withdraw2MonthsAfterStart:  false,
	HasResultFailed:            true,
}
