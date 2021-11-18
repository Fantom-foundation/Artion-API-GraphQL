// Package repository implements persistent data access and processing.
package repository

import "artion-api-graphql/internal/types"

// TextSearchToken tries to search for tokens relevant to the given textual phrase.
func (p *Proxy) TextSearchToken(phrase string, limit int32) ([]*types.Token, error) {
	return p.db.TextSearchToken(phrase, limit)
}

// TextSearchCollection tries to search for collections relevant to the given textual phrase.
func (p *Proxy) TextSearchCollection(phrase string, limit int32) ([]*types.LegacyCollection, error) {
	return p.shared.TextSearchLegacyCollection(phrase, limit)
}

// TextSearchUser tries to search for user profiles relevant to the given textual phrase.
func (p *Proxy) TextSearchUser(phrase string, limit int32) ([]*types.User, error) {
	return p.shared.TextSearchUser(phrase, limit)
}
