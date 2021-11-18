package resolvers

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type Listing types.Listing

func (l Listing) Token() (*Token, error) {
	return NewToken(&l.Contract, &l.TokenId)
}

func (l Listing) OwnerUser() (User, error) {
	return getUserByAddress(l.Owner)
}

type ListingEdge struct {
	Node *Listing
}

func (edge ListingEdge) Cursor() (types.Cursor, error) {
	return sorting.ListingSortingNone.GetCursor((*types.Listing)(edge.Node))
}

type ListingConnection struct {
	Edges      []ListingEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

func NewListingConnection(list *types.ListingList) (con *ListingConnection, err error) {
	con = new(ListingConnection)
	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]ListingEdge, len(list.Collection))
	for i := 0; i < len(list.Collection); i++ {
		con.Edges[i].Node = (*Listing)(list.Collection[i])
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
