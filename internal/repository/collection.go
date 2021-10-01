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

// AddNFTCollection adds the specified NFT collection contract record into persistent storage.
func (p *Proxy) AddNFTCollection(nft *types.NFTCollection) error {
	return p.db.AddNFTCollection(nft)
}
