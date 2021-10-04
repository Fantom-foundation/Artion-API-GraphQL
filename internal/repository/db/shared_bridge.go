// Package db provides access to the persistent storage.
package db

import (
	"context"
	"sync"
)
import "go.mongodb.org/mongo-driver/mongo"

// SharedMongoDbBridge represents Mongo DB abstraction layer.
type SharedMongoDbBridge struct {
	client *mongo.Client
	dbName string

	// init state marks
	initUsers       *sync.Once
}

// NewShared creates a new Mongo Db connection bridge.
func NewShared() *SharedMongoDbBridge {
	// log what we do
	log.Debugf("connecting database at %s/%s", cfg.SharedDb.Url, cfg.SharedDb.DbName)

	// open the database connection
	con, err := connectDb(&cfg.SharedDb)
	if err != nil {
		log.Criticalf("can not contact the database; %s", err.Error())
		return nil
	}

	log.Notice("database connection is open")
	return &SharedMongoDbBridge{
		client: con,
		dbName: cfg.SharedDb.DbName,
	}
}

// Close terminates the database connection.
func (db *SharedMongoDbBridge) Close() {
	err := db.client.Disconnect(context.Background())
	if err != nil {
		log.Errorf("can not disconnect database; %s", err.Error())
	}
}
