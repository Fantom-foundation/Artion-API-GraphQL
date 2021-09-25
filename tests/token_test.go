package tests

import (
	"artion-api-graphql/internal/types"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/onsi/gomega"
	"math/big"
	"testing"
)

func TestTokenIdGenerator(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	tok1 := new(types.Token)
	tok1.Nft = common.HexToAddress("0xf41270836df4db1d28f7fd0935270e3a603e78cc")
	tok1.TokenId = (hexutil.Big)(*big.NewInt(9292))
	tok1.GenerateId()
	g.Expect(hex.EncodeToString(tok1.Id)).To(gomega.Equal("71e825e157ec35c117c8c61a45a4564f"))
}
