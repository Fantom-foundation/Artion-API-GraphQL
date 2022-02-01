// Package db provides access to the persistent storage.
package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IndexDefinitionTokens provides a list of indexes expected to exist on tokens' collection.
func IndexDefinitionTokens() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 16)

	ixContractToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixContractToken}}

	ixOrdinal := "ix_ordinal"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "index", Value: -1}}, Options: &options.IndexOptions{Name: &ixOrdinal}}

	ixFulltext := "ix_fulltext"
	ix[2] = mongo.IndexModel{Keys: bson.D{
		{Key: "name", Value: "text"},
		{Key: "desc", Value: "text"},
		{Key: "symbol", Value: "text"},
	}, Options: &options.IndexOptions{Name: &ixFulltext}}

	ixCreatedBy := "ix_created_by"
	ix[3] = mongo.IndexModel{Keys: bson.D{{Key: "created_by", Value: 1}}, Options: &options.IndexOptions{Name: &ixCreatedBy}}

	// fields used for tokens list sorting - should match to sorting.TokenSorting
	ixCreated := "ix_created_index"
	ix[4] = mongo.IndexModel{Keys: bson.D{{Key: "created", Value: 1}, {Key: "index", Value: 1}}, Options: &options.IndexOptions{Name: &ixCreated}}

	ixLastList := "ix_last_list_index"
	ix[5] = mongo.IndexModel{Keys: bson.D{{Key: "last_list", Value: 1}, {Key: "index", Value: 1}}, Options: &options.IndexOptions{Name: &ixLastList}}

	ixLastTrade := "ix_last_trade_index"
	ix[6] = mongo.IndexModel{Keys: bson.D{{Key: "last_trade", Value: 1}, {Key: "index", Value: 1}}, Options: &options.IndexOptions{Name: &ixLastTrade}}

	ixAuctionUntil := "ix_auction_until_index"
	ix[7] = mongo.IndexModel{Keys: bson.D{{Key: "auction_until", Value: 1}, {Key: "index", Value: 1}}, Options: &options.IndexOptions{Name: &ixAuctionUntil}}

	ixPrice := "ix_price_index"
	ix[8] = mongo.IndexModel{Keys: bson.D{{Key: "price", Value: 1}, {Key: "index", Value: 1}}, Options: &options.IndexOptions{Name: &ixPrice}}

	ixListPrice := "ix_min_list_usd_index"
	ix[9] = mongo.IndexModel{Keys: bson.D{{Key: "min_list.usd", Value: 1}, {Key: "index", Value: 1}}, Options: &options.IndexOptions{Name: &ixListPrice}}

	ixOfferPrice := "ix_max_offer_usd_index"
	ix[10] = mongo.IndexModel{Keys: bson.D{{Key: "max_offer.usd", Value: 1}, {Key: "index", Value: 1}}, Options: &options.IndexOptions{Name: &ixOfferPrice}}

	ixLastTradeAmount := "ix_amo_trade_usd_index"
	ix[11] = mongo.IndexModel{Keys: bson.D{{Key: "amo_trade.usd", Value: 1}, {Key: "index", Value: 1}}, Options: &options.IndexOptions{Name: &ixLastTradeAmount}}

	ixViews := "ix_views_index"
	ix[12] = mongo.IndexModel{Keys: bson.D{{Key: "views", Value: 1}, {Key: "index", Value: 1}}, Options: &options.IndexOptions{Name: &ixViews}}

	ixLikes := "ix_likes_index"
	ix[13] = mongo.IndexModel{Keys: bson.D{{Key: "likes", Value: 1}, {Key: "index", Value: 1}}, Options: &options.IndexOptions{Name: &ixLikes}}

	ixPriceValid := "ix_price_valid" // prevents RAM exhaustion in TokenPriceRefreshSet
	ix[14] = mongo.IndexModel{Keys: bson.D{{Key: "price_valid", Value: 1}}, Options: &options.IndexOptions{Name: &ixPriceValid}}

	ixPriceUpdate := "ix_price_update" // prevents RAM exhaustion in TokenPriceRefreshSet
	ix[15] = mongo.IndexModel{Keys: bson.D{{Key: "price_update", Value: 1}}, Options: &options.IndexOptions{Name: &ixPriceUpdate}}

	return ix
}

// IndexDefinitionCollections provides list of indexes expected on collections.
func IndexDefinitionCollections() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 1)

	ixType := "ix_type"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "type", Value: 1}}, Options: &options.IndexOptions{Name: &ixType}}
	return ix
}

// IndexDefinitionListings provides list of indexes expected on listings.
func IndexDefinitionListings() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixToken}}

	ixOwner := "ix_owner"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "owner", Value: 1}}, Options: &options.IndexOptions{Name: &ixOwner}}
	return ix
}

// IndexDefinitionOwnership provides a list of indexes expected to exist on tokens' ownership records.
func IndexDefinitionOwnership() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixContractToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixContractToken}}

	ixOwner := "ix_owner"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "owner", Value: 1}}, Options: &options.IndexOptions{Name: &ixOwner}}
	return ix
}

// IndexDefinitionOffers provides list of indexes expected on listings.
func IndexDefinitionOffers() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixToken}}

	ixOfferedBy := "ix_offered_by"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "proposer", Value: 1}}, Options: &options.IndexOptions{Name: &ixOfferedBy}}

	ixCreated := "ix_created"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "created", Value: 1}}, Options: &options.IndexOptions{Name: &ixCreated}}
	return ix
}

// IndexDefinitionAuctions provides a list of indexes expected to exist on auctions' collection.
func IndexDefinitionAuctions() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixContractToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixContractToken}}

	ixOwner := "ix_owner"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "owner", Value: 1}}, Options: &options.IndexOptions{Name: &ixOwner}}

	ixOrdinal := "ix_ordinal"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "index", Value: -1}}, Options: &options.IndexOptions{Name: &ixOrdinal}}
	return ix
}

// IndexDefinitionAuctionBids provides a list of indexes expected to exist on auction bids' collection.
func IndexDefinitionAuctionBids() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixContractToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixContractToken}}

	ixOwner := "ix_bidder"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "bidder", Value: 1}}, Options: &options.IndexOptions{Name: &ixOwner}}
	return ix
}

// IndexDefinitionActivities provides a list of indexes expected to exist on the collection.
func IndexDefinitionActivities() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 4)

	ixContractToken := "ix_contract_token_ordinal"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}, {Key: "index", Value: -1}}, Options: &options.IndexOptions{Name: &ixContractToken}}

	ixOrdinal := "ix_ordinal"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "index", Value: -1}}, Options: &options.IndexOptions{Name: &ixOrdinal}}

	ixFrom := "ix_from_ordinal"
	ix[2] = mongo.IndexModel{Keys: bson.D{{Key: "from", Value: 1}, {Key: "index", Value: -1}}, Options: &options.IndexOptions{Name: &ixFrom}}

	ixTo := "ix_to_ordinal"
	ix[3] = mongo.IndexModel{Keys: bson.D{{Key: "to", Value: 1}, {Key: "index", Value: -1}}, Options: &options.IndexOptions{Name: &ixTo}}
	return ix
}