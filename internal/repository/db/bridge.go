// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository/db/registry"
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

	// sync DB related processes
	wg  sync.WaitGroup
	sig []chan bool
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
	br := &MongoDbBridge{
		client: con,
		dbName: cfg.Db.DbName,
		sig:    make([]chan bool, 0),
	}

	// update indexes to make sure we have the data in correct shape
	br.updateDatabaseIndexes()
	return br
}

// connectDb opens Mongo database connection
func connectDb(cfg *config.Database) (*mongo.Client, error) {
	// get empty unrestricted context
	ctx := context.Background()

	// create new Mongo client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Url).SetRegistry(registry.New()))
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
	// signal terminate and wait for processes to finish
	for _, sig := range db.sig {
		sig <- true
	}
	db.wg.Wait()

	// disconnect the db
	err := db.client.Disconnect(context.Background())
	if err != nil {
		log.Errorf("can not disconnect database; %s", err.Error())
	}
}
