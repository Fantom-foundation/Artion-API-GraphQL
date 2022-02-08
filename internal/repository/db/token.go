package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
	"strings"
	"time"
)

const (
	// CoTokens is the name of database collection.
	coTokens = "tokens"

	// fiTokenContract is the column storing the address of the NFT token contract.
	fiTokenContract = "contract"

	// fiTokenIsActive is the column storing the NFT token activity mark.
	fiTokenIsActive = "is_active"

	// fiTokenIsBanned is the column storing the NFT token banned mark.
	fiTokenIsBanned = "is_banned"

	// fiTokenIsColBanned is the column storing the NFT token collection banned mark.
	fiTokenIsColBanned = "is_col_banned"

	// fiTokenMetadataURI is the column storing the NFT token metadata URI.
	fiTokenMetadataURI = "uri"

	// FiTokenName is the column storing the name of the NFT token.
	fiTokenName = "name"

	// fiTokenDescription is the column storing the description of the NFT token.
	fiTokenDescription = "desc"

	// fiTokenSymbol is the column storing symbol of the token
	fiTokenSymbol = "symbol"

	// fiTokenIpRights is the column storing URL of IP document
	fiTokenIpRights = "ip_rights"

	// fiTokenImageURI is the column storing the image URI of the NFT token.
	fiTokenImageURI = "image"

	// fiTokenImageType is the column storing type of the image of the NFT token.
	fiTokenImageType = "image_type"

	// fiTokenCreatedBy is the column marking creator of the token.
	fiTokenCreatedBy = "created_by"

	// fiTokenHasListingSince is the column marking listed token earliest date/time.
	fiTokenHasListingSince = "listed_since"

	// fiTokenHasAuctionSince is the column marking the earliest time a token is auctioned.
	fiTokenHasAuctionSince = "auction_since"

	// fiTokenHasAuctionUntil is the column marking the latest time a token is auctioned.
	fiTokenHasAuctionUntil = "auction_until"

	// fiTokenHasOfferUntil is the column marking offered token latest date/time.
	fiTokenHasOfferUntil = "offer_until"

	// fiTokenHasBids is the column marking auctioned token with at least one bid.
	fiTokenHasBids = "has_bid"

	// fiTokenLastTrade is the column storing the last trade date/time.
	fiTokenLastTrade = "last_trade"

	// fiTokenLastListing is the column storing the last listing creation date/time.
	fiTokenLastListing = "last_list"

	// fiTokenLastOffer is the column storing the last offer creation date/time.
	fiTokenLastOffer = "last_offer"

	// fiTokenLastAuction is the column storing the latest auction creation date/time.
	fiTokenLastAuction = "last_auction"

	// fiTokenLastBid is the column storing the latest auction bid date/time.
	fiTokenLastBid = "last_bid"

	// fiTokenAmountLastTrade is the column storing the last price token was sold for.
	fiTokenAmountLastTrade = "amo_trade"

	// fiTokenAmountLastOffer is the column storing the last price token was sold for.
	fiTokenAmountLastOffer = "amo_offer"

	// fiTokenAmountLastBid is the column storing the last bid placed on the token.
	fiTokenAmountLastBid = "amo_bid"

	// fiTokenAmountLastList is the column storing the latest price token is listed at.
	fiTokenAmountLastList = "amo_list"

	// fiTokenReservePrice is the column storing reserve price of running auction.
	fiTokenReservePrice = "reserve"

	// fiTokenMinListPrice is the column storing minimal listing price.
	fiTokenMinListPrice = "min_list"

	// fiTokenListUsdPrice is the column storing listed price of token in USD aggregated from listings.
	fiTokenListUsdPrice = "min_list.usd"

	// fiTokenOfferUsdPrice is the column storing offered price of token in USD aggregated from listings.
	fiTokenOfferUsdPrice = "max_offer.usd"

	// fiTokenMaxOfferPrice is the column storing maximal offer price.
	fiTokenMaxOfferPrice = "max_offer"

	// fiTokenPrice is the column storing price of token in USD aggregated from listings and auctions.
	fiTokenPrice = "price"

	// fiTokenPriceValid is the column storing end of price validity, when the price should be updated
	fiTokenPriceValid = "price_valid"

	// fiTokenPriceUpdate is the column storing time of the last price update
	fiTokenPriceUpdate = "price_update"

	// fiTokenCategories is the column storing categories ids of the token.
	fiTokenCategories = "categories"

	// fiTokenRoyalty is the column storing token royalty.
	fiTokenRoyalty = "royalty"

	// fiTokenFeeRecipient is the column storing recipient of the royalty.
	fiTokenFeeRecipient = "fee_recipient"

	// fiTokenCachedLikes is the column storing amount of likes synced from shared db.
	fiTokenCachedLikes = "likes"

	// fiTokenCachedViews is the column storing amount of views synced from shared db.
	fiTokenCachedViews = "views"

	// fiTokenLikesUpdate is the column storing time of last CachedLikes/CachedViews update.
	fiTokenLikesUpdate = "likes_update"

	// fiTokenMetadataUpdate is the column storing the time
	// of the metadata update schedule of the NFT token.
	fiTokenMetadataUpdate = "meta_update"

	// fiTokenMetadataUpdate is the column storing the time
	// of the metadata update schedule of the NFT token.
	fiTokenMetadataUpdateFailures = "meta_failures"
)

// GetToken loads specific NFT token for the given contract address and token ID
func (db *MongoDbBridge) GetToken(contract *common.Address, tokenId *big.Int) (token *types.Token, err error) {
	col := db.client.Database(db.dbName).Collection(coTokens)
	result := col.FindOne(context.Background(), bson.D{{Key: fieldId, Value: types.TokenID(contract, tokenId)}})

	var row types.Token
	if err = result.Decode(&row); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Errorf("can not decode token %s at %s; %s", tokenId.String(), contract.String(), err.Error())
		return nil, err
	}
	return &row, err
}

// TokenKnown checks if the given token exists i the database.
func (db *MongoDbBridge) TokenKnown(contract *common.Address, tokenId *big.Int) bool {
	col := db.client.Database(db.dbName).Collection(coTokens)

	res := col.FindOne(
		context.Background(),
		bson.D{{Key: fieldId, Value: types.TokenID(contract, tokenId)}},
		options.FindOne().SetProjection(bson.D{{Key: fieldId, Value: 1}}),
	)
	if res.Err() != nil {
		return false
	}
	return true
}

// StoreToken inserts new NFT token or updates existing token in persistent database.
func (db *MongoDbBridge) StoreToken(token *types.Token) error {
	if token == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coTokens)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	// try to do the insert
	id := token.ID()
	rs, err := col.UpdateOne(
		ctx,
		bson.D{{Key: fieldId, Value: id}},
		bson.D{
			{Key: "$set", Value: token},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fieldId, Value: id},
			}},
		},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Errorf("can not store token %s at %s; %s", token.TokenId.String(), token.Contract.String(), err.Error())
		return err
	}
	if rs.UpsertedCount > 0 {
		log.Infof("token %s on contract %s added to database", token.TokenId.String(), token.Contract.String())
	}
	return nil
}

// UpdateToken updates the token data i the database using provided update data set.
func (db *MongoDbBridge) UpdateToken(contract *common.Address, tokenID *big.Int, data bson.M) error {
	// get the collection
	col := db.client.Database(db.dbName).Collection(coTokens)
	rs, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: fieldId, Value: types.TokenID(contract, tokenID)}},
		bson.D{{Key: "$set", Value: data}},
	)
	if err != nil {
		log.Errorf("can not update token %s/%s; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err.Error())
		return err
	}
	if rs.UpsertedCount > 0 {
		log.Infof("token %s/%s updated", contract.String(), (*hexutil.Big)(tokenID).String())
	}
	return nil
}

// UpdateTokenMetadata updates basic metadata of the NFT token.
func (db *MongoDbBridge) UpdateTokenMetadata(nft *types.Token) error {
	if nft == nil {
		return fmt.Errorf("no value to store")
	}

	return db.UpdateToken(&nft.Contract, (*big.Int)(&nft.TokenId), bson.M{
		fiTokenName: nft.Name,
		fiTokenDescription: nft.Description,
		fiTokenImageURI: nft.ImageURI,
		fiTokenImageType: nft.ImageType,
		fiTokenIpRights: nft.IpRights,
		fiTokenSymbol: nft.Symbol,
		fiTokenCategories: nft.Categories,
		fiTokenRoyalty: nft.Royalty,
		fiTokenFeeRecipient: nft.FeeRecipient,
		fiTokenMetadataUpdate: nft.MetaUpdate,
		fiTokenMetadataUpdateFailures: nft.MetaFailures,
		fiTokenIsActive: nft.IsActive,
	})
}

// UpdateTokenMetadataRefreshSchedule sets the NFT metadata update schedule time.
func (db *MongoDbBridge) UpdateTokenMetadataRefreshSchedule(nft *types.Token) error {
	if nft == nil {
		return fmt.Errorf("no value to store")
	}

	return db.UpdateToken(&nft.Contract, (*big.Int)(&nft.TokenId), bson.M{
		fiTokenMetadataUpdate: nft.MetaUpdate,
		fiTokenMetadataUpdateFailures: nft.MetaFailures,
	})
}

// TokenMarkOffered marks the given NFT as having offer for the given price.
func (db *MongoDbBridge) TokenMarkOffered(contract common.Address, tokenID hexutil.Big, price types.TokenPrice, ts *time.Time, priceCalc types.PriceCalcFunc) error {
	t, err := db.GetToken(&contract, tokenID.ToInt())
	if err != nil {
		log.Errorf("unable to load token %s/%s; %s", contract.String(), tokenID.String(), err)
		return err
	}
	if t == nil {
		log.Criticalf("unknown token %s/%s", contract.String(), tokenID.String())
		return nil
	}

	return db.updateTokenAndRecalcPrice(t, bson.M{
		fiTokenLastOffer: *ts,
		fiTokenHasOfferUntil: db.OpenOfferUntil(&contract, tokenID.ToInt()),
		fiTokenAmountLastOffer: price,
	}, priceCalc)
}

// TokenMarkListed marks the given NFT as listed for direct sale for the given price.
func (db *MongoDbBridge) TokenMarkListed(contract common.Address, tokenID hexutil.Big, price types.TokenPrice, ts *time.Time, priceCalc types.PriceCalcFunc) error {
	t, err := db.GetToken(&contract, tokenID.ToInt())
	if err != nil {
		log.Errorf("unable to load token %s/%s; %s", contract.String(), tokenID.String(), err)
		return err
	}
	if t == nil {
		log.Criticalf("unknown token %s/%s", contract.String(), tokenID.String())
		return nil
	}

	t.HasListingSince = db.OpenListingSince(&contract, tokenID.ToInt())

	return db.updateTokenAndRecalcPrice(t, bson.M{
		fiTokenLastListing: *ts,
		fiTokenHasListingSince: t.HasListingSince,
		fiTokenAmountLastList: price,
	}, priceCalc)
}

// TokenMarkAuctioned marks the given NFT as auctioned for the given price.
func (db *MongoDbBridge) TokenMarkAuctioned(contract common.Address, tokenID hexutil.Big, reservePrice types.TokenPrice, ts *time.Time, priceCalc types.PriceCalcFunc) error {
	t, err := db.GetToken(&contract, tokenID.ToInt())
	if err != nil {
		log.Errorf("unable to load token %s/%s; %s", contract.String(), tokenID.String(), err)
		return err
	}
	if t == nil {
		log.Criticalf("unknown token %s/%s", contract.String(), tokenID.String())
		return nil
	}

	auctionSince, auctionUntil := db.OpenAuctionRange(&contract, tokenID.ToInt())
	return db.updateTokenAndRecalcPrice(t, bson.M{
		fiTokenLastAuction: *ts,
		fiTokenHasAuctionSince: auctionSince,
		fiTokenHasAuctionUntil: auctionUntil,
		fiTokenHasBids: false,
		fiTokenReservePrice: reservePrice,
	}, priceCalc)
}

// TokenMarkBid marks the given NFT as having auction bid for the given price.
func (db *MongoDbBridge) TokenMarkBid(contract common.Address, tokenID hexutil.Big, bidPrice types.TokenPrice, ts *time.Time, priceCalc types.PriceCalcFunc) error {
	t, err := db.GetToken(&contract, tokenID.ToInt())
	if err != nil {
		log.Errorf("unable to load token %s/%s; %s", contract.String(), tokenID.String(), err)
		return err
	}
	if t == nil {
		log.Criticalf("unknown token %s/%s", contract.String(), tokenID.String())
		return nil
	}

	return db.updateTokenAndRecalcPrice(t, bson.M{
		fiTokenHasBids: true,
		fiTokenAmountLastBid: bidPrice,
		fiTokenLastBid: ts,
	}, priceCalc)
}

// TokenMarkUnlisted marks the given NFT as listed for direct sale for the given price.
func (db *MongoDbBridge) TokenMarkUnlisted(contract common.Address, tokenID hexutil.Big, priceCalc types.PriceCalcFunc) error {
	t, err := db.GetToken(&contract, tokenID.ToInt())
	if err != nil {
		log.Errorf("unable to load token %s/%s; %s", contract.String(), tokenID.String(), err)
		return err
	}
	if t == nil {
		log.Criticalf("unknown token %s/%s", contract.String(), tokenID.String())
		return nil
	}

	hasListingSince := db.OpenListingSince(&contract, tokenID.ToInt())

	return db.updateTokenAndRecalcPrice(t, bson.M{
		fiTokenHasListingSince: hasListingSince,
	}, priceCalc)
}

// TokenMarkUnOffered marks the given NFT as not having buy offers.
func (db *MongoDbBridge) TokenMarkUnOffered(contract common.Address, tokenID hexutil.Big, priceCalc types.PriceCalcFunc) error {
	t, err := db.GetToken(&contract, tokenID.ToInt())
	if err != nil {
		log.Errorf("unable to load token %s/%s; %s", contract.String(), tokenID.String(), err)
		return err
	}
	if t == nil {
		log.Criticalf("unknown token %s/%s", contract.String(), tokenID.String())
		return nil
	}

	return db.updateTokenAndRecalcPrice(t, bson.M{
		fiTokenHasOfferUntil: db.OpenOfferUntil(&contract, tokenID.ToInt()),
	}, priceCalc)
}

// TokenMarkUnAuctioned marks the given NFT as not having an active auction on.
func (db *MongoDbBridge) TokenMarkUnAuctioned(contract common.Address, tokenID hexutil.Big, priceCalc types.PriceCalcFunc) error {
	t, err := db.GetToken(&contract, tokenID.ToInt())
	if err != nil {
		log.Errorf("unable to load token %s/%s; %s", contract.String(), tokenID.String(), err)
		return err
	}
	if t == nil {
		log.Criticalf("unknown token %s/%s", contract.String(), tokenID.String())
		return nil
	}

	hasAuctionSince, hasAuctionUntil := db.OpenAuctionRange(&contract, tokenID.ToInt())
	t.HasBids = false

	return db.updateTokenAndRecalcPrice(t, bson.M{
		fiTokenHasAuctionSince: hasAuctionSince,
		fiTokenHasAuctionUntil: hasAuctionUntil,
		fiTokenHasBids: false,
	}, priceCalc)
}

// TokenMarkUnBid marks the given NFT as not having a bid anymore.
func (db *MongoDbBridge) TokenMarkUnBid(contract common.Address, tokenID hexutil.Big, priceCalc types.PriceCalcFunc) error {
	t, err := db.GetToken(&contract, tokenID.ToInt())
	if err != nil {
		log.Errorf("unable to load token %s/%s; %s", contract.String(), tokenID.String(), err)
		return err
	}
	if t == nil {
		log.Criticalf("unknown token %s/%s", contract.String(), tokenID.String())
		return nil
	}

	return db.updateTokenAndRecalcPrice(t, bson.M{
		fiTokenHasBids: false,
	}, priceCalc)
}

// TokenMarkSold marks the given NFT as transferred OR sold on a listing/offer/auction for the given price.
func (db *MongoDbBridge) TokenMarkSold(contract common.Address, tokenID hexutil.Big, price *types.TokenPrice, tradeTime *time.Time, priceCalc types.PriceCalcFunc) error {
	t, err := db.GetToken(&contract, tokenID.ToInt())
	if err != nil {
		log.Errorf("unable to load token %s/%s; %s", contract.String(), tokenID.String(), err)
		return err
	}
	if t == nil {
		log.Criticalf("unknown token %s/%s", contract.String(), tokenID.String())
		return nil
	}

	t.HasAuctionSince, t.HasAuctionUntil = db.OpenAuctionRange(&contract, tokenID.ToInt())
	t.HasListingSince = db.OpenListingSince(&contract, tokenID.ToInt())
	t.HasOfferUntil = db.OpenOfferUntil(&contract, tokenID.ToInt())

	if price != nil { // is this trade? (not free transfer only)
		t.LastTrade = (*types.Time)(tradeTime)
		t.AmountLastTrade = *price
	}

	return db.updateTokenAndRecalcPrice(t, bson.M{
		fiTokenLastTrade: t.LastTrade,
		fiTokenAmountLastTrade: t.AmountLastTrade,
		fiTokenHasListingSince: t.HasListingSince,
		fiTokenHasOfferUntil: t.HasOfferUntil,
		fiTokenHasAuctionSince: t.HasAuctionSince,
		fiTokenHasAuctionUntil: t.HasAuctionUntil,
		fiTokenHasBids: false,
	}, priceCalc)
}

func (db *MongoDbBridge) TokenMarkBanned(contract *common.Address, tokenID *big.Int, banned bool) error {
	return db.UpdateToken(contract, tokenID, bson.M{
		fiTokenIsBanned: banned,
	})
}

func (db *MongoDbBridge) TokenMarkCollectionBanned(contract *common.Address, banned bool) error {
	col := db.client.Database(db.dbName).Collection(coTokens)
	rs, err := col.UpdateMany(
		context.Background(),
		bson.D{{Key: fiTokenContract, Value: contract}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: fiTokenIsColBanned, Value: banned},
		}}},
	)
	if err != nil {
		log.Errorf("can not update tokens of collection %s; %s", contract.String(), err.Error())
		return err
	}
	if rs.UpsertedCount > 0 {
		log.Infof("tokens of %s updated", contract.String())
	}
	return nil
}

// TokenMetadataRefreshSet pulls s set of NFT tokens scheduled to be updated up to this time.
func (db *MongoDbBridge) TokenMetadataRefreshSet(setSize int64) ([]*types.Token, error) {
	list := make([]*types.Token, setSize)
	col := db.client.Database(db.dbName).Collection(coTokens)

	// load the set from database
	cur, err := col.Find(
		context.Background(),
		bson.D{
			{Key: fiTokenMetadataUpdate, Value: bson.D{{"$lt", time.Now()}}},
			{Key: fiTokenMetadataURI, Value: bson.D{{"$ne", ""}}},
		},
		options.Find().SetSort(bson.D{{Key: fiTokenMetadataUpdate, Value: 1}}).SetLimit(setSize),
	)
	if err != nil {
		log.Errorf("can not pull metadata refresh set; %s", err.Error())
		return nil, err
	}
	defer closeFindCursor("TokenMetadataRefreshSet", cur)

	var i int
	for cur.Next(context.Background()) {
		var row types.Token
		if err := cur.Decode(&row); err != nil {
			log.Errorf("can not decode Token; %s", err.Error())
			return nil, err
		}
		list[i] = &row
		i++
	}
	return list[:i], nil
}

// TokenPriceRefreshSet pulls s set of tokens scheduled to be their price updated.
func (db *MongoDbBridge) TokenPriceRefreshSet(setSize int) ([]*types.Token, error) {
	list := make([]*types.Token, setSize)
	now := time.Now()

	// load the set of expired prices (price changes because of timed event like end of auction)
	count, err := db.addTokensIntoRefreshSet(
		bson.M{
			fiTokenPriceValid: bson.M{ "$lt": now },
		},
		options.Find().SetSort(bson.D{{Key: fiTokenPriceValid, Value: 1}}).SetLimit(int64(setSize)),
		list,
		)
	if err != nil {
		log.Errorf("failed to list invalid prices tokens; %s", err.Error())
		return nil, err
	}

	// when exhausted expired prices, complete the set with the oldest prices (regular exchange rate update)
	count2 := 0
	if count < setSize {
		count2, err = db.addTokensIntoRefreshSet(
			bson.M{	"$or": bson.A{
				bson.M{	fiTokenPriceValid: bson.M{ "$gte": now } },
				bson.M{	fiTokenPriceValid: bson.M{ "$type": 10 } },
				},
			},
			options.Find().SetSort(bson.D{{Key: fiTokenPriceUpdate, Value: 1}}).SetLimit(int64(setSize - count)),
			list[count:],
		)
		if err != nil {
			log.Errorf("failed to list old prices tokens; %s", err.Error())
			return nil, err
		}
	}

	return list[:count+count2], nil
}

func (db *MongoDbBridge) addTokensIntoRefreshSet(filter bson.M, opts *options.FindOptions, list []*types.Token) (int, error) {
	col := db.client.Database(db.dbName).Collection(coTokens)
	cur, err := col.Find(context.Background(), filter, opts)
	if err != nil {
		log.Errorf("can not pull refresh set; %s", err)
		return 0, err
	}
	defer closeFindCursor("addIntoRefreshSet", cur)

	var i int
	for cur.Next(context.Background()) {
		var row types.Token
		if err = cur.Decode(&row); err != nil {
			log.Errorf("can not decode Token; %s", err.Error())
			return i, err
		}
		list[i] = &row
		i++
	}
	return i, nil
}

// TokenPriceRefresh recalculates token prices and updates them in database.
func (db *MongoDbBridge) TokenPriceRefresh(t *types.Token, priceCalc types.PriceCalcFunc) (err error) {
	update, err := db.getTokenPriceUpdate(t, priceCalc)
	if err != nil {
		return err
	}
	return db.UpdateToken(&t.Contract, t.TokenId.ToInt(), update)
}

// updateTokenAndRecalcPrice updates the token with given data and refresh the token prices
func (db *MongoDbBridge) updateTokenAndRecalcPrice(t *types.Token, data bson.M, priceCalc types.PriceCalcFunc) error {
	update, err := db.getTokenPriceUpdate(t, priceCalc)
	if err != nil {
		return fmt.Errorf("unable to refresh price of token %s/%s; %s", t.Contract.String(), t.TokenId.String(), err)
	}
	// merge update set from parameter with price update set
	for key, value := range data {
		update[key] = value
	}
	return db.UpdateToken(&t.Contract, t.TokenId.ToInt(), update)
}

// getTokenPriceUpdate provides mongo-update for all token price fields
func (db *MongoDbBridge) getTokenPriceUpdate(t *types.Token, priceCalc types.PriceCalcFunc) (update bson.M, err error) {

	minListPrice, err := db.MinListingPrice(t.Contract, t.TokenId, priceCalc)
	if err != nil {
		log.Errorf("obtaining listing price of %s/%s failed; %s", t.Contract.String(), t.TokenId.String(), err)
	}
	minListValidity, err := db.MinListingPriceValidity(t.Contract, t.TokenId) // when starts next token listing
	if err != nil {
		log.Errorf("obtaining listing price validity of %s/%s failed; %s", t.Contract.String(), t.TokenId.String(), err)
	}
	maxOfferPrice, maxOfferValidity, err := db.MaxOfferPrice(t.Contract, t.TokenId, priceCalc)
	if err != nil {
		log.Errorf("obtaining offer price of %s/%s failed; %s", t.Contract.String(), t.TokenId.String(), err)
	}
	auctionPrice, bidPrice, reservePrice, auctionValidity, err := db.AuctionPrice(t.Contract, t.TokenId, priceCalc)
	if err != nil {
		log.Errorf("obtaining auction price of %s/%s failed; %s", t.Contract.String(), t.TokenId.String(), err)
	}
	lastTradePrice, err := priceCalc(t.AmountLastTrade.PayToken, t.AmountLastTrade.Amount)
	if err != nil {
		log.Errorf("obtaining last trade price of %s/%s failed; %s", t.Contract.String(), t.TokenId.String(), err)
	}

	// one aggregated price for general sorting
	var tokenPrice types.TokenPrice
	var priceValidity *types.Time

	// has auction - use auction top bid (or reserve price)
	if auctionValidity != nil { // reserve price can be zero - need to use validity in the condition
		tokenPrice = auctionPrice
		priceValidity = auctionValidity
	}

	// has listing and the listing is cheaper than auction
	if minListPrice.Usd != 0 && (tokenPrice.Usd == 0 || tokenPrice.Usd > minListPrice.Usd) {
		tokenPrice = minListPrice
	}

	// if validity from listing is shorter, use it
	if minListValidity != nil && (priceValidity == nil || (*time.Time)(priceValidity).After(time.Time(*minListValidity))) {
		priceValidity = minListValidity
	}

	// if validity from offer is shorter, use it
	if maxOfferValidity != nil && (priceValidity == nil || (*time.Time)(priceValidity).After(time.Time(*maxOfferValidity))) {
		priceValidity = maxOfferValidity
	}

	// no listing or auction? use last trade price
	if tokenPrice.Usd == 0 {
		tokenPrice = lastTradePrice
	}

	return bson.M{
		fiTokenMinListPrice: minListPrice,
		fiTokenMaxOfferPrice: maxOfferPrice,
		fiTokenAmountLastBid: bidPrice,
		fiTokenReservePrice: reservePrice,
		fiTokenAmountLastTrade: lastTradePrice,
		fiTokenPrice: tokenPrice.Usd,
		fiTokenPriceValid: priceValidity,
		fiTokenPriceUpdate: time.Now(),
	}, nil
}

// TokenLikesViewsStore updates tokens views/likes from shared database.
func (db *MongoDbBridge) TokenLikesViewsStore(t *types.Token) error {
	return db.UpdateToken(&t.Contract, t.TokenId.ToInt(), bson.M{
		fiTokenCachedLikes: t.CachedLikes,
		fiTokenCachedViews: t.CachedViews,
		fiTokenLikesUpdate: t.LikesUpdate,
	})
}

// TokenLikesViewsRefreshSet pulls s set of tokens scheduled to be their views/likes updated.
func (db *MongoDbBridge) TokenLikesViewsRefreshSet(setSize int64) ([]*types.Token, error) {
	list := make([]*types.Token, setSize)
	col := db.client.Database(db.dbName).Collection(coTokens)

	// load the set from database
	cur, err := col.Find(
		context.Background(),
		bson.D{{Key: fiTokenIsActive, Value: true}},
		options.Find().SetSort(bson.D{{Key: fiTokenLikesUpdate, Value: 1}}).SetLimit(setSize),
	)
	if err != nil {
		log.Errorf("can not pull likes/views refresh set; %s", err.Error())
		return nil, err
	}
	defer closeFindCursor("TokenLikesViewsRefreshSet", cur)

	var i int
	for cur.Next(context.Background()) {
		var row types.Token
		if err := cur.Decode(&row); err != nil {
			log.Errorf("can not decode Token; %s", err.Error())
			return nil, err
		}
		list[i] = &row
		i++
	}
	return list[:i], nil
}

func (db *MongoDbBridge) TokensCount(filter *types.TokenFilter) (count int64, err error) {
	col := db.client.Database(db.dbName).Collection(coTokens)
	bsonFilter := tokenFilterToBson(filter)
	count, err = db.getTotalCount(col, bsonFilter)
	if err != nil {
		log.Errorf("tokens count failed, filter: %+v; %s", bsonFilter, err)
		return 0, fmt.Errorf("tokens count failed; %s", err)
	}
	return count, err
}

func (db *MongoDbBridge) ListTokens(
	filter *types.TokenFilter,
	sorting sorting.TokenSorting,
	sortDesc bool,
	cursor types.Cursor,
	count int,
	backward bool,
) (out *types.TokenList, err error) {
	var list types.TokenList
	col := db.client.Database(db.dbName).Collection(coTokens)
	ctx := context.Background()

	bsonFilter := tokenFilterToBson(filter)
	log.Debugf("Filter: %+v", bsonFilter)

	if count == 0 {
		return &list, nil // interested in TotalCount only
	}

	ld, err := db.findPaginated(col, bsonFilter, cursor, count, sorting, backward != sortDesc)
	if err != nil {
		log.Errorf("listing tokens failed, filter: %+v; %s", bsonFilter, err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer closeFindCursor("ListTokens", ld)

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.Token
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the token in list [%s]; %s", ld.Current.String(), err.Error())
				return nil, err
			}
			list.Collection = append(list.Collection, &row)
		} else {
			// skip the last row and set HasNext only
			list.HasNext = true
		}
	}

	if cursor != "" {
		list.HasPrev = true
	}
	if backward {
		list.Reverse()
	}
	return &list, nil
}

// tokenFilterToBson creates a BSON representation of the provided token filter.
func tokenFilterToBson(f *types.TokenFilter) bson.D {
	// return default filter BSON if nothing set
	if f == nil {
		return bson.D{
			{Key: fiTokenIsActive, Value: true},
		}
	}

	// make new filter
	filter := bson.D{}

	// exclude inactive tokens (if not requested to include them)
	if f.IncludeInactive == nil || *f.IncludeInactive != true {
		filter = append(filter,
			bson.E{Key: fiTokenIsActive, Value: true},
			bson.E{Key: fiTokenIsBanned, Value: bson.D{{Key: "$ne", Value: true}}},
			bson.E{Key: fiTokenIsColBanned, Value: bson.D{{Key: "$ne", Value: true}}},
		)
	}

	// regular expression search
	if f.Search != nil {
		filter = append(filter, bson.E{Key: fiTokenName, Value: primitive.Regex{
			Pattern: *f.Search,
			Options: "i",
		}})
	}

	now := types.Time(time.Now().UTC())
	if f.HasListing != nil {
		if *f.HasListing {
			filter = filterAddDateTimeLimit(filter, fiTokenHasListingSince, "$lte", now)
		} else {
			filter = filterAddDateTimeMiss(filter, fiTokenHasListingSince, "$gt", now)
		}
	}

	if f.HasAuction != nil {
		if *f.HasAuction {
			filter = filterAddDateTimeLimit(filter, fiTokenHasAuctionSince, "$lte", now)
			filter = filterAddDateTimeLimit(filter, fiTokenHasAuctionUntil, "$gt", now)
		} else {
			filter = filterAddDateTimeMiss(filter, fiTokenHasAuctionSince, "$gt", now)
			filter = filterAddDateTimeMiss(filter, fiTokenHasAuctionUntil, "$lte", now)
		}
	}

	if f.HasOffer != nil {
		if *f.HasOffer {
			filter = filterAddDateTimeLimit(filter, fiTokenHasOfferUntil, "$gt", now)
		} else {
			filter = filterAddDateTimeMiss(filter, fiTokenHasOfferUntil, "$lte", now)
		}
	}

	if f.HasBids != nil {
		filter = append(filter, bson.E{Key: fiTokenHasBids, Value: *f.HasBids})
		// filter for HasAuction time-range too, to exclude time-terminated auctions
		filter = filterAddDateTimeLimit(filter, fiTokenHasAuctionSince, "$lte", now)
		filter = filterAddDateTimeLimit(filter, fiTokenHasAuctionUntil, "$gt", now)
	}

	if f.Collections != nil && len(*f.Collections) > 0 {
		if len(*f.Collections) == 1 {
			filter = append(filter, bson.E{Key: fiTokenContract, Value: (*f.Collections)[0].String()})
		} else {
			values := make([]string, len(*f.Collections))
			for i, value := range *f.Collections {
				values[i] = value.String()
			}
			filter = append(filter, bson.E{Key: fiTokenContract, Value: bson.D{{Key: "$in", Value: values}}})
		}
	}

	if f.Categories != nil && len(*f.Categories) > 0 {
		if len(*f.Categories) == 1 {
			filter = append(filter, bson.E{Key: fiTokenCategories, Value: (*f.Categories)[0]})
		} else {
			filter = append(filter, bson.E{Key: fiTokenCategories, Value: bson.D{{Key: "$in", Value: *f.Categories}}})
		}
	}

	if f.CreatedBy != nil {
		filter = append(filter, bson.E{Key: fiTokenCreatedBy, Value: *f.CreatedBy})
	}

	if f.PriceMin != nil {
		filter = append(filter, bson.E{Key: fiTokenPrice, Value: bson.D{{Key: "$gte", Value: f.PriceMin.ToInt().Int64()}}})
	}
	if f.PriceMax != nil {
		filter = append(filter, bson.E{Key: fiTokenPrice, Value: bson.D{{Key: "$lte", Value: f.PriceMax.ToInt().Int64()}}})
	}

	if f.ListPriceMin != nil {
		filter = append(filter, bson.E{Key: fiTokenListUsdPrice, Value: bson.D{{Key: "$gte", Value: f.ListPriceMin.ToInt().Int64()}}})
	}
	if f.ListPriceMax != nil {
		filter = append(filter, bson.E{Key: fiTokenListUsdPrice, Value: bson.D{{Key: "$lte", Value: f.ListPriceMax.ToInt().Int64()}}})
	}

	if f.OfferPriceMin != nil {
		filter = append(filter, bson.E{Key: fiTokenOfferUsdPrice, Value: bson.D{{Key: "$gte", Value: f.OfferPriceMin.ToInt().Int64()}}})
	}
	if f.OfferPriceMax != nil {
		filter = append(filter, bson.E{Key: fiTokenOfferUsdPrice, Value: bson.D{{Key: "$lte", Value: f.OfferPriceMax.ToInt().Int64()}}})
	}

	return filter
}

// ExtendLegacyToken tries to load token metadata details from the shared legacy database.
func (sdb *SharedMongoDbBridge) ExtendLegacyToken(token *types.Token) (*types.Token, error) {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyTokens)

	result := col.FindOne(context.Background(), bson.D{
		{Key: fiLegacyTokenContract, Value: strings.ToLower(token.Contract.String())},
		{Key: fiLegacyTokenTokenId, Value: int32(token.TokenId.ToInt().Int64())},
	})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			log.Debugf("token %s / %s not found in legacy db", token.Contract.String(), token.TokenId.String())
			return token, nil
		}
		return nil, result.Err()
	}

	var row types.LegacyToken
	if err := result.Decode(&row); err != nil {
		log.Errorf("can not decode LegacyToken %s / %s; %s", token.Contract.String(), token.TokenId.String(), err.Error())
		return token, nil
	}

	token.Name = row.Name
	token.ImageURI = row.ImageURL
	token.IsActive = row.IsActive
	return token, nil
}
