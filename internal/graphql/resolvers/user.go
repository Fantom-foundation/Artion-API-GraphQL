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

type User struct {
	Address common.Address
	dbUser  *types.User // data for user loaded from Mongo
}

type UserEdge struct {
	Node    *User
}

// Cursor generates unique row identifier of the scrollable Tokens list.
func (edge UserEdge) Cursor() (types.Cursor, error) {
	// dbUser is always already loaded when in Edge
	return sorting.UserSortingNone.GetCursor(edge.Node.dbUser)
}

type UserConnection struct {
	Edges      []UserEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

// NewUserConnection creates new resolver of scrollable User list connector.
func NewUserConnection(list *types.UserList) (con *UserConnection, err error) {
	con = new(UserConnection)

	con.TotalCount = (hexutil.Big)(*big.NewInt(list.TotalCount))
	con.Edges = make([]UserEdge, len(list.Collection))

	for i := 0; i < len(list.Collection); i++ {
		user := User{
			Address: list.Collection[i].Address,
			dbUser: list.Collection[i],
		}
		con.Edges[i] = UserEdge{
			Node:    &user,
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

func (user User) Username() (*string, error) {
	if user.dbUser == nil {
		return nil, nil
	}
	return user.dbUser.Username, nil
}

func (user User) Bio() (*string, error) {
	if user.dbUser == nil {
		return nil, nil
	}
	return user.dbUser.Bio, nil
}

func (user User) Email(ctx context.Context) (*string, error) {
	logged, _ := auth.GetIdentityOrNil(ctx) // email available only for the user itself
	if logged == nil || !bytes.Equal(logged.Bytes(), user.Address.Bytes()) || user.dbUser == nil {
		return nil, nil
	}
	return user.dbUser.Email, nil
}

func (user User) Avatar() (*string, error) {
	if user.dbUser == nil || user.dbUser.Avatar == nil || *user.dbUser.Avatar == "" {
		return nil, nil
	}
	return user.dbUser.Avatar, nil
}

func (user User) AvatarThumb() (*string, error) {
	if user.dbUser == nil || user.dbUser.Avatar == nil || *user.dbUser.Avatar == "" {
		return nil, nil
	}
	url := fmt.Sprintf("/images/avatar/%s/%s", user.Address.String(), *user.dbUser.Avatar)
	return &url, nil
}

func (user User) Banner() (*string, error) {
	if user.dbUser == nil || user.dbUser.Banner == nil || *user.dbUser.Banner == "" {
		return nil, nil
	}
	return user.dbUser.Banner, nil
}

func (user User) Ownerships(args struct{ PaginationInput }) (con *OwnershipConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListOwnerships(nil, nil, &user.Address, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewOwnershipConnection(list)
}

func (user User) TokenLikes(args struct{ PaginationInput }) (con *TokenLikeConnection, err error) {
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

func (user User) CreatedTokens(args struct {
	PaginationInput
}) (con *TokenConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	filter := types.TokenFilter{
		CreatedBy: &user.Address,
	}
	list, err := repository.R().ListTokens(&filter, sorting.TokenSortingCreated, true, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewTokenConnection(list, sorting.TokenSortingCreated)
}

func (user User) Followers(args struct{ PaginationInput }) (con *FollowConnection, err error) {
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

func (user User) Following(args struct{ PaginationInput }) (con *FollowConnection, err error) {
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

func (user User) Activities(args struct{
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

func (user User) MyOffers(args struct{ PaginationInput }) (con *OfferConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}
	list, err := repository.R().ListOffers(nil, nil, &user.Address, cursor, count, backward)
	if err != nil {
		return nil, err
	}
	return NewOfferConnection(list)
}

func getUserByAddress(address common.Address) (user User, err error) {
	dbUser, err := repository.R().GetUser(address)
	if err != nil {
		return User{}, err
	} else {
		return User{Address: address, dbUser: dbUser}, nil
	}
}

func (rs *RootResolver) User(args struct{ Address common.Address }) (user User, err error) {
	return getUserByAddress(args.Address)
}

func (rs *RootResolver) LoggedUser(ctx context.Context) (user *User, err error) {
	address, err := auth.GetIdentityOrNil(ctx)
	if address == nil || err != nil {
		return nil, err
	}
	dbUser, err := repository.R().GetUser(*address)
	if err != nil {
		return nil, err
	} else {
		return &User{Address: *address, dbUser: dbUser}, nil
	}
}

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
		Address:  *address,
		Username: args.Username,
		Bio:      args.Bio,
		Email:    &args.Email,
	}
	err = repository.R().StoreUser(&user)
	return err == nil, err
}

func (rs *RootResolver) InitiateLogin() (string, error) {
	return auth.GetAuthenticator().GenerateChallenge()
}

func (rs *RootResolver) Login(args struct {
	User      common.Address
	Challenge string
	Signature string
}) (string, error) {
	return auth.GetAuthenticator().GenerateBearer(args.Challenge, args.User, args.Signature)
}

func (rs *RootResolver) Users(args struct {
	Search  *string
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
