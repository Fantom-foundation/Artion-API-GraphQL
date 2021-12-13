package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type BannedNft types.BannedNft

type BannedNftEdge struct {
	Node *BannedNft
}

func (edge BannedNftEdge) Cursor() (types.Cursor, error) {
	return sorting.BannedNftSortingUpdated.GetCursor((*types.BannedNft)(edge.Node))
}

type BannedNftConnection struct {
	Edges      []BannedNftEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

func NewBannedNftConnection(list *types.BannedNftList) (con *BannedNftConnection, err error) {
	con = new(BannedNftConnection)
	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]BannedNftEdge, len(list.Collection))
	for i := 0; i < len(list.Collection); i++ {
		con.Edges[i].Node = (*BannedNft)(list.Collection[i])
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

func (o BannedNft) Token() (*Token, error) {
	return NewToken(&o.Contract, &o.TokenId)
}

func (rs *RootResolver) BannedTokens(ctx context.Context, args struct {
	PaginationInput
}) (con *BannedNftConnection, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return nil, err
	}
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListBannedNfts(cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewBannedNftConnection(list)
}
