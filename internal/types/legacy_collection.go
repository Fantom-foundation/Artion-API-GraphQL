package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
)

// LegacyCollection represents token collection from old Artion.
// Keeps off-chain data about the collection.
type LegacyCollection struct {
	Address         common.Address  `bson:"erc721Address"`
	Name            string          `bson:"collectionName"`
	Description     string          `bson:"description"`
	NamedCategories []string        `bson:"categories"`
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
	IsAppropriate   bool            `bson:"isAppropriate"`
	IsInternal      bool            `bson:"isInternal"`
	IsOwnable       bool            `bson:"isOwnerble"`
	IsVerified      bool            `bson:"isVerified"`
	Status          bool            `bson:"status"`
}

// CategoriesAsInt provides a list of category ID-s
// converted to a slice of integers instead of strings.
func (lc LegacyCollection) CategoriesAsInt() ([]int32, error) {
	var out []int32
	for _, value := range lc.NamedCategories {
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
