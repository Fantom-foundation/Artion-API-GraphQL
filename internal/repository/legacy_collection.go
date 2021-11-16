package repository

import (
	"artion-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func (p *Proxy) GetLegacyCollection(address common.Address) (*types.LegacyCollection, error) {
	return p.shared.GetLegacyCollection(address)
}

func (p *Proxy) ListLegacyCollections(search *string, mintableBy *common.Address, cursor types.Cursor, count int, backward bool) (out *types.LegacyCollectionList, err error) {
	return p.shared.ListLegacyCollections(search, mintableBy, cursor, count, backward)
}

func (p *Proxy) UploadCollectionApplication(app types.CollectionApplication, image types.Image, owner common.Address) (err error) {
	imageCid, err := p.pinner.PinFile("collection-image-"+app.Contract.String(), image.Data)
	if err != nil {
		return fmt.Errorf("uploading collection image failed; %s", err)
	}
	collection := app.ToCollection(imageCid, &owner)
	return p.shared.InsertLegacyCollection(collection)
}
