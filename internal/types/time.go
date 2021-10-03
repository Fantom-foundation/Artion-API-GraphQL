package types

import (
	"encoding/json"
	"fmt"
	"time"
)

// Processed on the basis of a package "github.com/graph-gophers/definition-go"

// Time is a custom GraphQL type to represent an instant in time. It has to be added to a schema
// via "scalar Time" since it is not a predeclared GraphQL type like "ID".
type Time time.Time

// ImplementsGraphQLType maps this custom Go type
// to the definition scalar type in the schema.
func (Time) ImplementsGraphQLType(name string) bool {
	return name == "Time"
}

// UnmarshalGraphQL is a custom unmarshaler for Time
//
// This function will be called whenever you use the
// time scalar as an input
func (t *Time) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case time.Time:
		*t = Time(input)
		return nil
	case string:
		var err error
		var parsed time.Time
		parsed, err = time.Parse(time.RFC3339, input)
		if err == nil {
			*t = Time(parsed)
		}
		return err
	case int32:
		*t = Time(time.Unix(int64(input), 0))
		return nil
	case int64:
		*t = Time(time.Unix(input, 0))
		return nil
	case float64:
		*t = Time(time.Unix(int64(input), 0))
		return nil
	default:
		return fmt.Errorf("wrong type for Time: %T", input)
	}
}

// MarshalJSON is a custom marshaler for Time
//
// This function will be called whenever you
// query for fields that use the Time type
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t))
}
