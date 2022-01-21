package resolvers

import (
	"artion-api-graphql/internal/auth"
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
	"strings"
	"time"
)

// Token object is constructed from query, data from db are loaded on demand into "dbToken" field.
type Token types.Token

// TokenEdge represents an edge in scrollable Tokens list.
type TokenEdge struct {
	Node    *Token
	sorting sorting.TokenSorting
}

// TokenConnection represents scrollable tokens list connector.
type TokenConnection struct {
	Edges      []TokenEdge
	PageInfo   PageInfo
	filter     *types.TokenFilter
}

// NewToken creates a new instance of the resolvable Token.
func NewToken(contract *common.Address, tokenID *hexutil.Big) (*Token, error) {
	tok, err := repository.R().Token(contract, tokenID)
	if err != nil {
		return nil, err
	}
	return (*Token)(tok), nil
}

// NewTokenConnection creates new resolver of scrollable token list connector.
func NewTokenConnection(list *types.TokenList, sorting sorting.TokenSorting, filter *types.TokenFilter) (con *TokenConnection, err error) {
	con = new(TokenConnection)

	con.Edges = make([]TokenEdge, len(list.Collection))
	con.filter = filter

	for i := 0; i < len(list.Collection); i++ {
		tok, err := NewToken(&list.Collection[i].Contract, &list.Collection[i].TokenId)
		if err != nil {
			return nil, err
		}

		con.Edges[i] = TokenEdge{
			Node:    tok,
			sorting: sorting,
		}
	}

	con.PageInfo.HasNextPage = list.HasNext
	con.PageInfo.HasPreviousPage = list.HasPrev

	if len(list.Collection) > 0 {
		startCur, err := con.Edges[0].Cursor()
		if err != nil {
			return nil, err
		}

		endCur, err := con.Edges[len(con.Edges)-1].Cursor()
		if err != nil {
			return nil, err
		}

		con.PageInfo.StartCursor = &startCur
		con.PageInfo.EndCursor = &endCur
	}
	return con, err
}

func (t *TokenConnection) TotalCount() (hexutil.Big, error) {
	count, err := repository.R().TokensCount(t.filter)
	return (hexutil.Big)(*big.NewInt(count)), err
}

// Image resolves URI of the token image.
func (t *Token) Image() *string {
	if t.ImageURI == "" {
		return nil
	}
	if strings.HasPrefix(t.ImageURI, "ipfs://") {
		uri := "https://artion.mypinata.cloud/ipfs/" + t.ImageURI[7:]
		return &uri
	}
	return &t.ImageURI
}

// ImageMimetype resolves mimetype of the token image.
func (t *Token) ImageMimetype() *string {
	if t.ImageURI == "" || t.ImageType == types.ImageTypeUnknown {
		return nil
	}
	mimetype := t.ImageType.Mimetype()
	return &mimetype
}

// HasListing checks if the given token has any active listing right now.
func (t *Token) HasListing() bool {
	if nil == t.HasListingSince {
		return false
	}
	return (*time.Time)(t.HasListingSince).Before(time.Now().UTC())
}

// HasOffer checks if the given token has any active offers right now.
func (t *Token) HasOffer() bool {
	if nil == t.HasOfferUntil {
		return false
	}
	return (*time.Time)(t.HasOfferUntil).After(time.Now().UTC())
}

// HasAuction checks if the given token has any active auction right now.
func (t *Token) HasAuction() bool {
	if nil == t.HasAuctionSince {
		return false
	}
	now := time.Now().UTC()
	return (*time.Time)(t.HasAuctionSince).Before(now) && (*time.Time)(t.HasAuctionUntil).After(now)
}

// ImageThumb generates REST path providing the token image thumbnail from this Artion API.
func (t *Token) ImageThumb() *string {
	if t.ImageURI == "" {
		return nil
	}
	uri := fmt.Sprintf("/images/token/%s/%s", t.Contract.String(), t.TokenId.String())
	return &uri
}

func (t *Token) Likes() (hexutil.Big, error) {
	count, err := repository.R().GetTokenLikesCount(&t.Contract, (*big.Int)(&t.TokenId))
	if err != nil {
		return hexutil.Big{}, err
	}
	return hexutil.Big(*big.NewInt(count)), nil
}

func (t *Token) IsLiked(ctx context.Context) (bool, error) {
	identity, err := auth.GetIdentityOrNil(ctx)
	if identity == nil || err != nil {
		return false, err
	}
	return repository.R().IsTokenLiked(identity, &t.Contract, (*big.Int)(&t.TokenId))
}

func (t *Token) IsLikedBy(args struct{ User *common.Address }) (bool, error) {
	if args.User == nil {
		return false, nil
	}
	return repository.R().IsTokenLiked(args.User, &t.Contract, (*big.Int)(&t.TokenId))
}

func (t *Token) Views() (hexutil.Big, error) {
	count, err := repository.R().GetTokenViews(t.Contract, big.Int(t.TokenId))
	if err != nil {
		return hexutil.Big{}, err
	}
	return hexutil.Big(*count), nil
}

// Collection resolves collection of the token.
func (t *Token) Collection() (*Collection, error) {
	return NewCollection(&t.Contract)
}

func (t *Token) FeeRecipientUser() (*User, error) {
	return getUserByAddressPtr(t.FeeRecipient)
}

func (t *Token) Ownerships(args struct{ PaginationInput }) (con *OwnershipConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListOwnerships(&t.Contract, &t.TokenId, nil, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewOwnershipConnection(list)
}

func (t *Token) Listings(args struct{ PaginationInput }) (con *ListingConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListListings(&t.Contract, &t.TokenId, nil, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewListingConnection(list)
}

func (t *Token) Offers(args struct{ PaginationInput }) (con *OfferConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListOffers(&t.Contract, &t.TokenId, nil, nil, false, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewOfferConnection(list)
}

func (t Token) Activities(args struct {
	Filter *ActivityFilter
	PaginationInput
}) (con *ActivityConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	var actTypes []types.ActivityType
	if args.Filter != nil && args.Filter.Types != nil {
		for _, strType := range *args.Filter.Types {
			actTypes = append(actTypes, ActivityTypeFromString(strType))
		}
	}
	list, err := repository.R().ListActivities(&t.Contract, &t.TokenId, nil, actTypes, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewActivityConnection(list)
}

func (t *Token) Auction() (auction *Auction, err error) {
	a, err := repository.R().GetLastAuction(&t.Contract, (*big.Int)(&t.TokenId))
	if err != nil {
		return nil, err
	}
	return (*Auction)(a), nil
}

func (t *Token) ListingPrice() (*types.TokenPrice, error) {
	if t.MinListPrice.Amount.ToInt().Uint64() == 0 {
		return nil, nil
	}
	return &t.MinListPrice, nil
}

func (t *Token) AuctionedPrice() (*types.TokenPrice, error) {
	if !t.HasAuction() || !t.HasBids {
		return nil, nil
	}
	return &t.AmountLastBid, nil
}

func (t *Token) AuctionReservePrice() (*types.TokenPrice, error) {
	if !t.HasAuction() {
		return nil, nil
	}
	return &t.ReservePrice, nil
}

func (t *Token) OfferedPrice() (*types.TokenPrice, error) {
	if t.MaxOfferPrice.Amount.ToInt().Uint64() == 0 {
		return nil, nil
	}
	return &t.MaxOfferPrice, nil
}

func (t *Token) LastTradePrice() (*types.TokenPrice, error) {
	if t.AmountLastTrade.Amount.ToInt().Uint64() == 0 {
		return nil, nil
	}
	return &t.AmountLastTrade, nil
}

func (t *Token) UsdPrice() string {
	return strconv.FormatInt(t.AmountPrice, 10)
}

func (t *Token) PriceHistory(args struct {
	From types.Time
	To   types.Time
}) ([]types.PriceHistory, error) {
	return repository.R().TokenPriceHistory(&t.Contract, &t.TokenId, time.Time(args.From), time.Time(args.To))
}

// Cursor generates unique row identifier of the scrollable Tokens list.
func (edge TokenEdge) Cursor() (types.Cursor, error) {
	return edge.sorting.GetCursor((*types.Token)(edge.Node))
}

func (rs *RootResolver) Token(args struct {
	Contract common.Address
	TokenId  hexutil.Big
}) (*Token, error) {
	token, err := NewToken(&args.Contract, &args.TokenId)
	if token != nil && (token.IsBanned || token.IsColBanned) {
		return nil, nil // hide banned token
	}
	return token, err
}

func (rs *RootResolver) Tokens(args struct {
	Filter  *types.TokenFilter
	SortBy  *string
	SortDir *string
	PaginationInput
}) (con *TokenConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	srt, err := tokenSortingFromString(args.SortBy)
	if err != nil {
		return nil, err
	}
	sortDesc := isSortingDirectionDesc(args.SortDir)
	if args.Filter == nil {
		args.Filter = &types.TokenFilter{}
	}

	// default sorting Recently Created
	if srt == sorting.TokenSortingNone {
		srt = sorting.TokenSortingCreated
		sortDesc = true
	}

	// when sorting by price, skip zero price (not-listed/auctioned) items
	if (srt == sorting.TokenSortingPrice || args.Filter.PriceMax != nil) && args.Filter.PriceMin == nil {
		args.Filter.PriceMin = (*hexutil.Big)(big.NewInt(1))
	}

	// when filtering max 15.00 USD, include anything what is displayed as 15.00 (like 15.004)
	if args.Filter.PriceMax != nil {
		max := args.Filter.PriceMax.ToInt()
		max.Add(max, big.NewInt(10000)) // add 0.01 USD
	}

	hasListing := args.Filter.HasListing != nil && *args.Filter.HasListing
	hasAuction := args.Filter.HasAuction != nil && *args.Filter.HasAuction
	hasBids := args.Filter.HasBids != nil && *args.Filter.HasBids
	hasOffer := args.Filter.HasOffer != nil && *args.Filter.HasOffer

	// when filtering listed only, replace general-price by listing-price sorting/filtering
	if hasListing && !hasAuction && !hasBids {
		if srt == sorting.TokenSortingPrice {
			srt = sorting.TokenSortingListPrice
		}
		args.Filter.ListPriceMin = args.Filter.PriceMin
		args.Filter.PriceMin = nil
		args.Filter.ListPriceMax = args.Filter.PriceMax
		args.Filter.PriceMax = nil
	}

	// when filtering offers only, replace general-price by offer-price sorting/filtering
	if hasOffer && !hasListing && !hasAuction && !hasBids {
		if srt == sorting.TokenSortingPrice {
			srt = sorting.TokenSortingOfferPrice
		}
		args.Filter.OfferPriceMin = args.Filter.PriceMin
		args.Filter.PriceMin = nil
		args.Filter.OfferPriceMax = args.Filter.PriceMax
		args.Filter.PriceMax = nil
	}

	// when filtering auctions, general-price is OK!

	list, err := repository.R().ListTokens(args.Filter, srt, sortDesc, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewTokenConnection(list, srt, args.Filter)
}

// IncrementTokenViews increments amount of views of the token.
func (rs *RootResolver) IncrementTokenViews(args struct {
	Contract common.Address
	TokenId  hexutil.Big
}) (bool, error) {
	err := repository.R().IncrementTokenViews(args.Contract, big.Int(args.TokenId))
	return err == nil, err
}
