package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const fieldId = "_id"

// getTotalCount obtains count of documents in MongoDB collection corresponding the filter.
func (db *MongoDbBridge) getTotalCount(col *mongo.Collection, filter *bson.D) (int64, error) {
	ctx := context.Background()
	totalCount, err := col.CountDocuments(ctx, *filter)
	if err != nil {
		log.Errorf("can not get total count")
	}
	return totalCount, err
}

func cursorToFilter(cursor types.Cursor, sorting sorting.Sorting, backward bool) (bson.E, error) {
	curParams, err := cursor.ToParams()
	if err != nil {
		return bson.E{}, err
	}

	x, _ := json.Marshal(curParams)
	log.Errorf("curParams %s", x)

	cmpOp := "$gt" // greater than - for ascending sorting
	if backward {
		cmpOp = "$lt" // less than - for descending sorting
	}

	sortedField := sorting.SortedFieldBson()
	ordinalField := sorting.OrdinalFieldBson()

	if sortedField == "" {
		return bson.E{ Key: ordinalField, Value: bson.D{{cmpOp, curParams[ordinalField] }} }, nil
	}

	return bson.E{ Key: "$or", Value: bson.A{
		bson.D{{sortedField, bson.D{{cmpOp, curParams[sortedField] }} }},
		bson.E{ Key: "$and", Value: bson.A{
			bson.D{{sortedField, bson.D{{"$eq", curParams[sortedField] }} }},
			bson.D{{ordinalField, bson.D{{cmpOp, curParams[ordinalField] }} }},
		}},
	}}, nil
}

// findPaginated obtains one page of filtered results from given collection of MongoDB.
func (db *MongoDbBridge) findPaginated(col *mongo.Collection, inputFilter *bson.D, cursor types.Cursor, count int, sorting sorting.Sorting, backward bool) (mc *mongo.Cursor, err error) {
	ctx := context.Background()
	filter := *inputFilter

	if cursor != "" {
		paginationFilter, err := cursorToFilter(cursor, sorting, backward)
		if err != nil {
			log.Errorf("unable to decode cursor %s; %s", cursor, err)
			return nil, err
		}
		filter = append(filter, paginationFilter)
	}

	x, _ := json.Marshal(filter)
	log.Errorf("FILTER %s", x)

	opt := options.Find()
	direction := 1
	if backward {
		direction = -1
	}
	if sorting.SortedFieldBson() != "" {
		opt.SetSort(bson.D{{Key: sorting.SortedFieldBson(), Value: direction}, {Key: fieldId, Value: direction}})
	} else {
		opt.SetSort(bson.D{{Key: fieldId, Value: direction}})
	}
	opt.SetLimit(int64(count + 1))

	mc, err = col.Find(ctx, filter, opt)
	if err != nil {
		log.Errorf("error loading list; %s", err.Error())
	}
	return mc, err
}
