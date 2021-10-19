package types

// Category is set of token collections.
type Category struct {
	Id     int32  `bson:"_id"`
	Name   string  `bson:"name"`
	Icon   *string  `bson:"icon"`
}
