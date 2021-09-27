package uri

import "artion-api-graphql/internal/types"

// Downloader handles downloading of JSON metadata and images from given URI
// Handles various protocols like HTTP, IPFS or data encoded directly into URI.
type Downloader interface {

	// GetJsonMetadata obtains token metadata from JSON at given URI
	GetJsonMetadata(uri string) (*types.JsonMetadata, error)

	// GetImage obtains image (including mimetype) from given URI
	GetImage(uri string) (image *types.Image, err error)

}
