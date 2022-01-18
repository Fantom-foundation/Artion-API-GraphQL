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

	// fiOfferProposer is the name of the DB column of the offer creator address.
	fiOfferOwners = "owners"

	// fiOfferClosed is the name of the DB column of the offer having been closed date/time.
	fiOfferClosed = "closed"

	// fiOfferDeadline is the name of the DB column of the offer expiring date/time.
	fiOfferDeadline = "deadline"

	// fiOfferUnifiedPrice is the name of the DB column storing price converted to USD.
	fiOfferUnifiedPrice = "uprice"
)

// GetOffer provides the token offer stored in the database, if available.
func (db *MongoDbBridge) GetOffer(contract *common.Address, tokenID *big.Int, proposer *common.Address, marketplace *common.Address) (*types.Offer, error) {
	// get the collection
	col := db.client.Database(db.dbName).Collection(coOffers)

	sr := col.FindOne(context.Background(), bson.D{{Key: fieldId, Value: types.OfferID(contract, tokenID, proposer, marketplace)}})
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
func (db *MongoDbBridge) StoreOffer(offer *types.Offer) (err error) {
	if offer == nil {
		return fmt.Errorf("no value to store")
	}

	// get receivers of the offer
	offer.Owners, err = db.GetTokenOwners(offer.Contract, offer.TokenId)
	if err != nil {
		return fmt.Errorf("unable to load owners for offer %s/%s; %s", offer.Contract.String(), offer.TokenId.String(), err)
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

// UpdateOffersOwners updates token owners in offers for given token
func (db *MongoDbBridge) UpdateOffersOwners(contract common.Address, tokenID hexutil.Big) {
	owners, err := db.GetTokenOwners(contract, tokenID)
	if err != nil {
		log.Errorf("failed to load offers owners for %s/%s; %s",
			contract.String(), tokenID.String(), err)
	}
	col := db.client.Database(db.dbName).Collection(coOffers)
	_, err = col.UpdateMany(context.Background(), bson.D{
		{Key: fiOfferContract, Value: contract},
		{Key: fiOfferTokenId, Value: tokenID},
	}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: fiOfferOwners, Value: owners},
		}},
	})
	if err != nil {
		log.Errorf("failed to update offers owners for %s/%s; %s",
			contract.String(), tokenID.String(), err)
	}
}

// OpenOfferUntil provides the latest active offer date/time if any.
func (db *MongoDbBridge) OpenOfferUntil(contract *common.Address, tokenID *big.Int) *types.Time {
	var row struct {
		Until types.Time `bson:"val"`
	}

	col := db.client.Database(db.dbName).Collection(coOffers)
	err := db.AggregateSingle(col, &mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: fiOfferContract, Value: *contract},
				{Key: fiOfferTokenId, Value: hexutil.Big(*tokenID)},
				{Key: fiOfferClosed, Value: bson.D{{Key: "$type", Value: 10}}},
			}},
		},
		bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: nil},
				{Key: "val", Value: bson.D{{Key: "$max", Value: "$deadline"}}},
			}},
		},
	}, &row)

	if err != nil {
		// no offer at all?
		if err == mongo.ErrNoDocuments {
			log.Debugf("no open offer available for %s/%s", contract.String(), (*hexutil.Big)(tokenID).String())
			return nil
		}
		log.Criticalf("failed latest offer check of %s/%s; %s", contract.String(), (*hexutil.Big)(tokenID).String(), err.Error())
		return nil
	}

	log.Infof("open offer on %s/%s until %s", contract.String(), (*hexutil.Big)(tokenID).String(), time.Time(row.Until).Format(time.RFC1123))
	return &row.Until
}

// MaxOfferPrice obtains maximal offered price for the token and time until when is the information valid.
func (db *MongoDbBridge) MaxOfferPrice(contract *common.Address, tokenID *big.Int) (tokenPrice types.TokenPrice, valid *types.Time) {
	col := db.client.Database(db.dbName).Collection(coOffers)
	now := time.Now()

	// get minimal price starting before now
	var maxOffer types.Offer
	sr := col.FindOne(context.Background(), bson.D{
		{Key: fiOfferContract, Value: *contract},
		{Key: fiOfferTokenId, Value: hexutil.Big(*tokenID)},
		{Key: fiOfferClosed, Value: bson.D{{Key: "$type", Value: 10}}},   // not closed yet
		{Key: fiOfferDeadline, Value: bson.D{{Key: "$gte", Value: now}}}, // before deadline
	}, options.FindOne().SetSort(bson.D{{Key: fiOfferUnifiedPrice, Value: -1}}))
	if sr.Err() != nil {
		if sr.Err() != mongo.ErrNoDocuments {
			log.Errorf("error loading max offer price; %s", sr.Err())
		}
		return
	}
	err := sr.Decode(&maxOffer)
	if err != nil {
		log.Errorf("error decoding max offer price; %s", err)
		return
	}
	tokenPrice = types.TokenPrice{
		Usd:      maxOffer.UnifiedPrice,
		Amount:   maxOffer.UnitPrice,
		PayToken: maxOffer.PayToken,
	}
	valid = &maxOffer.Deadline
	return
}

func (db *MongoDbBridge) ListOffers(contract *common.Address, tokenId *hexutil.Big, creator *common.Address, owner *common.Address, activeOnly bool, cursor types.Cursor, count int, backward bool) (out *types.OfferList, err error) {
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
	if owner != nil {
		filter = append(filter, primitive.E{Key: fiOfferOwners, Value: bson.D{{Key: "$all", Value: primitive.A{owner.String()}}}})
	}
	if activeOnly {
		filter = append(filter,
			primitive.E{Key: fiOfferClosed, Value: bson.D{{Key: "$type", Value: 10}}}, // not closed yet
			primitive.E{Key: fiOfferDeadline, Value: bson.D{{Key: "$gte", Value: time.Now()}}}, // before deadline
			)
	}
	return db.listOffers(filter, cursor, count, backward)
}

func (db *MongoDbBridge) listOffers(filter bson.D, cursor types.Cursor, count int, backward bool) (out *types.OfferList, err error) {
	var list types.OfferList
	col := db.client.Database(db.dbName).Collection(coOffers)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return &list, nil // interested in TotalCount only
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.OfferSortingCreated, !backward)
	if err != nil {
		log.Errorf("error loading Offers list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer closeFindCursor("listOffers", ld)

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
