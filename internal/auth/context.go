package auth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// identityContextKeyName is the name of the identity field
// in the context values.
const identityContextKeyName = "identity"

// SetIdentity adds the given identity to the provided context
// returning a new derived context.
func SetIdentity(ctx context.Context, i *common.Address) context.Context {
	return context.WithValue(ctx, identityContextKeyName, *i)
}

// GetIdentityOrNil extracts the active identity from the provided context
// and returns nil if the identity can not be extracted.
func GetIdentityOrNil(ctx context.Context) (*common.Address, error) {
	value := ctx.Value(identityContextKeyName)
	if value == nil {
		return nil, nil
	}
	address, ok := value.(common.Address)
	if !ok {
		return nil, fmt.Errorf("identity type mismatch")
	}
	return &address, nil
}

// GetIdentityOrErr extracts the active identity from the provided context
// and returns an error if the identity can not be extracted.
func GetIdentityOrErr(ctx context.Context) (*common.Address, error) {
	value := ctx.Value(identityContextKeyName)
	if value == nil {
		return nil, fmt.Errorf("not authenticated")
	}
	address, ok := value.(common.Address)
	if !ok {
		return nil, fmt.Errorf("identity type mismatch")
	}
	return &address, nil
}
