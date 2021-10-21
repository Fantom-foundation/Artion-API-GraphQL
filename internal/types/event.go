package types

type Event struct {
	Type string
	Auction *Auction
	Offer *Offer
}

type EventListener struct {
	StopChan   <-chan struct{}
	EventsChan chan Event
}
