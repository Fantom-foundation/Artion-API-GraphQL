// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// addCollectionQueueCapacity is the capacity of the queue being filled with
// new collections to be added into the repository.
const addCollectionQueueCapacity = 10

// NewCollectionQueue provides queue filled with addresses of collection contracts
// to be added into the API.
func (p *Proxy) NewCollectionQueue() chan common.Address {
	return p.newCollectionQueue
}

// AddNewCollection pushes the given address to be validated and added as a new collection.
func (p *Proxy) AddNewCollection(adr *common.Address) {
	p.newCollectionQueue <- *adr
}

// CollectionName provides a name of an Artion ERC721 and/or ERC1155 token.
func (p *Proxy) CollectionName(adr *common.Address) (string, error) {
	return p.rpc.CollectionName(adr)
}

// CollectionSymbol provides a symbol of an Artion ERC721 and/or ERC1155 token.
func (p *Proxy) CollectionSymbol(adr *common.Address) (string, error) {
	return p.rpc.CollectionSymbol(adr)
}

// StoreCollection adds the specified NFT collection contract record into persistent storage.
func (p *Proxy) StoreCollection(nft *types.Collection) error {
	return p.db.StoreCollection(nft)
}

func (p *Proxy) GetCollection(address common.Address) (*types.Collection, error) {
	return p.db.GetCollection(address)
}

func (p *Proxy) ListCollections(cursor types.Cursor, count int, backward bool) (*types.CollectionList, error) {
	return p.db.ListCollections(cursor, count, backward)
}
