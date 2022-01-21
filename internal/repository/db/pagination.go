package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const fieldId = "_id"

// getTotalCount obtains count of documents in MongoDB collection corresponding the filter.
func (db *MongoDbBridge) getTotalCount(col *mongo.Collection, filter bson.D) (int64, error) {
	ctx := context.Background()
	start := time.Now()
	totalCount, err := col.CountDocuments(ctx, filter)
	duration := time.Now().Sub(start)
	if err != nil {
		log.Errorf("can not get total count; filter: %+v, %s", filter, err)
	}
	if duration > 100 * time.Millisecond {
		log.Infof("SlowMongo TotalCount dur=%s col=%s filter=%s", duration, col.Name(), filter)
	}
	return totalCount, err
}

// cursorToFilter converts cursor and sorting setting into Mongo filter
func cursorToFilter(cursor types.Cursor, sort sorting.Sorting, backward bool) (bson.E, error) {
	curParams, err := sorting.CursorToParams(cursor)
	if err != nil {
		return bson.E{}, err
	}

	cmpOp := "$gt" // greater than - for ascending sorting
	if backward {
		cmpOp = "$lt" // less than - for descending sorting
	}

	sortedField := sort.SortedFieldBson()
	ordinalField := sort.OrdinalFieldBson()

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
func (db *MongoDbBridge) findPaginated(col *mongo.Collection, filter bson.D, cursor types.Cursor, count int, sorting sorting.Sorting, backward bool) (mc *mongo.Cursor, err error) {
	ctx := context.Background()

	if cursor != "" {
		paginationFilter, err := cursorToFilter(cursor, sorting, backward)
		if err != nil {
			log.Errorf("unable to decode cursor %s; %s", cursor, err)
			return nil, err
		}
		filter = append(filter, paginationFilter)
	}

	opt := options.Find()
	direction := 1
	if backward {
		direction = -1
	}
	if sorting.SortedFieldBson() != "" {
		opt.SetSort(bson.D{{Key: sorting.SortedFieldBson(), Value: direction}, {Key: sorting.OrdinalFieldBson(), Value: direction}})
	} else {
		opt.SetSort(bson.D{{Key: sorting.OrdinalFieldBson(), Value: direction}})
	}
	opt.SetLimit(int64(count + 1))

	start := time.Now()
	mc, err = col.Find(ctx, filter, opt)
	duration := time.Now().Sub(start)
	if err != nil {
		log.Errorf("error loading list; %s", err.Error())
	}
	if duration > 100 * time.Millisecond {
		log.Infof("SlowMongo Find dur=%s col=%s filter=%s cursor=%s count=%s", duration, col.Name(), filter, cursor, count)
	}
	return mc, err
}
