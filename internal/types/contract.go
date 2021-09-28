// Package types provides high level structures for the API server.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	// ContractTypeERC721 represents an NFT contract type ERC-721
	ContractTypeERC721 = "erc721"

	// ContractTypeERC1155 represents an NFT contract type ERC-1155
	ContractTypeERC1155 = "erc1155"
)

// ObservedContract represents a contract observed by the API server.
type ObservedContract struct {
	Address     common.Address
	Type        string
	Created     time.Time
	Creator     common.Address
	BlockNumber uint64
	DeployedBy  common.Hash
}

// MarshalBSON provides BSON representation of the observed contract structure.
func (oc *ObservedContract) MarshalBSON() ([]byte, error) {
	return bson.Marshal(struct {
		Addr    string    `bson:"_id"`
		Type    string    `bson:"type"`
		Created time.Time `bson:"created"`
		Creator string    `bson:"creator"`
		Block   int64     `bson:"block"`
		Trx     string    `bson:"trx"`
	}{
		Addr:    oc.Address.String(),
		Type:    oc.Type,
		Created: oc.Created,
		Creator: oc.Creator.String(),
		Block:   int64(oc.BlockNumber),
		Trx:     oc.DeployedBy.String(),
	})
}

// UnmarshalBSON decodes a BSON representation of the observed contract into its instance.
func (oc *ObservedContract) UnmarshalBSON(data []byte) error {
	var row struct {
		Addr    string    `bson:"_id"`
		Type    string    `bson:"type"`
		Created time.Time `bson:"created"`
		Creator string    `bson:"creator"`
		Block   int64     `bson:"block"`
		Trx     string    `bson:"trx"`
	}
	if err := bson.Unmarshal(data, &row); err != nil {
		return err
	}
	oc.Address = common.HexToAddress(row.Addr)
	oc.Type = row.Type
	oc.Created = row.Created
	oc.Creator = common.HexToAddress(row.Creator)
	oc.BlockNumber = uint64(row.Block)
	oc.DeployedBy = common.HexToHash(row.Trx)
	return nil
}
