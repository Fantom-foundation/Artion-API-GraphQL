package types

// JsonMetadata describes a token as defined in ERC-721/ERC-1155 Metadata JSON Schema.
type JsonMetadata struct {
	Name        string  `json:"name"`        // Identifies the asset to which this token represents
	Description string  `json:"description"` // Describes the asset to which this token represents
	Image       *string `json:"image"`       // A URI pointing to a resource with mime type image/* representing the asset to which this token represents.
	Decimals    *int    `json:"decimals"`    // ERC-1155 only: The number of decimal places that the token amount should display
}
