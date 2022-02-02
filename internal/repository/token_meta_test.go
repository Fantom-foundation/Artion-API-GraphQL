package repository

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestCidFromIpfsUri(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	cid := getCidFromIpfsUri("/ipfs/QmXNdxJiyPpM2NTLC3panhpFde16KGNHZ4PUPVXz3W6DBR/150")
	g.Expect(cid).To(gomega.Equal("QmXNdxJiyPpM2NTLC3panhpFde16KGNHZ4PUPVXz3W6DBR"))
}
