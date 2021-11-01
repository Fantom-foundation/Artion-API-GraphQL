package types

import "encoding/json"

// JsonMetadata describes a token as defined in ERC-721/ERC-1155 Metadata JSON Schema.
type JsonMetadata struct {
	Name        string  `json:"name"`        // Identifies the asset to which this token represents
	Description *string `json:"description"` // Describes the asset to which this token represents
	Image       *string `json:"image"`       // A URI pointing to a resource with mime type image/* representing the asset to which this token represents.
	Decimals    *int    `json:"decimals"`    // ERC-1155 only: The number of decimal places that the token amount should display
	Properties  JsonMetadataProps `json:"properties"`
}

type JsonMetadataProps struct {
	Symbol     *string `json:"symbol"`
	Address    string  `json:"address"` // minter
	Royalty    *string `json:"royalty"` // percents
	Recipient  string  `json:"recipient"` // royalty recipient
	IpRights   *string `json:"IP_Rights"`
	CreatedAt  string  `json:"createdAt"`
	Collection *string `json:"collection"`
}

// DecodeJsonMetadata parses the NFT token Metadata JSON.
func DecodeJsonMetadata(data []byte) (*JsonMetadata, error) {
	var out JsonMetadata
	err := json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
