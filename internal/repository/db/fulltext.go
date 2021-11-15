// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TextSearchToken uses existing text index to search list of tokens to get relevant records.
func (db *MongoDbBridge) TextSearchToken(phrase string, limit int32) ([]*types.Token, error) {
	col := db.client.Database(db.dbName).Collection(coTokens)

	// execute the search
	cur, err := col.Find(context.Background(), bson.D{
		{Key: "$text", Value: bson.D{
			{Key: "$search", Value: phrase},
			{Key: "$caseSensitive", Value: false},
			{Key: "$diacriticSensitive", Value: false},
		}},
	}, options.Find().SetLimit(int64(limit)))
	if err != nil {
		log.Errorf("fulltext search %s failed; %s", phrase, err.Error())
		return nil, err
	}

	defer closeFindCursor("token text search", cur)

	// load data
	list := make([]*types.Token, 0)
	for cur.Next(context.Background()) {
		var row types.Token

		if err := cur.Decode(&row); err != nil {
			log.Errorf("could not decode token; %s", err.Error())
			return nil, err
		}

		list = append(list, &row)
	}

	return list, nil
}

// TextSearchLegacyCollection uses existing text index to search legacy collection list to get relevant collections.
func (sdb *SharedMongoDbBridge) TextSearchLegacyCollection(phrase string, limit int32) ([]*types.LegacyCollection, error) {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)

	// execute the search
	cur, err := col.Find(context.Background(), bson.D{
		{Key: "$text", Value: bson.D{
			{Key: "$search", Value: phrase},
			{Key: "$caseSensitive", Value: false},
			{Key: "$diacriticSensitive", Value: false},
		}},
	}, options.Find().SetLimit(int64(limit)))
	if err != nil {
		log.Errorf("fulltext search %s failed; %s", phrase, err.Error())
		return nil, err
	}

	defer closeFindCursor("legacy collection text search", cur)

	// load data
	list := make([]*types.LegacyCollection, 0)
	for cur.Next(context.Background()) {
		var row types.LegacyCollection

		if err := cur.Decode(&row); err != nil {
			log.Errorf("could not decode collection; %s", err.Error())
			return nil, err
		}

		list = append(list, &row)
	}

	return list, nil
}
