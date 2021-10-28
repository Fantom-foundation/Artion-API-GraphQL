// Package db provides access to the persistent storage.
package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

// indexListQueueCapacity represents the capacity of index lists queue.
const indexListQueueCapacity = 10

// IndexListProvider represents a function providing list of indexes.
type IndexListProvider func() []mongo.IndexModel

// IndexList represents a prescribed list of indexes for the given collection
type IndexList struct {
	Collection *mongo.Collection
	Indexes    []mongo.IndexModel
}

// updateDatabaseIndexes checks for indexes existence; if an expected index is not found, it creates it.
func (db *MongoDbBridge) updateDatabaseIndexes() {
	// define index list loaders
	var ixLoaders = map[string]IndexListProvider{
		coActivities:      IndexDefinitionActivities,
		coAuctions:        IndexDefinitionAuctions,
		coAuctionBids:     IndexDefinitionAuctionBids,
		coCollection:      IndexDefinitionCollections,
		coListings:        IndexDefinitionListings,
		coOffers:          IndexDefinitionOffers,
		coTokenOwnerships: IndexDefinitionOwnership,
		coTokens:          IndexDefinitionTokens,
		coUsers:           IndexDefinitionUsers,
	}

	// the DB bridge needs a way to terminate this thread
	sig := make(chan bool, 1)
	db.sig = append(db.sig, sig)

	// prep queue and start the updater
	iq := make(chan *IndexList, indexListQueueCapacity)
	db.wg.Add(1)
	go db.indexUpdater(iq, sig)

	// check indexes
	for cn, ld := range ixLoaders {
		iq <- &IndexList{
			Collection: db.client.Database(db.dbName).Collection(cn),
			Indexes:    ld(),
		}
	}

	// close the channel as no more updates will be sent
	close(iq)
}

// indexUpdater processes queue of index updates.
func (db *MongoDbBridge) indexUpdater(iq chan *IndexList, sig chan bool) {
	defer func() {
		db.wg.Done()
		log.Noticef("db index updater closed")
	}()

	for {
		select {
		case <-sig:
			return
		case ix, more := <-iq:
			// is this the end?
			if !more {
				log.Noticef("all indexes processed")
				return
			}

			// do the update
			err := db.updateIndexes(ix.Collection, ix.Indexes, sig)
			if err != nil {
				log.Errorf("%s index list sync failed; %s", ix.Collection.Name(), err.Error())
			}
		}
	}
}

// updateIndexes synchronizes indexes of the given DB collection with the given prescription.
func (db *MongoDbBridge) updateIndexes(col *mongo.Collection, list []mongo.IndexModel, sig chan bool) error {
	view := col.Indexes()

	// load list of all indexes known
	known, err := view.ListSpecifications(context.Background())
	if err != nil {
		return err
	}

	// loop prescriptions and make sure each index exists by name
	for _, ix := range list {
		// respect possible terminate signal
		select {
		case <-sig:
			sig <- true
			return nil
		default:
		}

		err = db.updateIndex(col, &view, ix, known)
		if err != nil {
			log.Errorf(err.Error())
		}
	}
	return nil
}

// updateIndex checks specific index model and creates the index on the DB collection if needed.
func (db *MongoDbBridge) updateIndex(col *mongo.Collection, view *mongo.IndexView, ix mongo.IndexModel, known []*mongo.IndexSpecification) error {
	// throw if index is not explicitly named
	if ix.Options.Name == nil {
		return fmt.Errorf("index name not defined on %s", col.Name())
	}

	// do we know the index?
	var ok = false
	for _, spec := range known {
		if spec.Name == *ix.Options.Name {
			ok = true
			break
		}
	}

	if !ok {
		ina, err := view.CreateOne(context.Background(), ix)
		if err != nil {
			return fmt.Errorf("failed to create index %s on %s", *ix.Options.Name, col.Name())
		}
		log.Noticef("created index %s on %s", ina, col.Name())
	}
	return nil
}
