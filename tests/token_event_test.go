package tests

import (
	"artion-api-graphql/internal/types"
	"encoding/hex"
	"github.com/onsi/gomega"
	"testing"
)

func TestTokenEventIdGenerator(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	event := new(types.TokenEvent)

	event.GenerateId(0x1234, 0x4321, 0x98)
	g.Expect(hex.EncodeToString(event.Id[:])).To(gomega.Equal("0000000000001234000000000000432100000098"))

	event.GenerateId(0x614B51A5, 0x10AF083, 0x1234)
	g.Expect(hex.EncodeToString(event.Id[:])).To(gomega.Equal("00000000614b51a500000000010af08300001234"))

	event.GenerateId(0x614B51A5, 0x10AF083123456, 0x1)
	g.Expect(hex.EncodeToString(event.Id[:])).To(gomega.Equal("00000000614b51a500010af08312345600000001"))
}
