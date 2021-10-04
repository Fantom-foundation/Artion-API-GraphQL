package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// NftCollection object is constructed from query, data from db are loaded on demand into "dbNftCollection" field.
type NftCollection struct {
	Address         common.Address
	dbNftCollection *types.NFTCollection // data for collection loaded from Mongo
}

type NftCollectionEdge struct {
	Node *NftCollection
}

func (edge NftCollectionEdge) Cursor() (types.Cursor, error) {
	return types.CursorFromId(edge.Node.Address.Bytes()), nil
}

type NftCollectionConnection struct {
	Edges      []NftCollectionEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

func NewCollectionConnection(list *types.NFTCollectionList) (con *NftCollectionConnection, err error) {
	con = new(NftCollectionConnection)
	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]NftCollectionEdge, len(list.Collection))
	for i := 0; i < len(list.Collection); i++ {
		resolverCollection := NftCollection{
			Address:         list.Collection[i].Address,
			dbNftCollection: list.Collection[i],
		}
		con.Edges[i].Node = &resolverCollection
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


func (t *NftCollection) load() error {
	if t.dbNftCollection == nil {
		tok, err := repository.R().GetNFTCollection(t.Address)
		if err != nil {
			return err
		}
		t.dbNftCollection = tok
	}
	return nil
}

func (t NftCollection) Type() (string, error) {
	err := t.load()
	if err != nil {
		return "", err
	}
	return t.dbNftCollection.Type, nil
}

func (t NftCollection) Name() (string, error) {
	err := t.load()
	if err != nil {
		return "", err
	}
	return t.dbNftCollection.Name, nil
}

func (t NftCollection) Symbol() (string, error) {
	err := t.load()
	if err != nil {
		return "", err
	}
	return t.dbNftCollection.Symbol, nil
}

func (t NftCollection) Created() (types.Time, error) {
	err := t.load()
	if err != nil {
		return types.Time{}, err
	}
	return types.Time(t.dbNftCollection.Created), nil
}

func (t NftCollection) IsActive() (bool, error) {
	err := t.load()
	if err != nil {
		return false, err
	}
	return t.dbNftCollection.IsActive, nil
}

func (t *NftCollection) Tokens(args struct{ PaginationInput }) (con *TokenConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListCollectionTokens(t.Address, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewTokenConnection(list)
}
