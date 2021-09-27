package uri

import "artion-api-graphql/internal/types"

type Downloader interface {

	GetJsonMetadata(uri string) (*types.JsonMetadata, error)

	GetImage(uri string) ([]byte, error)

}
