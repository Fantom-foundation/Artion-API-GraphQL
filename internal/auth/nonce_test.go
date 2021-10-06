package auth

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/onsi/gomega"
	"testing"
)

func TestNonce(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	secret := hexutil.MustDecode("0xABCDEF")

	nonce, err := generateNonce(secret)
	g.Expect(err).To(gomega.BeNil())

	err = verifyNonce(nonce, secret)
	g.Expect(err).To(gomega.BeNil())
}
