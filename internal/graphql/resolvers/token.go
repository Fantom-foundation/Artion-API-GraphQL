package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// Token object is constructed from query, data from db are loaded on demand into "dbToken" field.
type Token struct {
	Contract common.Address
	TokenId  hexutil.Big
	dbToken  *types.Token // data for token loaded from Mongo
}

type TokenEdge struct {
	Node *Token
	sorting sorting.TokenSorting
}

func (edge TokenEdge) Cursor() (types.Cursor, error) {
	// dbToken is always already loaded when in Edge
	return edge.sorting.GetCursor(edge.Node.dbToken)
}

type TokenConnection struct {
	Edges      []TokenEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

func NewTokenConnection(list *types.TokenList, sorting sorting.TokenSorting) (con *TokenConnection, err error) {
	con = new(TokenConnection)
	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]TokenEdge, len(list.Collection))
	for i := 0; i < len(list.Collection); i++ {
		resolverToken := Token{
			Contract: list.Collection[i].Contract,
			TokenId:  list.Collection[i].TokenId,
			dbToken:  list.Collection[i],
		}
		con.Edges[i] = TokenEdge{
			Node: &resolverToken,
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

func (t *Token) load() error {
	if t.dbToken == nil {
		tok, err := repository.R().Token(&t.Contract, &t.TokenId)
		if err != nil {
			return fmt.Errorf("unable to load token from database; %s", err)
		}
		t.dbToken = tok
	}
	return nil
}

func (t Token) Name() (string, error) {
	err := t.load()
	if err != nil {
		return "", err
	}
	return t.dbToken.Name, nil
}

func (t Token) Description() (string, error) {
	err := t.load()
	if err != nil {
		return "", err
	}
	return t.dbToken.Description, nil
}

func (t Token) Image() (*string, error) {
	err := t.load()
	if err != nil || t.dbToken.ImageURI == "" {
		return nil, err
	}
	return &t.dbToken.ImageURI, nil
}

func (t Token) ImageProxy() (*string, error) {
	err := t.load()
	if err != nil || t.dbToken.ImageURI == "" {
		return nil, err
	}
	url := "/token-image/" + t.Contract.String() + "/" + t.TokenId.String()
	return &url, nil
}

func (t Token) Created() (types.Time, error) {
	err := t.load()
	if err != nil {
		return types.Time{}, err
	}
	return t.dbToken.Created, nil
}

func (t Token) HasListing() (bool, error) {
	err := t.load()
	if err != nil {
		return false, err
	}
	return t.dbToken.HasListing, nil
}

func (t Token) HasOffer() (bool, error) {
	err := t.load()
	if err != nil {
		return false, err
	}
	return t.dbToken.HasOffer, nil
}

func (t Token) HasAuction() (bool, error) {
	err := t.load()
	if err != nil {
		return false, err
	}
	return t.dbToken.HasAuction, nil
}

func (t Token) HasBids() (bool, error) {
	err := t.load()
	if err != nil {
		return false, err
	}
	return t.dbToken.HasBid, nil
}

func (t Token) LastListing() (*types.Time, error) {
	err := t.load()
	if err != nil || time.Time(t.dbToken.LastList).IsZero() {
		return nil, err
	}
	return &t.dbToken.LastList, nil
}

func (t Token) LastTrade() (*types.Time, error) {
	err := t.load()
	if err != nil || time.Time(t.dbToken.LastTrade).IsZero() {
		return nil, err
	}
	return &t.dbToken.LastTrade, nil
}

func (t Token) LastOffer() (*types.Time, error) {
	err := t.load()
	if err != nil || time.Time(t.dbToken.LastOffer).IsZero() {
		return nil, err
	}
	return &t.dbToken.LastOffer, nil
}

func (t Token) LastBid() (*types.Time, error) {
	err := t.load()
	if err != nil || time.Time(t.dbToken.LastBid).IsZero() {
		return nil, err
	}
	return &t.dbToken.LastBid, nil
}

func (t Token) Ownerships(args struct{ PaginationInput }) (con *OwnershipConnection, err error) {
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

func (t Token) Listings(args struct{ PaginationInput }) (con *ListingConnection, err error) {
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

func (t Token) Offers(args struct{ PaginationInput }) (con *OfferConnection, err error) {
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
