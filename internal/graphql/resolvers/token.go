package resolvers

import (
	"artion-api-graphql/internal/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type Token struct {
	Address common.Address
	TokenId hexutil.Big
}

func (t *Token) Events(args struct{ PaginationInput }) (list *TokenEventConnection, err error) {
	cursor, count, err := args.ToCursorCount()
	if err != nil {
		return nil, err
	}
	out, err := repository.R().ListTokenEvents(t.Address, t.TokenId, cursor, count)
	if err != nil {
		return nil, err
	}

	list = new(TokenEventConnection)
	list.TotalCount = (hexutil.Big)(*big.NewInt(out.TotalCount))
	list.Edges = make([]TokenEventEdge, len(out.Collection))
	for i := 0; i < len(out.Collection); i++ {
		list.Edges[i].Node = (*TokenEvent)(out.Collection[i])
	}
	list.PageInfo.HasNextPage = out.HasNext
	list.PageInfo.HasPreviousPage = out.HasPrev
	if len(out.Collection) > 0 {
		startCur, err := list.Edges[0].Cursor()
		if err != nil {
			return nil, err
		}
		endCur, err := list.Edges[len(list.Edges)-1].Cursor()
		if err != nil {
			return nil, err
		}
		list.PageInfo.StartCursor = &startCur
		list.PageInfo.EndCursor = &endCur
	}
	return list, err
}
