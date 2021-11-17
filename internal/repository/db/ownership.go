// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	// coTokenOwnerships is the name of the collection keeping token ownership.
	coTokenOwnerships = "ownerships"

	// fiOwnershipContract is the name of the DB column of the NFT contract address.
	fiOwnershipContract = "contract"

	// fiOwnershipTokenId is the name of the DB column of the token ID.
	fiOwnershipTokenId = "token"

	// fiOwnershipOwner is the name of the DB column of the token owner address.
	fiOwnershipOwner = "owner"

	// coTokenOwnershipsQueryTimeout represents the timeout applied to owners collection queries.
	coTokenOwnershipsQueryTimeout = 10 * time.Second
)

// StoreOwnership updates NFT ownership in the persistent storage. If the Qty dropped to zero,
// the record is deleted.
func (db *MongoDbBridge) StoreOwnership(to *types.Ownership) error {
	if to == nil {
		return fmt.Errorf("no value to store")
	}

	// remove record with zero Qty
	if 0 == to.Qty.ToInt().Uint64() {
		return db.DeleteOwnership(to)
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coTokenOwnerships)
	ctx, cancel := context.WithTimeout(context.Background(), coTokenOwnershipsQueryTimeout)
	defer func() {
		cancel()
	}()

	// try to do the insert
	id := to.ID()
	rs, err := col.UpdateOne(
		ctx,
		bson.D{{Key: fieldId, Value: id}},
		bson.D{
			{Key: "$set", Value: to},
			{Key: "$setOnInsert", Value: bson.D{{
				Key: fieldId, Value: id,
			}}},
		},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Errorf("can not store ownership %s of %s / #%s; %s",
			to.Owner.String(), to.Contract.String(), to.TokenId.String(), err.Error())
		return err
	}
	if rs.UpsertedCount > 0 {
		log.Debugf("token %s / #%s ownership updated to %s", to.Contract.String(), to.TokenId.String(), to.Owner.String())

		err = db.UpdateOffersOwners(to.Contract, to.TokenId)
		if err != nil {
			log.Errorf("failed to update offers owners for %s/%s; %s",
				to.Contract.String(), to.TokenId.String(), err.Error())
			return err
		}
	}
	return nil
}

// DeleteOwnership removes NFT ownership from the persistent storage.
// We do this if the token qty drops to zero on the owner address.
func (db *MongoDbBridge) DeleteOwnership(to *types.Ownership) error {
	if to == nil {
		return fmt.Errorf("no value to store")
	}

	col := db.client.Database(db.dbName).Collection(coTokenOwnerships)
	ctx, cancel := context.WithTimeout(context.Background(), coTokenOwnershipsQueryTimeout)
	defer func() {
		cancel()
	}()

	dr, err := col.DeleteOne(ctx, bson.D{{Key: fieldId, Value: to.ID()}})
	if err != nil {
		log.Errorf("can not delete ownership %s of %s / #%s; %s",
			to.Owner.String(), to.Contract.String(), to.TokenId.String(), err.Error())
		return err
	}

	if dr.DeletedCount > 0 {
		log.Debugf("token %s / #%s ownership by %s deleted", to.Contract.String(), to.TokenId.String(), to.Owner.String())

		err = db.UpdateOffersOwners(to.Contract, to.TokenId)
		if err != nil {
			log.Errorf("failed to update offers owners for %s/%s; %s",
				to.Contract.String(), to.TokenId.String(), err.Error())
			return err
		}
	}
	return nil
}

func (db *MongoDbBridge) IsOwnerOf(contract common.Address, tokenId hexutil.Big, owner common.Address) (bool, error) {
	filter := bson.D{
		{Key: fiOwnershipContract, Value: contract.String()},
		{Key: fiOwnershipTokenId, Value: tokenId.String()},
		{Key: fiOwnershipOwner, Value: owner.String()},
	}
	col := db.client.Database(db.dbName).Collection(coTokenOwnerships)
	count, err := db.getTotalCount(col, filter)
	return count > 0, err
}

// GetTokenOwners provides list of owners of given token
func (db *MongoDbBridge) GetTokenOwners(contract common.Address, tokenId hexutil.Big) ([]common.Address, error) {
	col := db.client.Database(db.dbName).Collection(coTokenOwnerships)
	ctx := context.Background()
	var list []common.Address

	ld, err := col.Find(ctx, bson.D{
		{Key: fiOwnershipContract, Value: contract.String()},
		{Key: fiOwnershipTokenId, Value: tokenId.String()},
	})
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			log.Errorf("error closing all token owners cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		var ownership types.Ownership
		if err = ld.Decode(&ownership); err != nil {
			log.Errorf("can not decode ownerships in list; %s", err.Error())
			return nil, err
		}
		list = append(list, ownership.Owner)
	}
	return list, nil
}

func (db *MongoDbBridge) ListOwnerships(contract *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor types.Cursor, count int, backward bool) (out *types.OwnershipList, err error) {
	filter := bson.D{}
	if contract != nil {
		filter = append(filter, primitive.E{Key: fiOwnershipContract, Value: contract.String()})
	}
	if tokenId != nil {
		filter = append(filter, primitive.E{Key: fiOwnershipTokenId, Value: tokenId.String()})
	}
	if owner != nil {
		filter = append(filter, primitive.E{Key: fiOwnershipOwner, Value: owner.String()})
	}
	return db.listOwnerships(filter, cursor, count, backward)
}

func (db *MongoDbBridge) listOwnerships(filter bson.D, cursor types.Cursor, count int, backward bool) (out *types.OwnershipList, err error) {
	var list types.OwnershipList
	col := db.client.Database(db.dbName).Collection(coTokenOwnerships)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.OwnershipSortingNone, backward)
	if err != nil {
		log.Errorf("error loading ownerships list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			log.Errorf("error closing ownerships list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.Ownership
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the ownerships in list; %s", err.Error())
				return nil, err
			}
			list.Collection = append(list.Collection, &row)
		} else {
			list.HasNext = true
		}
	}

	if cursor != "" {
		list.HasPrev = true
	}
	if backward {
		list.Reverse()
	}
	return &list, nil
}
