package types

import (
	"crypto/md5"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
)

const (
	// CoTokens is the name of database collection.
	CoTokens = "Tokens"

	// BSON attributes names used in the database collection.
	FiTokenNft     = "nft"
	FiTokenTokenId = "tokenId"
	FiTokenName    = "name"
)

// Token represents item listed in marketplace.
type Token struct {
	Id            []byte
	Nft           common.Address
	TokenId       hexutil.Big
	Name          string
	Description   string
}

func (t *Token) GenerateId() {
	hash := md5.New()
	hash.Write(t.Nft.Bytes())
	hash.Write((*big.Int)(&t.TokenId).Bytes())
	t.Id = hash.Sum(nil)
}

type tokenBson struct {
	Id           []byte   `bson:"_id"`
	Nft          string   `bson:"nft"`
	TokenId      string   `bson:"tokenId"`
	Name         string   `bson:"name"`
	Description  string   `bson:"description"`
}

// MarshalBSON prepares data to be stored in MongoDB.
func (t *Token) MarshalBSON() ([]byte, error) {
	return bson.Marshal(tokenBson{
		Id:          t.Id,
		Nft:         t.Nft.String(),
		TokenId:     (&t.TokenId).String(),
		Name:        t.Name,
		Description: t.Description,
	})
}

// UnmarshalBSON parses data from MongoDB.
func (ev *Token) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode TokenEvent; %s", err.Error())
		}
	}()

	// try to decode the BSON data
	var row tokenBson
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	ev.Id = row.Id
	ev.Nft = common.HexToAddress(row.Nft)
	ev.TokenId = (hexutil.Big)(*hexutil.MustDecodeBig(row.TokenId))
	ev.Name = row.Name
	ev.Description = row.Description
	return nil
}
