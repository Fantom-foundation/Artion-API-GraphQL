package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// initOfferCollection initializes collection with indexes and additional parameters.
func (db *MongoDbBridge) initOfferCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index sender and recipient
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiOfferNft, Value: 1}, {Key: types.FiOfferTokenId, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiOfferCreator, Value: 1}}})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		log.Panicf("can not create indexes for transaction collection; %s", err.Error())
	}

	// log we are done that
	log.Debugf("transactions collection initialized")
}

func (db *MongoDbBridge) AddOffer(event *types.Offer) error {
	if event == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(types.CoOffers)

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), event); err != nil {
		log.Errorf("can not store Offer; %s", err)
		return err
	}
	// make sure gas price collection is initialized
	if db.initOffers != nil {
		db.initOffers.Do(func() { db.initOfferCollection(col); db.initOffers = nil })
	}
	return nil
}

func (db *MongoDbBridge) UpdateOffer(Offer *types.Offer) error {
	if Offer == nil {
		return fmt.Errorf("no value to store")
	}
	col := db.client.Database(db.dbName).Collection(types.CoOffers)

	filter := bson.D{
		{ Key: types.FiOfferCreator, Value: Offer.Creator.String() },
		{ Key: types.FiOfferNft, Value: Offer.Nft.String() },
		{ Key: types.FiOfferTokenId, Value: Offer.TokenId.String() },
	}

	if _, err := col.ReplaceOne(context.Background(), filter, Offer); err != nil {
		log.Errorf("can not update Offer; %s", err)
		return err
	}
	return nil
}

func (db *MongoDbBridge) RemoveOffer(creator common.Address, nft common.Address, tokenId hexutil.Big) error {
	col := db.client.Database(db.dbName).Collection(types.CoOffers)

	filter := bson.D{
		{ Key: types.FiOfferCreator, Value: creator.String() },
		{ Key: types.FiOfferNft, Value: nft.String() },
		{ Key: types.FiOfferTokenId, Value: tokenId.String() },
	}

	if _, err := col.DeleteOne(context.Background(), filter); err != nil {
		log.Errorf("can not update Offer; %s", err)
		return err
	}
	return nil
}

func (db *MongoDbBridge) ListOffers(nft *common.Address, tokenId *hexutil.Big, creator *common.Address, cursor types.Cursor, count int, backward bool) (out *types.OfferList, err error) {
	filter := bson.D{}
	if nft != nil {
		filter = append(filter, primitive.E{Key: types.FiOfferNft, Value: nft.String() })
	}
	if tokenId != nil {
		filter = append(filter, primitive.E{Key: types.FiOfferTokenId, Value: tokenId.String() })
	}
	if creator != nil {
		filter = append(filter, primitive.E{Key: types.FiOfferCreator, Value: creator.String() })
	}
	return db.listOffers(&filter, cursor, count, backward)
}

func (db *MongoDbBridge) listOffers(filter *bson.D, cursor types.Cursor, count int, backward bool) (out *types.OfferList, err error) {
	var list types.OfferList
	col := db.client.Database(db.dbName).Collection(types.CoOffers)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, backward)
	if err != nil {
		log.Errorf("error loading Offers list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			log.Errorf("error closing Offers list cursor; %s", err.Error())
		}
	}()

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.Offer
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the Offer in list; %s", err.Error())
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
