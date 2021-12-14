package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"time"
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
	Email           string          `bson:"email"`
	SiteUrl         string          `bson:"siteUrl"`
	DiscordUrl      string          `bson:"discord"`
	TelegramUrl     string          `bson:"telegram"`
	MediumUrl       string          `bson:"mediumHandle"`
	TwitterUrl      string          `bson:"twitterHandle"`
	Instagram       string          `bson:"instagramHandle"`
	IsAppropriate   bool            `bson:"isAppropriate"` // is reviewed and royalties registered on chain
	IsInternal      bool            `bson:"isInternal"` // is created using factory contract?
	IsOwnerOnly     bool            `bson:"isOwnerble"` // is only Owner allowed to mint?
	IsVerified      bool            `bson:"isVerified"` // is boosted by admin? (moderator is not sufficient)
	IsReviewed      bool            `bson:"status"` // false = in review, true = approved (removed on reject)
	AppropriateUpdate time.Time     `bson:"appropriateUpdate"` // when was "isAppropriate" changed last time?
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

// CollectionApplication is input for new LegacyCollection registration
type CollectionApplication struct {
	Contract        common.Address `json:"contract"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Royalty         json.Number    `json:"royalty"` // percents of fee
	FeeRecipient    common.Address `json:"feeRecipient"`
	Categories      []int32        `bson:"categories"`
	Discord         string         `bson:"discord"`
	Email           string         `bson:"email"`
	Telegram        string         `bson:"telegram"`
	SiteUrl         string         `bson:"siteUrl"`
	MediumHandle    string         `bson:"mediumHandle"`
	TwitterHandle   string         `bson:"twitterHandle"`
	InstagramHandle string         `bson:"instagramHandle"`
}

// DecodeCollectionApplication parses the collection registration application JSON.
func DecodeCollectionApplication(data []byte) (*CollectionApplication, error) {
	var out CollectionApplication
	err := json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (app CollectionApplication) ToCollection(image string, owner *common.Address) LegacyCollection {
	categoriesStr := make([]string, len(app.Categories))
	for i, categoryId := range app.Categories {
		categoriesStr[i] = strconv.Itoa(int(categoryId))
	}
	return LegacyCollection{
		Address:         app.Contract,
		Name:            app.Name,
		Description:     app.Description,
		CategoriesStr:   categoriesStr,
		Image:         image,
		Owner:         owner,
		FeeRecipient:  &app.FeeRecipient,
		RoyaltyValue:  app.Royalty,
		Email:         app.Email,
		SiteUrl:       app.SiteUrl,
		DiscordUrl:    app.Discord,
		TelegramUrl:   app.Telegram,
		MediumUrl:     app.MediumHandle,
		TwitterUrl:    app.TwitterHandle,
		Instagram:     app.InstagramHandle,
		IsAppropriate: false,
		IsInternal:    false,
		IsOwnerOnly:   false,
		IsVerified:    false,
		IsReviewed:    false,
	}
}
