package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

func (p *Proxy) GetUser(address common.Address) (*types.User, error) {
	key := "GetUser-" + address.String()
	user, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.shared.GetUser(address)
	})
	return user.(*types.User), err
}

func (p *Proxy) ListUserUsers(search string, cursor types.Cursor, count int, backward bool) (out *types.UserList, err error) {
	return p.shared.ListUserUsers(search, cursor, count, backward)
}

func (p *Proxy) StoreUser(User *types.User) error {
	return p.shared.StoreUser(User)
}

func (p *Proxy) UploadUserAvatar(address common.Address, image types.Image) error {
	cid, err := p.pinner.PinFile("user-avatar-"+address.String()+image.Type.Extension(), image.Data)
	if err != nil {
		return err
	}
	return p.shared.SetUserAvatar(address, cid)
}

func (p *Proxy) UploadUserBanner(address common.Address, image types.Image) error {
	cid, err := p.pinner.PinFile("user-banner-"+address.String()+image.Type.Extension(), image.Data)
	if err != nil {
		return err
	}
	return p.shared.SetUserBanner(address, cid)
}
