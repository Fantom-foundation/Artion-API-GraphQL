package resolvers

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type Ownership types.Ownership

func (o Ownership) Token() (Token, error) {
	return Token{
		Contract: o.Contract,
		TokenId:  o.TokenId,
	}, nil
}

type OwnershipEdge struct {
	Node *Ownership
}

func (edge OwnershipEdge) Cursor() (types.Cursor, error) {
	return sorting.OwnershipSortingNone.GetCursor((*types.Ownership)(edge.Node))
}

type OwnershipConnection struct {
	Edges      []OwnershipEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

func NewOwnershipConnection(list *types.OwnershipList) (con *OwnershipConnection, err error) {
	con = new(OwnershipConnection)
	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]OwnershipEdge, len(list.Collection))
	for i := 0; i < len(list.Collection); i++ {
		con.Edges[i].Node = (*Ownership)(list.Collection[i])
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
