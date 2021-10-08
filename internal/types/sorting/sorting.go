package sorting

// Sorting collection-specific enums defines sorting possibilities for individual collections.
// Because sorting defines what pagination cursor needs to contain, the Sorting enum
// also ensure cursor encoding/decoding.
type Sorting interface {
	// SortedFieldBson provides BSON name of currently sorting column (empty string for sorting by ordinal only)
	SortedFieldBson() string
	// OrdinalFieldBson provides BSON name of collection-specific ordinal field - must be unique in collection
	OrdinalFieldBson() string
}
