package resolvers

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type Auction types.Auction

type AuctionEdge struct {
	Node *Auction
}

type AuctionConnection struct {
	Edges      []AuctionEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

func (o Auction) Token() (*Token, error) {
	return NewToken(&o.Contract, &o.TokenId)
}

func (edge AuctionEdge) Cursor() (types.Cursor, error) {
	return sorting.AuctionSortingNone.GetCursor((*types.Auction)(edge.Node))
}

func NewAuctionConnection(list *types.AuctionList) (con *AuctionConnection, err error) {
	con = new(AuctionConnection)
	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]AuctionEdge, len(list.Collection))
	for i := 0; i < len(list.Collection); i++ {
		con.Edges[i].Node = (*Auction)(list.Collection[i])
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
