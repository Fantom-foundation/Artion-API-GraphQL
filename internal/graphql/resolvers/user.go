package resolvers

import (
	"artion-api-graphql/internal/auth"
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// User represents a user profile in Artion database.
type User types.User

// UserEdge is an item in the User list.
type UserEdge struct {
	Node *User
}

// UserConnection represents resolvable paginated list of users.
type UserConnection struct {
	Edges      []UserEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

// Cursor generates unique row identifier of the scrollable Tokens list.
func (edge UserEdge) Cursor() (types.Cursor, error) {
	// dbUser is always already loaded when in Edge
	return sorting.UserSortingNone.GetCursor((*types.User)(edge.Node))
}

// NewUserConnection creates new resolver of scrollable User list connector.
func NewUserConnection(list *types.UserList) (con *UserConnection, err error) {
	con = new(UserConnection)

	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]UserEdge, len(list.Collection))

	for i := 0; i < len(list.Collection); i++ {
		con.Edges[i] = UserEdge{
			Node: (*User)(list.Collection[i]),
		}
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

// Email of the user is resolved only for the user himself, or to logged-in party.
func (user *User) Email(ctx context.Context) (*string, error) {
	logged, _ := auth.GetIdentityOrNil(ctx)
	if logged == nil || 0 != bytes.Compare(logged.Bytes(), user.Address.Bytes()) {
		return nil, nil
	}
	return user.EmailAddress, nil
}

// AvatarThumb resolves a thumbnail to user's avatar image, if available.
func (user *User) AvatarThumb() (*string, error) {
	if user.Avatar == nil || *user.Avatar == "" {
		return nil, nil
	}
	url := fmt.Sprintf("/images/avatar/%s/%s", user.Address.String(), *user.Avatar)
	return &url, nil
}

// Ownerships resolves a paginated list of tokens owned by the given user.
func (user *User) Ownerships(args struct {
	Collection *common.Address
	PaginationInput
}) (con *OwnershipConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	list, err := repository.R().ListOwnerships(args.Collection, nil, &user.Address, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewOwnershipConnection(list)
}

// TokenLikes resolves a list of likes of the user
func (user *User) TokenLikes(args struct{ PaginationInput }) (con *TokenLikeConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	list, err := repository.R().ListUserTokenLikes(&user.Address, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewTokenLikeConnection(list)
}

// CreatedTokens resolved a list of tokens created by the user.
func (user *User) CreatedTokens(args struct {
	PaginationInput
}) (con *TokenConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	filter := types.TokenFilter{
		CreatedBy:       &user.Address,
		IncludeInactive: true,
	}
	list, err := repository.R().ListTokens(&filter, sorting.TokenSortingCreated, true, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewTokenConnection(list, sorting.TokenSortingCreated)
}

// Followers resolves a list of users following this user.
func (user *User) Followers(args struct{ PaginationInput }) (con *FollowConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	list, err := repository.R().ListUserFollowers(user.Address, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewFollowConnection(list)
}

// Following resolves a list of users this user follows.
func (user *User) Following(args struct{ PaginationInput }) (con *FollowConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	list, err := repository.R().ListUserFollowed(user.Address, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewFollowConnection(list)
}

// Activities resolves a list of activity events of the user.
func (user *User) Activities(args struct {
	Filter *ActivityFilter
	PaginationInput
}) (con *ActivityConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	var actTypes []types.ActivityType
	if args.Filter != nil && args.Filter.Types != nil {
		for _, strType := range *args.Filter.Types {
			actTypes = append(actTypes, ActivityTypeFromString(strType))
		}
	}

	list, err := repository.R().ListActivities(nil, nil, &user.Address, actTypes, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewActivityConnection(list)
}

// Offers resolves a list of offers given to the user.
func (user *User) Offers(args struct{ PaginationInput }) (con *OfferConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListOffers(nil, nil, nil, &user.Address, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewOfferConnection(list)
}

// MyOffers resolves a list of offers this user gave to others.
func (user User) MyOffers(args struct{ PaginationInput }) (con *OfferConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	list, err := repository.R().ListOffers(nil, nil, &user.Address, nil, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewOfferConnection(list)
}

// User resolves a user profile for the given address
func (rs *RootResolver) User(args struct{ Address common.Address }) (*User, error) {
	return getUserByAddress(args.Address)
}

// LoggedUser resolves a profile of the current logged-in user.
func (rs *RootResolver) LoggedUser(ctx context.Context) (*User, error) {
	address, err := auth.GetIdentityOrNil(ctx)
	if address == nil || err != nil {
		return nil, err
	}

	return getUserByAddress(*address)
}

// UpdateUser resolves a mutation for user profile modification.
func (rs *RootResolver) UpdateUser(ctx context.Context, args struct {
	Username *string
	Bio      *string
	Email    string
}) (bool, error) {
	address, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return false, err
	}

	user := types.User{
		Address:      *address,
		Username:     args.Username,
		Bio:          args.Bio,
		EmailAddress: &args.Email,
	}
	err = repository.R().StoreUser(&user)
	return err == nil, err
}

// InitiateLogin provides a random challenge to be used to authenticate a user.
func (rs *RootResolver) InitiateLogin() (string, error) {
	return auth.GetAuthenticator().GenerateChallenge()
}

// Login authenticates a user using given challenge and signature.
func (rs *RootResolver) Login(args struct {
	User      common.Address
	Challenge string
	Signature string
}) (string, error) {
	return auth.GetAuthenticator().GenerateBearer(args.Challenge, args.User, args.Signature)
}

// Users resolves a list of user profiles.
func (rs *RootResolver) Users(args struct {
	Search *string
	PaginationInput
}) (con *UserConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	search := ""
	if args.Search != nil {
		search = *args.Search
	}

	list, err := repository.R().ListUserUsers(search, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewUserConnection(list)
}

// getUserByAddress loads a User by his address.
func getUserByAddress(address common.Address) (user *User, err error) {
	dbUser, err := repository.R().GetUser(address)
	if err != nil {
		return nil, err
	}
	return (*User)(dbUser), nil
}
