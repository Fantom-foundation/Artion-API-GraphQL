package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
)

// Token object is constructed from query, data from db are loaded on demand into "dbToken" field.
type Token struct {
	*types.Token
}

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
	return &Token{Token: tok}, nil
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

// ImageProxy generates REST path providing the token image from this Artion API.
func (t Token) ImageProxy() *string {
	if t.ImageURI == "" {
		return nil
	}

	var sb strings.Builder
	sb.WriteString("/token-image/")
	sb.WriteString(t.Contract.String())
	sb.WriteString("/")
	sb.WriteString(t.TokenId.String())
	uri := sb.String()

	return &uri
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

// Cursor generates unique row identifier of the scrollable Tokens list.
func (edge TokenEdge) Cursor() (types.Cursor, error) {
	// dbToken is always already loaded when in Edge
	return edge.sorting.GetCursor(edge.Node.Token)
}
