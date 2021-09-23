package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"math/big"
	"strconv"
)

// BigInt is type for 256-bit integers with BSON serialization support (hex in graphql).
type BigInt big.Int

// MarshalBSONValue converts the number into form to be stored in MongoDB (hex string).
func (bi *BigInt) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(hexutil.EncodeBig((*big.Int)(bi)))
}

// UnmarshalBSONValue converts the number from form in which it is stored in MongoDB (hex string).
func (bi *BigInt) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	rv := bson.RawValue{Type: t, Value: data}
	var hexa string
	err := rv.Unmarshal(&hexa)
	if err != nil {
		return err
	}
	out, err := hexutil.DecodeBig(hexa)
	if err != nil {
		return err
	}
	(*big.Int)(bi).Set(out)
	return nil
}

func (BigInt) ImplementsGraphQLType(name string) bool {
	return name == "BigInt"
}

// UnmarshalGraphQL converts the number from form to be used in GraphQL (hex string).
func (bi *BigInt) UnmarshalGraphQL(input interface{}) error {
	if hexa, ok := input.(string); ok {
		out, err := hexutil.DecodeBig(hexa)
		if err != nil {
			return err
		}
		(*big.Int)(bi).Set(out)
		return nil
	}
	return fmt.Errorf("unable to unmarshal type %T into BigInt", input)
}

// MarshalJSON converts the number into form to be used in GraphQL (hex string).
func (bi BigInt) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, hexutil.EncodeBig((*big.Int)(&bi))), nil
}

// ToFilter converts value to type to be used in Mongo query filter
func (bi *BigInt) ToFilter() string {
	return hexutil.EncodeBig((*big.Int)(bi))
}
