package resolvers

import (
	"artion-api-graphql/internal/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Token struct {
	Address common.Address
	TokenId hexutil.Big
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
