package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
)

// Collection represents a resolvable collection of NFT tokens.
type Collection types.LegacyCollection

// CollectionConnection represents a resolvable connection
// between Collection list and its edges.
type CollectionConnection struct {
	Edges      []CollectionEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

// CollectionEdge represents an edge on Collection list.
type CollectionEdge struct {
	Node *Collection
}

// Collection resolves an NFT collection for the given contract address.
func (rs *RootResolver) Collection(args struct {
	Contract common.Address
}) (*Collection, error) {
	return NewCollection(&args.Contract)
}

// Collections resolve a list of NFT Collection for the given criteria.
func (rs *RootResolver) Collections(args struct {
	Search *string
	MintableBy *common.Address
	PaginationInput
}) (con *CollectionConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	list, err := repository.R().ListLegacyCollections(types.CollectionFilter{
		Search:     args.Search,
		MintableBy: args.MintableBy,
	}, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewCollectionConnection(list)
}

// NewCollection loads a Collection structure for the given address.
func NewCollection(adr *common.Address) (*Collection, error) {
	col, err := repository.R().GetLegacyCollection(*adr)
	if err != nil {
		return nil, err
	}
	return (*Collection)(col), nil
}

// Cursor generates new unique identifier of the collection list edge.
func (edge CollectionEdge) Cursor() (types.Cursor, error) {
	return sorting.LegacyCollectionSortingName.GetCursor((*types.LegacyCollection)(edge.Node))
}

// NewCollectionConnection creates a new connection of a Collection list.
func NewCollectionConnection(list *types.LegacyCollectionList) (*CollectionConnection, error) {
	// create new connection
	con := &CollectionConnection{
		Edges:      make([]CollectionEdge, len(list.Collection)),
		TotalCount: (hexutil.Big)(*big.NewInt(list.TotalCount)),
		PageInfo:   PageInfo{},
	}

	// connect edges
	for i := 0; i < len(list.Collection); i++ {
		con.Edges[i].Node = (*Collection)(list.Collection[i])
	}

	// setup page info
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
	return con, nil
}

// Contract resolves thr address of the NFT collection contract.
func (t *Collection) Contract() common.Address {
	return t.Address
}

// Categories resolves list of Collection categories as a slice of PK indexes.
func (t *Collection) Categories() ([]int32, error) {
	return (*types.LegacyCollection)(t).CategoriesAsInt()
}

// Royalty returns percents of royalty fee as a string value.
func (t *Collection) Royalty() string {
	return (*types.LegacyCollection)(t).RoyaltyValue.String()
}

func (t *Collection) Discord() string {
	if idx := strings.Index(t.DiscordUrl, "discord.gg/"); idx != -1 {
		return "https://discord.gg/" + t.DiscordUrl[idx + 11:]
	}
	if t.DiscordUrl != "" {
		return "https://discord.gg/" + t.DiscordUrl
	}
	return ""
}

func (t *Collection) Site() string {
	if strings.HasPrefix(t.SiteUrl, "https://") || strings.HasPrefix(t.SiteUrl, "http://") {
		return t.SiteUrl
	}
	if t.SiteUrl != "" {
		return "https://" + t.SiteUrl
	}
	return ""
}

func (t *Collection) Telegram() string {
	if idx := strings.Index(t.TelegramUrl, "t.me/"); idx != -1 {
		return "https://t.me/" + t.TelegramUrl[idx + 5:]
	}
	if strings.HasPrefix(t.TelegramUrl, "@") && len(t.TelegramUrl) > 1 {
		return "https://t.me/" + t.TelegramUrl[1:]
	}
	if t.TelegramUrl != "" {
		return "https://t.me/" + t.TelegramUrl
	}
	return ""
}

func (t *Collection) Twitter() string {
	if idx := strings.Index(t.TwitterUrl, "twitter.com/"); idx != -1 {
		return "https://twitter.com/" + t.TwitterUrl[idx + 12:]
	}
	if strings.HasPrefix(t.TwitterUrl, "@") {
		return "https://twitter.com/" + t.TwitterUrl[1:]
	}
	if t.TwitterUrl != "" {
		return "https://twitter.com/" + t.TwitterUrl
	}
	return ""
}

func (t *Collection) Medium() string {
	if strings.HasPrefix(t.MediumUrl, "https://") || strings.HasPrefix(t.MediumUrl, "http://") {
		return t.MediumUrl
	}
	if strings.HasPrefix(t.MediumUrl, "@") {
		return "https://medium.com/" + t.MediumUrl
	}
	if t.MediumUrl != "" {
		return "https://medium.com/@" + t.MediumUrl
	}
	return ""
}

func (t *Collection) CreatedTime() types.Time {
	return types.Time(t.Id.Timestamp())
}

func (t *Collection) ChangedTime() *types.Time {
	if t.AppropriateUpdate.Unix() <= 0 {
		return nil
	}
	tm := types.Time(t.AppropriateUpdate)
	return &tm
}

func (t *Collection) OwnerUser() (*User, error) {
	return getUserByAddressPtr(t.Owner)
}

func (t *Collection) FeeRecipientUser() (*User, error) {
	return getUserByAddressPtr(t.FeeRecipient)
}

// CanMint resolves the minting privilege for the given user by address.
func (t *Collection) CanMint(args struct {
	User common.Address
	Fee  *hexutil.Big
}) (bool, error) {
	return repository.R().CanMint(&t.Address, &args.User, (*big.Int)(args.Fee))
}
