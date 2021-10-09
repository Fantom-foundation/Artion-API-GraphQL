package sorting

import (
	"artion-api-graphql/internal/repository/db/registry"
	"artion-api-graphql/internal/types"
	"encoding/base64"
	"go.mongodb.org/mongo-driver/bson"
)

var bsonRegistry = registry.New()

// CursorFromParams encodes map of params for pagination into cursor string.
func CursorFromParams(params map[string]interface{}) (cur types.Cursor, err error) {
	var bytes []byte
	bytes, err = bson.MarshalWithRegistry(bsonRegistry, params)
	if err != nil {
		return
	}
	cur = types.Cursor(base64.StdEncoding.EncodeToString(bytes))
	return
}

// CursorToParams decodes map of params for pagination stored in the cursor string.
func CursorToParams(cur types.Cursor) (params map[string]interface{}, err error) {
	bytes, err := base64.StdEncoding.DecodeString(string(cur))
	if err != nil {
		return
	}
	err = bson.UnmarshalWithRegistry(bsonRegistry, bytes, &params)
	return
}
