// Package db provides access to the persistent storage.
package db

// SharedMongoDbBridge represents MongoDbBridge attached to database shared/replicated between nodes.
type SharedMongoDbBridge MongoDbBridge

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
		sig:    make([]chan bool, 0),
	}
}

// Close terminates the database connection.
func (sdb *SharedMongoDbBridge) Close() {
	(*MongoDbBridge)(sdb).Close()
}
