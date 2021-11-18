// Package resolvers implements GraphQL resolvers and resolvable objects for the API.
package resolvers

import "artion-api-graphql/internal/repository"

const (
	// defaultTextSearchOutputLimit represents the default length of text search export list per type
	defaultTextSearchOutputLimit = 15

	// maxTextSearchOutputLimit represents the maximal length of text search export list per type
	maxTextSearchOutputLimit = 100

	// minTextSearchPhraseLength is the minimal accepted text search phrase length
	minTextSearchPhraseLength = 3
)

// TextSearchEdge represents a single row of the TextSearch result set.
type TextSearchEdge struct {
	Collection *Collection
	Token      *Token
	User       *User
}

// TextSearch executes fulltext search on data subset using Mongo text index.
func (rs *RootResolver) TextSearch(args struct {
	Phrase    string
	MaxLength int32
}) ([]*TextSearchEdge, error) {
	// limit output size to the predefined range
	if args.MaxLength < 1 || args.MaxLength > maxTextSearchOutputLimit {
		args.MaxLength = defaultTextSearchOutputLimit
	}

	// make sure the phrase length is at least the minimal runes
	if len([]rune(args.Phrase)) < minTextSearchPhraseLength {
		return make([]*TextSearchEdge, 0), nil
	}

	collections, err := repository.R().TextSearchCollection(args.Phrase, args.MaxLength)
	if err != nil {
		return nil, err
	}

	tokens, err := repository.R().TextSearchToken(args.Phrase, args.MaxLength)
	if err != nil {
		return nil, err
	}

	users, err := repository.R().TextSearchUser(args.Phrase, args.MaxLength)
	if err != nil {
		return nil, err
	}

	// make the output list
	list := make([]*TextSearchEdge, len(collections)+len(tokens)+len(users))
	for i, c := range collections {
		list[i] = &TextSearchEdge{
			Collection: (*Collection)(c),
		}
	}

	offset := len(collections)
	for i, t := range tokens {
		list[offset+i] = &TextSearchEdge{
			Token: (*Token)(t),
		}
	}

	offset = offset + len(tokens)
	for i, u := range users {
		list[offset+i] = &TextSearchEdge{
			User: &User{Address: u.Address, dbUser: u},
		}
	}
	return list, nil
}
