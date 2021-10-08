package types

import (
	"encoding/base64"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

// Cursor represents a string key of an element position in a sequential list of edges.
type Cursor string

// CursorFromId converts bytes of MongoDB field value GraphQL Cursor type.
func CursorFromId(id []byte) Cursor {
	return Cursor(hexutil.Encode(id))
}

// CursorFromObjectId converts MongoDB ObjectID to GraphQL Cursor type.
func CursorFromObjectId(id primitive.ObjectID) Cursor {
	return Cursor(hexutil.Encode(id[:]))
}

func (c Cursor) Bytes() ([]byte, error) {
	return hexutil.Decode(string(c))
}

func (c Cursor) ToObjectId() (primitive.ObjectID, error) {
	id, err := hexutil.Decode(string(c))
	if err != nil {
		return primitive.ObjectID{}, err
	}
	var bytes [12]byte
	copy(bytes[:], id)
	return bytes, nil
}

// CursorFromParams encodes map of params for pagination into cursor string.
func CursorFromParams(params map[string]interface{}) (cur Cursor, err error) {
	var bytes []byte
	bytes, err = bson.Marshal(params)
	if err != nil { return }
	cur = Cursor(base64.StdEncoding.EncodeToString(bytes))
	return
}

// ToParams decodes map of params for pagination stored in the cursor string.
func (c Cursor) ToParams() (params map[string]interface{}, err error) {
	bytes, err := base64.StdEncoding.DecodeString(string(c))
	if err != nil { return }
	err = bson.Unmarshal(bytes, &params)
	return
}

// ImplementsGraphQLType notifies the GraphQL that this type resolves Cursor scalar.
func (Cursor) ImplementsGraphQLType(name string) bool {
	return name == "Cursor"
}

// UnmarshalGraphQL unmarshal incoming Cursor into a local variable.
func (c *Cursor) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		*c = Cursor(input)
	case int32:
		*c = Cursor(strconv.Itoa(int(input)))
	default:
		err = errors.New("wrong cursor type")
	}
	return err
}

// MarshalJSON encodes a cursor to JSON for transport.
func (c Cursor) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(c)), nil
}
