package db

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)
import "go.mongodb.org/mongo-driver/mongo"

// config represents the configuration setup used by the repository
// to establish and maintain required connectivity to external services
// as needed.
var cfg *config.Config

// log represents the logger to be used by the repository.
var log logger.Logger

// MongoDbBridge represents Mongo DB abstraction layer.
type MongoDbBridge struct {
	client *mongo.Client
	dbName string

	// init state marks
	initTokenEvents *sync.Once
}

// New creates a new Mongo Db connection bridge.
func New() *MongoDbBridge {
	// log what we do
	log.Debugf("connecting database at %s/%s", cfg.Db.Url, cfg.Db.DbName)

	// open the database connection
	con, err := connectDb(&cfg.Db)
	if err != nil {
		log.Criticalf("can not contact the database; %s", err.Error())
		return nil
	}

	log.Notice("database connection is open")
	return &MongoDbBridge{
		client: con,
		dbName: cfg.Db.DbName,
	}
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

// SetConfig sets the repository configuration to be used to establish
// and maintain external repository connections.
func SetConfig(c *config.Config) {
	cfg = c
}

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l.ModuleLogger("db")
}

// Close terminates the database connection.
func (db *MongoDbBridge) Close() {
	err := db.client.Disconnect(context.Background())
	if err != nil {
		log.Errorf("can not disconnect database; %s", err.Error())
	}
}
