package repository

import (
	"artion-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

func (p *Proxy) GetLegacyCollection(address common.Address) (*types.LegacyCollection, error) {
	key := "LCol-" + address.String()
	user, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.shared.GetLegacyCollection(address)
	})
	return user.(*types.LegacyCollection), err
}

func (p *Proxy) ListLegacyCollections(collectionFilter types.CollectionFilter, cursor types.Cursor, count int, backward bool) (out *types.LegacyCollectionList, err error) {
	return p.shared.ListLegacyCollections(collectionFilter, cursor, count, backward)
}

func (p *Proxy) UploadCollectionApplication(app types.CollectionApplication, image types.Image, owner common.Address) (err error) {
	imageCid, err := p.pinner.PinFile("collection-image-"+app.Contract.String(), image.Data)
	if err != nil {
		return fmt.Errorf("uploading collection image failed; %s", err)
	}
	collection := app.ToCollection(imageCid, &owner)
	return p.shared.InsertLegacyCollection(collection)
}

// MustCollectionName provides a name of an Artion ERC721 and/or ERC1155 token,
// or an empty string, if the name is not available.
func (p *Proxy) MustCollectionName(adr *common.Address) string {
	c, err := p.shared.GetLegacyCollection(*adr)
	if err != nil {
		return adr.String()
	}
	if c.Name == "" {
		return adr.String()
	}
	return c.Name
}

func (p *Proxy) ApproveCollection(address common.Address) error {
	return p.shared.ApproveCollection(address)
}

func (p *Proxy) DeclineCollection(address common.Address) error {
	return p.shared.DeclineCollection(address)
}

func (p *Proxy) BanCollection(address common.Address) error {
	return p.shared.BanCollection(address)
}

func (p *Proxy) UnbanCollection(address common.Address) error {
	return p.shared.UnbanCollection(address)
}

func (p *Proxy) ListCollectionsWithAppropriateUpdate(after time.Time, maxAmount int64) (out []*types.LegacyCollection, err error) {
	return p.shared.ListCollectionsWithAppropriateUpdate(after, maxAmount)
}
