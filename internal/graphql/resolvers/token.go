package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Token object is constructed from query, data from db are loaded on demand into "loaded" field.
type Token struct {
	Address common.Address
	TokenId hexutil.Big
	loaded *types.Token
}

func (t *Token) load() error {
	if t.loaded == nil { // TODO: need to add synchronization to prevent concurrent loads!
		tok, err := repository.R().GetToken(t.Address, t.TokenId)
		if err != nil {
			return err
		}
		t.loaded = tok
	}
	return nil
}

func (t *Token) Name() (string, error) {
	err := t.load(); if err != nil {
		return "", err
	}
	return t.loaded.Name, nil
}

func (t *Token) Description() (string, error) {
	err := t.load(); if err != nil {
		return "", err
	}
	return t.loaded.Description, nil
}

func (t *Token) Events(args struct{ PaginationInput }) (con *TokenEventConnection, err error) {
	cursor, count, backward, err := args.ToCursorCount()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListTokenEvents(&t.Address, &t.TokenId, nil, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewTokenEventConnection(list)
}
