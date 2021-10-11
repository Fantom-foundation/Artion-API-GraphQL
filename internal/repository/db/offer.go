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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
	"time"
)

const (
	// coOffers is the name of database collection.
	coOffers = "offers"

	// fiOfferContract is the name of the DB column of the NFT contract address.
	fiOfferContract = "contract"

	// fiOfferTokenId is the name of the DB column of the token ID.
	fiOfferTokenId = "token"

	// fiOfferProposer is the name of the DB column of the offer creator address.
	fiOfferProposer = "proposer"

	// fiOfferClosed is the name of the DB column of the offer having been closed date/time.
	fiOfferClosed = "closed"

	// fiOfferDeadline is the name of the DB column of the offer expiring date/time.
	fiOfferDeadline = "deadline"
)

// GetOffer provides the token offer stored in the database, if available.
func (db *MongoDbBridge) GetOffer(contract *common.Address, tokenID *big.Int, proposer *common.Address) (*types.Offer, error) {
	// get the collection
	col := db.client.Database(db.dbName).Collection(coOffers)

	sr := col.FindOne(context.Background(), bson.D{{Key: fieldId, Value: types.OfferID(contract, tokenID, proposer)}})
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			log.Warningf("could not find offer %s/%s proposed by %s",
				contract.String(), (*hexutil.Big)(tokenID).String(), proposer.String())
			return nil, sr.Err()
		}

		log.Errorf("failed to lookup offer %s/%s proposed by %s; %s",
			contract.String(), (*hexutil.Big)(tokenID).String(), proposer.String(), sr.Err().Error())
		return nil, sr.Err()
	}

	// decode
	var row types.Offer
	if err := sr.Decode(&row); err != nil {
		log.Errorf("could not decode offer %s/%s proposed by %s; %s",
			contract.String(), (*hexutil.Big)(tokenID).String(), proposer.String(), sr.Err().Error())
		return nil, sr.Err()
	}
	return &row, nil
}

// StoreOffer adds the provided offer into the database.
func (db *MongoDbBridge) StoreOffer(offer *types.Offer) error {
	if offer == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coOffers)

	// try to do the insert
	id := offer.ID()
	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: fieldId, Value: id}},
		bson.D{
			{Key: "$set", Value: offer},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fieldId, Value: id},
			}},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		log.Errorf("can not add offer %s/%s of %s; %s", offer.Contract.String(), offer.TokenId.String(), offer.ProposedBy.String(), err)
		return err
	}
	return nil
}

// HasOpenOffers checks if the given NFT has standing offer(s).
func (db *MongoDbBridge) HasOpenOffers(contract *common.Address, tokenID *big.Int) bool {
	col := db.client.Database(db.dbName).Collection(coOffers)

	// check for count of any un-closed offers with future deadline; we need only 1
	count, err := col.CountDocuments(context.Background(), bson.D{
		{Key: fiOfferContract, Value: *contract},
		{Key: fiOfferTokenId, Value: hexutil.Big(*tokenID)},
		{Key: fiOfferClosed, Value: bson.D{{Key: "$type", Value: 10}}},
		{Key: fiOfferDeadline, Value: bson.D{{Key: "$gt", Value: types.Time(time.Now().UTC())}}},
	}, options.Count().SetLimit(1))

	if err != nil {
		log.Errorf("can not count offers; %s", err.Error())
		return false
	}

	return count > 0
}

func (db *MongoDbBridge) ListOffers(contract *common.Address, tokenId *hexutil.Big, creator *common.Address, cursor types.Cursor, count int, backward bool) (out *types.OfferList, err error) {
	filter := bson.D{}
	if contract != nil {
		filter = append(filter, primitive.E{Key: fiOfferContract, Value: contract.String()})
	}
	if tokenId != nil {
		filter = append(filter, primitive.E{Key: fiOfferTokenId, Value: tokenId.String()})
	}
	if creator != nil {
		filter = append(filter, primitive.E{Key: fiOfferProposer, Value: creator.String()})
	}
	return db.listOffers(&filter, cursor, count, backward)
}

func (db *MongoDbBridge) listOffers(filter *bson.D, cursor types.Cursor, count int, backward bool) (out *types.OfferList, err error) {
	var list types.OfferList
	col := db.client.Database(db.dbName).Collection(coOffers)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.OfferSortingNone, backward)
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
