// Package registry provides project specific BSON encoders and decoders.
package registry

import (
	"artion-api-graphql/internal/types"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
	"time"
)

// tTime represents the reflection type of the types.Time.
var tTime = reflect.TypeOf(types.Time{})

// TimeDecodeValue is the BSON value decoder for types.Time.
func TimeDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tTime {
		return bsoncodec.ValueDecoderError{Name: "TimeDecodeValue", Types: []reflect.Type{tTime}, Received: val}
	}

	var timeVal types.Time
	switch vrType := vr.Type(); vrType {
	case bsontype.Int64:
		i64, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		timeVal = types.Time(time.Unix(i64/1000, i64%1000*1000000))
	case bsontype.DateTime:
		dt, err := vr.ReadDateTime()
		if err != nil {
			return err
		}
		timeVal = types.Time(time.Unix(dt/1000, dt%1000*1000000))
	case bsontype.Timestamp:
		t, _, err := vr.ReadTimestamp()
		if err != nil {
			return err
		}
		timeVal = types.Time(time.Unix(int64(t), 0))
	case bsontype.Null:
		if err := vr.ReadNull(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot decode %v into a types.Time", vrType)
	}

	val.Set(reflect.ValueOf(timeVal))
	return nil
}

// TimeEncodeValue is the BSON value encoder for types.Time.
func TimeEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tTime {
		return bsoncodec.ValueEncoderError{Name: "TimeEncodeValue", Types: []reflect.Type{tTime}, Received: val}
	}
	tt := time.Time(val.Interface().(types.Time))
	return vw.WriteDateTime(tt.Unix()*1e3 + int64(tt.Nanosecond())/1e6)
}
