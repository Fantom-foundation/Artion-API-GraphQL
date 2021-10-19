// Package types provides high level structures for the API server.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"strconv"
)

// Collection represents an Artion token collection, represented by an NFT contract.
// Artion basically recognizes NFT contracts deployed form a designated factory.
type Collection struct {
	Address    common.Address `bson:"_id"`
	Type       string         `bson:"type"`
	Name       string         `bson:"name"`
	Symbol     string         `bson:"symbol"`
	Created    Time           `bson:"created"`
	Categories []int32        `bson:"categories"`
	IsActive   bool           `bson:"is_active"`
}

// LegacyCollection represents token collection from old Artion.
// Keeps off-chain data about the collection.
type LegacyCollection struct {
	Address       common.Address `bson:"erc721Address"`
	Name          string         `bson:"collectionName"`
	Description   string         `bson:"description"`
	Categories    []string       `bson:"categories"`
	Image         string         `bson:"logoImageHash"`
	Owner         common.Address `bson:"owner"`
	FeeRecipient  common.Address `bson:"feeRecipient"`
	Discord       string         `bson:"discord"`
	Email         string         `bson:"email"`
	Telegram      string         `bson:"telegram"`
	SiteUrl       string         `bson:"siteUrl"`
	MediumHandle  string         `bson:"mediumHandle"`
	TwitterHandle string         `bson:"twitterHandle"`
	IsAppropriate bool           `bson:"isAppropriate"`
	IsInternal    bool           `bson:"isInternal"`
	IsOwnerble    bool           `bson:"isOwnerble"`
	IsVerified    bool           `bson:"isVerified"`
	IsActive      bool           `bson:"status"`
}

func (lc LegacyCollection) CategoriesAsInts() ([]int32, error) {
	var out []int32
	for _, value := range lc.Categories {
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
