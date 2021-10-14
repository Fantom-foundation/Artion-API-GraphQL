package types

import "github.com/ethereum/go-ethereum/common"

// ShippingAddress represents user shipping address for tokens redeem.
type ShippingAddress struct {
	User      common.Address `bson:"user"`
	FullName  string         `bson:"fullname"`
	Phone     string         `bson:"phone"`     // for delivery
	Street    string         `bson:"street"`    // 1th line of address
	Apartment string         `bson:"apartment"` // 2th line of address
	City      string         `bson:"city"`
	State     string         `bson:"state"`
	Country   string         `bson:"country"`
	Zip       string         `bson:"zip"`
	Updated   Time           `bson:"updatedAt"`
}
