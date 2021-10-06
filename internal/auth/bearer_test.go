package auth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/onsi/gomega"
	"testing"
)

func TestBearer(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	address := common.HexToAddress("0x83A6524Be9213B1Ce36bCc0DCEfb5eb51D87aD10")
	secret := hexutil.MustDecode("0x0123456789")

	token, err := generateBearer(&address, secret)
	g.Expect(err).To(gomega.BeNil())

	outAddress, err := verifyBearer(token, secret)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(*outAddress).To(gomega.Equal(address))
}
