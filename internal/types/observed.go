// Package types provides high level structures for the API server.
package types

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	// ContractTypeERC721 represents an NFT contract type ERC-721
	ContractTypeERC721 = "erc721"

	// ContractTypeERC1155 represents an NFT contract type ERC-1155
	ContractTypeERC1155 = "erc1155"
)

// ObservedContract represents a contract observed by the API server.
type ObservedContract struct {
	Address     common.Address  `bson:"_id"`
	Name        string          `bson:"name"`
	Type        string          `bson:"type"`
	Created     Time            `bson:"created"`
	Creator     common.Address  `bson:"creator"`
	BlockNumber uint64          `bson:"block"`
	DeployedBy  common.Hash     `bson:"trx"`
}
