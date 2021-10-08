package resolvers

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type Listing types.Listing

func (l Listing) Token() (Token, error) {
	return Token{
		Contract: l.Contract,
		TokenId:  l.TokenId,
	}, nil
}

type ListingEdge struct {
	Node *Listing
}

func (edge ListingEdge) Cursor() (types.Cursor, error) {
	id := (*types.Listing)(edge.Node).ID()
	return types.CursorFromId(id[:]), nil
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
