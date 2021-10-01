// Package db provides access to the persistent storage.
package db

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
)

var (
	// tAddress represents the reflection type of the Opera blockchain Address.
	tAddress = reflect.TypeOf(common.Address{})

	// tHash represents the reflection type of the Opera blockchain generic Hash.
	tHash = reflect.TypeOf(common.Hash{})

	// tHexBigInt represents the reflection type of the hexutil.Big integer.
	tHexBigInt = reflect.TypeOf(hexutil.Big{})
)

// AddressBSONEncodeValue encodes Opera blockchain address into BSON data stream.
func AddressBSONEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tAddress {
		return bsoncodec.ValueEncoderError{Name: "AddressBSONEncodeValue", Types: []reflect.Type{tAddress}, Received: val}
	}
	return vw.WriteString(val.Interface().(common.Address).String())
}

// AddressBSONDecodeValue decodes Opera blockchain address from BSON data stream.
func AddressBSONDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tAddress {
		return bsoncodec.ValueDecoderError{Name: "AddressBSONDecodeValue", Types: []reflect.Type{tAddress}, Received: val}
	}

	var adr common.Address
	switch vrType := vr.Type(); vrType {
	case bsontype.String:
		str, err := vr.ReadString()
		if err != nil {
			return err
		}
		adr = common.HexToAddress(str)
	default:
		return vr.ReadUndefined()
	}

	val.Set(reflect.ValueOf(adr).Elem())
	return nil
}

// HashBSONEncodeValue encodes Opera blockchain hash into BSON data stream.
func HashBSONEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tHash {
		return bsoncodec.ValueEncoderError{Name: "HashBSONEncodeValue", Types: []reflect.Type{tHash}, Received: val}
	}
	return vw.WriteString(val.Interface().(common.Hash).String())
}

// HashBSONDecodeValue decodes Opera blockchain hash from BSON data stream.
func HashBSONDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tHash {
		return bsoncodec.ValueDecoderError{Name: "HashBSONDecodeValue", Types: []reflect.Type{tHash}, Received: val}
	}

	var hash common.Hash
	switch vrType := vr.Type(); vrType {
	case bsontype.String:
		str, err := vr.ReadString()
		if err != nil {
			return err
		}
		hash = common.HexToHash(str)
	default:
		return vr.ReadUndefined()
	}

	val.Set(reflect.ValueOf(hash).Elem())
	return nil
}

// HexBigIntBSONEncodeValue encodes hexutil.Big into BSON data stream.
func HexBigIntBSONEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tHexBigInt {
		return bsoncodec.ValueEncoderError{Name: "HexBigIntBSONEncodeValue", Types: []reflect.Type{tHexBigInt}, Received: val}
	}
	v := val.Interface().(hexutil.Big)
	return vw.WriteString((&v).String())
}

// HexBigIntBSONDecodeValue decodes hexutil.Big from BSON data stream.
func HexBigIntBSONDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tHexBigInt {
		return bsoncodec.ValueDecoderError{Name: "HexBigIntBSONDecodeValue", Types: []reflect.Type{tHexBigInt}, Received: val}
	}

	var big hexutil.Big
	switch vrType := vr.Type(); vrType {
	case bsontype.String:
		str, err := vr.ReadString()
		if err != nil {
			return err
		}

		v, err := hexutil.DecodeBig(str)
		if err != nil {
			return err
		}

		big = hexutil.Big(*v)
	default:
		return vr.ReadUndefined()
	}

	val.Set(reflect.ValueOf(big).Elem())
	return nil
}
