package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// FOR TESTING ONLY! remove before production use!
func (rs *RootResolver) PushTestingData() (*string, error) {

	tok := types.Token{
		Nft: common.HexToAddress("0xf41270836dF4Db1D28F7fd0935270e3A603e78cC"),
		TokenId: hexutil.Big(*big.NewInt(9292)),
		Uri: "ipfs://QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi",
	}
	tok.GenerateId()
	err := repository.R().StoreToken(&tok)
	if err != nil {
		log.Errorf("error in storing token", err)
	}

	tok = types.Token{
		Nft: common.HexToAddress("0x61af4d29f672e27a097291f72fc571304bc93521"),
		TokenId: hexutil.Big(*big.NewInt(3224)),
		Uri: "https://artion1.mypinata.cloud/ipfs/QmcAtMen7niz53eHpnehWJvYPVEMkxaT1y8mtkPjz5rAwf",
	}
	tok.GenerateId()
	err = repository.R().StoreToken(&tok)
	if err != nil {
		log.Errorf("error in storing token", err)
	}

	listing := types.Listing{
		Owner: common.HexToAddress("0xabcd12345678927a097291f72fc571304321abcd"),
		Nft: common.HexToAddress("0x61af4d29f672e27a097291f72fc571304bc93521"),
		TokenId: hexutil.Big(*big.NewInt(3224)),
		Quantity: hexutil.Big(*big.NewInt(2)),
		PayToken: common.HexToAddress("0x432112345678927a097291f72fc5713043218888"),
		PricePerItem: hexutil.Big(*big.NewInt(15876900)),
		StartTime: types.Time(time.Now()),
	}
	listing.GenerateId()
	err = repository.R().StoreListing(&listing)
	if err != nil {
		log.Errorf("error in storing listing", err)
	}

	offer := types.Offer{
		Creator: common.HexToAddress("0x111112345678927a097291f72fc5713043211111"),
		Nft: common.HexToAddress("0x61af4d29f672e27a097291f72fc571304bc93521"),
		TokenId: hexutil.Big(*big.NewInt(3224)),
		Quantity: hexutil.Big(*big.NewInt(2)),
		PayToken: common.HexToAddress("0x432112345678927a097291f72fc5713043218888"),
		PricePerItem: hexutil.Big(*big.NewInt(15876900)),
		StartTime: types.Time(time.Now()),
		Deadline: types.Time(time.Now().Add(time.Hour * 48)),
	}
	offer.GenerateId()
	err = repository.R().StoreOffer(&offer)
	if err != nil {
		log.Errorf("error in storing offer", err)
	}

	out := "Loaded OK"
	return &out, nil
}
