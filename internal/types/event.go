package types

type Event struct {
	Type string
	Auction *Auction
	Offer *Offer
}
