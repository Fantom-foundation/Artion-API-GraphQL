// Package db provides access to the persistent storage.
package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IndexDefinitionTokens provides a list of indexes expected to exist on tokens' collection.
func IndexDefinitionTokens() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixContractToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixContractToken}}

	ixOrdinal := "ix_ordinal"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "index", Value: -1}}, Options: &options.IndexOptions{Name: &ixOrdinal}}
	return ix
}

// IndexDefinitionCollections provides list of indexes expected on collections.
func IndexDefinitionCollections() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 1)

	ixType := "ix_type"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "type", Value: 1}}, Options: &options.IndexOptions{Name: &ixType}}
	return ix
}

// IndexDefinitionListings provides list of indexes expected on listings.
func IndexDefinitionListings() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixToken}}

	ixOwner := "ix_owner"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "owner", Value: 1}}, Options: &options.IndexOptions{Name: &ixOwner}}
	return ix
}

// IndexDefinitionOwnership provides a list of indexes expected to exist on tokens' ownership records.
func IndexDefinitionOwnership() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixContractToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixContractToken}}

	ixOwner := "ix_owner"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "owner", Value: 1}}, Options: &options.IndexOptions{Name: &ixOwner}}
	return ix
}

// IndexDefinitionUsers provides a list of indexes expected to exist on users' collection.
func IndexDefinitionUsers() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixUser := "ix_username"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "username", Value: 1}}, Options: &options.IndexOptions{Name: &ixUser}}

	ixEmail := "ix_email"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "email", Value: 1}}, Options: &options.IndexOptions{Name: &ixEmail}}
	return ix
}

// IndexDefinitionOffers provides list of indexes expected on listings.
func IndexDefinitionOffers() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixToken}}

	ixOfferedBy := "ix_offered_by"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "proposer", Value: 1}}, Options: &options.IndexOptions{Name: &ixOfferedBy}}

	ixCreated := "ix_created"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "created", Value: 1}}, Options: &options.IndexOptions{Name: &ixCreated}}
	return ix
}

// IndexDefinitionAuctions provides a list of indexes expected to exist on auctions' collection.
func IndexDefinitionAuctions() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixContractToken := "ix_contract_token"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "contract", Value: 1}, {Key: "token", Value: 1}}, Options: &options.IndexOptions{Name: &ixContractToken}}

	ixOwner := "ix_owner"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "owner", Value: 1}}, Options: &options.IndexOptions{Name: &ixOwner}}

	ixOrdinal := "ix_ordinal"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "index", Value: -1}}, Options: &options.IndexOptions{Name: &ixOrdinal}}
	return ix
}
