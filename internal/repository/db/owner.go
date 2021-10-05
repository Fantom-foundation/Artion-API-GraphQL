// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	// coTokenOwners is the name of the collection keeping token ownership.
	coTokenOwners = "owners"

	// coTokenOwnersQueryTimeout represents the timeout applied to owners collection queries.
	coTokenOwnersQueryTimeout = 10 * time.Second
)

// StoreOwner updates NFT ownership in the persistent storage. If the Qty dropped to zero,
// the record is deleted.
func (db *MongoDbBridge) StoreOwner(to *types.NFTOwner) error {
	if to == nil {
		return fmt.Errorf("no value to store")
	}

	// remove record with zero Qty
	if 0 == to.Qty.ToInt().Uint64() {
		return db.DeleteOwner(to)
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coTokenOwners)
	ctx, cancel := context.WithTimeout(context.Background(), coTokenOwnersQueryTimeout)
	defer func() {
		cancel()
	}()

	// try to do the insert
	id := to.ID()
	rs, err := col.UpdateOne(
		ctx,
		bson.D{{Key: fieldId, Value: id}},
		bson.D{{Key: "$set", Value: to}, {Key: "$setOnInsert", Value: id}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Errorf("can not store owner %s of %s / #%s; %s",
			to.Owner.String(), to.Contract.String(), to.TokenID.String(), err.Error())
		return err
	}
	if rs.UpsertedCount > 0 {
		log.Infof("token %s / #%s owner updated to %s", to.Contract.String(), to.TokenID.String(), to.Owner.String())
	}
	return nil
}

// DeleteOwner removes NFT ownership from the persistent storage.
// We do this if the token qty drops to zero on the owner address.
func (db *MongoDbBridge) DeleteOwner(to *types.NFTOwner) error {
	if to == nil {
		return fmt.Errorf("no value to store")
	}

	col := db.client.Database(db.dbName).Collection(coTokenOwners)
	ctx, cancel := context.WithTimeout(context.Background(), coTokenOwnersQueryTimeout)
	defer func() {
		cancel()
	}()

	dr, err := col.DeleteOne(ctx, bson.D{{Key: fieldId, Value: to.ID()}})
	if err != nil {
		log.Errorf("can not delete owner %s of %s / #%s; %s",
			to.Owner.String(), to.Contract.String(), to.TokenID.String(), err.Error())
		return err
	}

	if dr.DeletedCount > 0 {
		log.Infof("token %s / #%s owner %s deleted", to.Contract.String(), to.TokenID.String(), to.Owner.String())
	}
	return nil
}
