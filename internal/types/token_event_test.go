package types

import (
	"encoding/hex"
	"github.com/onsi/gomega"
	"testing"
)

func TestTokenEventIdGenerator(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	event := new(TokenEvent)

	event.GenerateId(0x1234, 0x4321, 0x98)
	g.Expect(hex.EncodeToString(event.Id[:])).To(gomega.Equal("000012349800000000004321"))

	event.GenerateId(0x614B51A5, 0x10AF083, 0x1234)
	g.Expect(hex.EncodeToString(event.Id[:])).To(gomega.Equal("614b51a534120000010af083"))

	event.GenerateId(0x614B51A5, 0x10AF083123456, 0x1)
	g.Expect(hex.EncodeToString(event.Id[:])).To(gomega.Equal("614b51a501010af083123456"))
}

