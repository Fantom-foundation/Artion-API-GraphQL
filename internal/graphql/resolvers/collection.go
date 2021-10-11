package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// Collection object is constructed from query, data from db are loaded on demand into "dbCollection" field.
type Collection struct {
	Contract     common.Address
	dbCollection *types.Collection // data for collection loaded from Mongo
}

type CollectionEdge struct {
	Node *Collection
}

func (edge CollectionEdge) Cursor() (types.Cursor, error) {
	return sorting.CollectionSortingNone.GetCursor(edge.Node.dbCollection)
}

type CollectionConnection struct {
	Edges      []CollectionEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

func NewCollectionConnection(list *types.CollectionList) (con *CollectionConnection, err error) {
	con = new(CollectionConnection)
	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]CollectionEdge, len(list.Collection))
	for i := 0; i < len(list.Collection); i++ {
		resolverCollection := Collection{
			Contract:     list.Collection[i].Address,
			dbCollection: list.Collection[i],
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

func (t *Collection) load() error {
	if t.dbCollection == nil {
		tok, err := repository.R().GetCollection(t.Contract)
		if err != nil {
			return err
		}
		t.dbCollection = tok
	}
	return nil
}

func (t Collection) Type() (string, error) {
	err := t.load()
	if err != nil {
		return "", err
	}
	return t.dbCollection.Type, nil
}

func (t Collection) Name() (string, error) {
	err := t.load()
	if err != nil {
		return "", err
	}
	return t.dbCollection.Name, nil
}

func (t Collection) Symbol() (string, error) {
	err := t.load()
	if err != nil {
		return "", err
	}
	return t.dbCollection.Symbol, nil
}

func (t Collection) Created() (types.Time, error) {
	err := t.load()
	if err != nil {
		return types.Time{}, err
	}
	return types.Time(t.dbCollection.Created), nil
}

func (t Collection) IsActive() (bool, error) {
	err := t.load()
	if err != nil {
		return false, err
	}
	return t.dbCollection.IsActive, nil
}
