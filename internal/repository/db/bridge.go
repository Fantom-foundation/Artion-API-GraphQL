package db

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)
import "go.mongodb.org/mongo-driver/mongo"

// MongoDbBridge represents Mongo DB abstraction layer.
type MongoDbBridge struct {
	client *mongo.Client
	log    logger.Logger
	dbName string

	// init state marks
	initTokenEvents *sync.Once
}

// New creates a new Mongo Db connection bridge.
func New(cfg *config.Config, log logger.Logger) (*MongoDbBridge, error) {
	// log what we do
	log.Debugf("connecting database at %s/%s", cfg.Db.Url, cfg.Db.DbName)

	// open the database connection
	con, err := connectDb(&cfg.Db)
	if err != nil {
		log.Criticalf("can not contact the database; %s", err.Error())
		return nil, err
	}

	// log the event
	log.Notice("database connection established")

	// return the bridge
	db := &MongoDbBridge{
		client: con,
		log:    log,
		dbName: cfg.Db.DbName,
	}
	return db, nil
}

// connectDb opens Mongo database connection
func connectDb(cfg *config.Database) (*mongo.Client, error) {
	// get empty unrestricted context
	ctx := context.Background()

	// create new Mongo client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Url))
	if err != nil {
		return nil, err
	}

	// validate the connection was indeed established
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
