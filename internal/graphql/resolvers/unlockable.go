package resolvers

import (
	"artion-api-graphql/internal/auth"
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (t *Token) HasUnlockable() (bool, error) {
	unlockable, err := repository.R().GetUnlockable(t.Contract, *t.TokenId.ToInt())
	if unlockable == nil {
		return false, err
	}
	return true, err
}

func (rs *RootResolver) UnlockableContent(ctx context.Context, args struct{
	Contract common.Address
	TokenId  hexutil.Big
}) (*string, error) {
	identity, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return nil, err
	}
	isOwner, err := repository.R().IsOwnerOf(args.Contract, args.TokenId, *identity)
	if err != nil {
		return nil, err
	}
	if ! isOwner {
		return nil, fmt.Errorf("not authorized - not owner of the token")
	}
	unlockable, err := repository.R().GetUnlockable(args.Contract, *args.TokenId.ToInt())
	if unlockable == nil {
		return nil, err
	}
	return &unlockable.Content, err
}

func (rs *RootResolver) SetUnlockableContent(ctx context.Context, args struct{
	Contract common.Address
	TokenId  hexutil.Big
	Content  string
}) (bool, error) {
	identity, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return false, err
	}
	isOwner, err := repository.R().IsOwnerOf(args.Contract, args.TokenId, *identity)
	if err != nil {
		return false, err
	}
	if ! isOwner {
		return false, fmt.Errorf("not authorized - not owner of the token")
	}
	token, err := repository.R().Token(&args.Contract, &args.TokenId)
	if token == nil {
		return false, err
	}
	if ! bytes.Equal(token.CreatedBy.Bytes(), identity.Bytes()) {
		return false, fmt.Errorf("not authorized - not minter of the token (minter is %s)", token.CreatedBy.String())
	}
	unlockable := types.Unlockable{
		Contract: args.Contract,
		TokenId: int32(args.TokenId.ToInt().Int64()),
		Content: args.Content,
	}
	err = repository.R().InsertUnlockable(&unlockable)
	return err == nil, err
}
