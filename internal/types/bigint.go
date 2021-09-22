package types

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"math/big"
)

// BigInt is type for 256-bit integers with BSON serialization support (in hex string).
type BigInt big.Int

func (bi *BigInt) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(hexutil.EncodeBig((*big.Int)(bi)))
}

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
