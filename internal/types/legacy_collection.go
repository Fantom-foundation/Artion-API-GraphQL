package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
)

// LegacyCollection represents token collection from old Artion.
// Keeps off-chain data about the collection.
type LegacyCollection struct {
	Address         common.Address  `bson:"erc721Address"` // unique index
	Name            string          `bson:"collectionName"`
	Description     string          `bson:"description"`
	CategoriesStr   []string        `bson:"categories"`
	Image           string          `bson:"logoImageHash"`
	Owner           *common.Address `bson:"owner"`
	FeeRecipient    *common.Address `bson:"feeRecipient"`
	RoyaltyValue    json.Number     `bson:"royalty"` // percents of fee (mostly int32, but sometime float)
	Discord         string          `bson:"discord"`
	Email           string          `bson:"email"`
	Telegram        string          `bson:"telegram"`
	SiteUrl         string          `bson:"siteUrl"`
	MediumHandle    string          `bson:"mediumHandle"`
	TwitterHandle   string          `bson:"twitterHandle"`
	IsAppropriate   bool            `bson:"isAppropriate"` // is reviewed and royalties registered on chain
	IsInternal      bool            `bson:"isInternal"` // is created using factory contract?
	IsOwnerOnly     bool            `bson:"isOwnerble"` // is only Owner allowed to mint?
	IsVerified      bool            `bson:"isVerified"`
	IsReviewed      bool            `bson:"status"` // false = in review, true = approved (removed on reject)
}

// CategoriesAsInt provides a list of category ID-s
// converted to a slice of integers instead of strings.
func (lc LegacyCollection) CategoriesAsInt() ([]int32, error) {
	var out []int32
	for _, value := range lc.CategoriesStr {
		if value == "" {
			continue
		}
		converted, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		out = append(out, int32(converted))
	}
	return out, nil
}
