package repository

import (
	"artion-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

func (p *Proxy) GetLegacyCollection(address common.Address) (*types.LegacyCollection, error) {
	key := "lcol_-" + address.String()
	user, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.cache.GetLegacyCollection(address, p.shared.GetLegacyCollection)
	})
	return user.(*types.LegacyCollection), err
}

func (p *Proxy) ListLegacyCollections(collectionFilter types.CollectionFilter, cursor types.Cursor, count int, backward bool) (out *types.LegacyCollectionList, err error) {
	if collectionFilter.IsUsed() || cursor != "" || backward {
		return p.shared.ListLegacyCollections(collectionFilter, cursor, count, backward)
	}

	// cache result for the default (mostly used) params
	list, err, _ := p.callGroup.Do("lcols", func() (interface{}, error) {
		return p.cache.GetLegacyCollectionList(func() (*types.LegacyCollectionList, error) {
			return p.shared.ListLegacyCollections(types.CollectionFilter{}, "", 5000, false)
		})
	})

	colList := list.(*types.LegacyCollectionList)
	if colList != nil && len(colList.Collection) > count {
		newColList := *colList // clone callGroup output before modification
		newColList.Collection = newColList.Collection[:count]
		newColList.HasNext = true
		return &newColList, err
	}
	return colList, err
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

func (p *Proxy) ApproveCollection(address common.Address) (err error) {
	err = p.shared.ApproveCollection(address)
	p.cache.InvalidateLegacyCollection(address)
	return err
}

func (p *Proxy) DeclineCollection(address common.Address) (err error) {
	err = p.shared.DeclineCollection(address)
	p.cache.InvalidateLegacyCollection(address)
	return err
}

func (p *Proxy) BanCollection(address common.Address) (err error) {
	err = p.shared.BanCollection(address)
	p.cache.InvalidateLegacyCollection(address)
	return err
}

func (p *Proxy) UnbanCollection(address common.Address) (err error) {
	err = p.shared.UnbanCollection(address)
	p.cache.InvalidateLegacyCollection(address)
	return err
}

func (p *Proxy) ListCollectionsWithAppropriateUpdate(after time.Time, maxAmount int64) (out []*types.LegacyCollection, err error) {
	return p.shared.ListCollectionsWithAppropriateUpdate(after, maxAmount)
}
