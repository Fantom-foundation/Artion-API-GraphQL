package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	// CoUsers is the name of database collection.
	CoUsers = "Users"
)

// User represents user account/profile.
type User struct {
	Address   common.Address
	Username  string
	Bio       string
	Email     string
	Avatar    string
}

type UserBson struct {
	Id           string    `bson:"_id"`
	Username     string    `bson:"username"`
	Bio          string    `bson:"bio"`
	Email        string    `bson:"email"`
	Avatar       string    `bson:"avatar"`
}

// MarshalBSON prepares data to be stored in MongoDB.
func (o *User) MarshalBSON() ([]byte, error) {
	row := UserBson {
		Id:           o.Address.String(),
		Username:     o.Username,
		Bio:          o.Bio,
		Email:        o.Email,
		Avatar:       o.Avatar,
	}
	return bson.Marshal(row)
}

// UnmarshalBSON parses data from MongoDB.
func (o *User) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode Offer; %s", err.Error())
		}
	}()

	// try to decode the BSON data
	var row UserBson
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	o.Address = common.HexToAddress(row.Id)
	o.Username = row.Username
	o.Bio = row.Bio
	o.Email = row.Email
	o.Avatar = row.Avatar
	return nil
}
