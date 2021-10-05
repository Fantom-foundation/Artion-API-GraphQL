// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// CollectionName provides a name of an Artion ERC721 and/or ERC1155 token.
func (p *Proxy) CollectionName(adr *common.Address) (string, error) {
	return p.rpc.CollectionName(adr)
}

// CollectionSymbol provides a symbol of an Artion ERC721 and/or ERC1155 token.
func (p *Proxy) CollectionSymbol(adr *common.Address) (string, error) {
	return p.rpc.CollectionSymbol(adr)
}

// AddCollection adds the specified NFT collection contract record into persistent storage.
func (p *Proxy) AddCollection(nft *types.Collection) error {
	return p.db.AddCollection(nft)
}

func (p *Proxy) GetCollection(address common.Address) (*types.Collection, error) {
	return p.db.GetCollection(address)
}

func (p *Proxy) ListCollections(cursor types.Cursor, count int, backward bool) (*types.CollectionList, error) {
	return p.db.ListCollections(cursor, count, backward)
}
