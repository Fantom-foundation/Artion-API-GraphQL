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
	TotalCount hexutil.Big
	PageInfo   PageInfo
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
func NewTokenConnection(list *types.TokenList, sorting sorting.TokenSorting) (con *TokenConnection, err error) {
	con = new(TokenConnection)

	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]TokenEdge, len(list.Collection))

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

// Image resolves URI of the token image.
func (t *Token) Image() *string {
	if t.ImageURI == "" {
		return nil
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

func (t *Token) Views() (hexutil.Big, error) {
	count, err := repository.R().GetTokenViews(t.Contract, big.Int(t.TokenId))
	if err != nil {
		return hexutil.Big{}, err
	}
	return hexutil.Big(*count), nil
}

// Collection resolves collection of the token.
func (t *Token) Collection() Collection {
	return Collection{
		Contract: t.Contract,
	}
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
	list, err := repository.R().ListOffers(&t.Contract, &t.TokenId, nil, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewOfferConnection(list)
}

func (t Token) Activities(args struct{
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
	a, err := repository.R().GetAuction(&t.Contract, (*big.Int)(&t.TokenId))
	if err != nil {
		return nil, err
	}
	return (*Auction)(a), nil
}

func (t *Token) ListingPrice() (*hexutil.Uint64, error) {
	if t.MinListAmount == 0 {
		return nil, nil
	}
	out := hexutil.Uint64(t.MinListAmount)
	return &out, nil
}

func (t *Token) Price() (*hexutil.Uint64, error) {
	if t.AmountPrice == 0 {
		return nil, nil
	}
	out := hexutil.Uint64(t.AmountPrice)
	return &out, nil
}

// Cursor generates unique row identifier of the scrollable Tokens list.
func (edge TokenEdge) Cursor() (types.Cursor, error) {
	return edge.sorting.GetCursor((*types.Token)(edge.Node))
}

func (rs *RootResolver) Token(args struct {
	Contract common.Address
	TokenId  hexutil.Big
}) (*Token, error) {
	return NewToken(&args.Contract, &args.TokenId)
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
	sorting, err := tokenSortingFromString(args.SortBy)
	if err != nil {
		return nil, err
	}

	list, err := repository.R().ListTokens(args.Filter, sorting, isSortingDirectionDesc(args.SortDir), cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewTokenConnection(list, sorting)
}

// IncrementTokenViews increments amount of views of the token.
func (rs *RootResolver) IncrementTokenViews(args struct {
	Contract common.Address
	TokenId  hexutil.Big
}) (bool, error) {
	err := repository.R().IncrementTokenViews(args.Contract, big.Int(args.TokenId))
	return err == nil, err
}
