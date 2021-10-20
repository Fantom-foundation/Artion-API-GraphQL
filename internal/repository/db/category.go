package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// coCategories is the name of database collection.
	coCategories = "colcats"

	// fiCategoryName is the DB column of the category having name of the category.
	fiCategoryName = "name"
)

func (sdb *SharedMongoDbBridge) ListCategories() (out []types.Category, err error) {
	col := sdb.client.Database(sdb.dbName).Collection(coCategories)
	ctx := context.Background()

	mc, err := col.Find(ctx, bson.D{}, options.Find().SetSort(bson.D{{Key: fiCategoryName, Value: 1}}))
	if err != nil {
		log.Errorf("error loading categories list; %s", err.Error())
		return nil, err
	}
	defer func() {
		err = mc.Close(ctx)
		if err != nil {
			log.Errorf("error closing categories list cursor; %s", err.Error())
		}
	}()

	for mc.Next(ctx) {
		var row types.Category
		if err = mc.Decode(&row); err != nil {
			log.Errorf("can not decode the category in list; %s", err.Error())
			return nil, err
		}
		out = append(out, row)
	}
	return out, nil
}
