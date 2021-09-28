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

// Token represents item list-able in the marketplace.
type Token struct {
	Id            []byte
	Nft           common.Address
	TokenId       hexutil.Big
	Uri           string
}

func TokenIdFromAddress(Nft *common.Address, TokenId *big.Int) []byte {
	hash := md5.New()
	hash.Write(Nft.Bytes())
	hash.Write(TokenId.Bytes())
	return hash.Sum(nil)
}

func (t *Token) GenerateId() {
	t.Id = TokenIdFromAddress(&t.Nft, (*big.Int)(&t.TokenId))
}

type tokenBson struct {
	Id           []byte   `bson:"_id"`
	Nft          string   `bson:"nft"`
	TokenId      string   `bson:"tokenId"`
	Name         string   `bson:"name"`
	Description  string   `bson:"description"`
	Uri          string   `bson:"uri"`
}

// MarshalBSON prepares data to be stored in MongoDB.
func (t *Token) MarshalBSON() ([]byte, error) {
	return bson.Marshal(tokenBson{
		Id:          t.Id,
		Nft:         t.Nft.String(),
		TokenId:     (&t.TokenId).String(),
		Uri:         t.Uri,
	})
}

// UnmarshalBSON parses data from MongoDB.
func (t *Token) UnmarshalBSON(data []byte) (err error) {
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

	t.Id = row.Id
	t.Nft = common.HexToAddress(row.Nft)
	t.TokenId = (hexutil.Big)(*hexutil.MustDecodeBig(row.TokenId))
	t.Uri = row.Uri
	return nil
}
