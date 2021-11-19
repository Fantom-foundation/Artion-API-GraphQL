package types

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// no properties
	data := "{\"name\":\"tst\",\"image\":\"https://artion2.mypinata.cloud/abc\",\"description\":\"aaa\"}"
	meta, err := DecodeJsonMetadata([]byte(data))
	g.Expect(err).To(gomega.BeNil())
	g.Expect(meta.Name).To(gomega.Equal("tst"))

	// empty properties
	data = "{\"name\":\"tst\",\"image\":\"https://artion2.mypinata.cloud/abc\",\"description\":\"aaa\",\"properties\":{}}"
	meta, err = DecodeJsonMetadata([]byte(data))
	g.Expect(err).To(gomega.BeNil())
	g.Expect(meta.Name).To(gomega.Equal("tst"))

	data = "{\"name\":\"tst\",\"image\":\"https://artion2.mypinata.cloud/abc\",\"description\":\"aaa\",\"properties\":{\"symbol\":\"tst\",\"address\":\"0x83a6524be9213b1ce36bcc0dcefb5eb51d87ad10\",\"royalty\":\"0\",\"recipient\":\"0x83a6524be9213b1ce36bcc0dcefb5eb51d87ad10\",\"IP_Rights\":\"https://artion2.mypinata.cloud/def\",\"createdAt\":\"12:04:54 GMT+0000 (Coordinated Universal Time)\",\"collection\":\"Fantom Powered NFT Collection\"}}"
	meta, err = DecodeJsonMetadata([]byte(data))
	g.Expect(err).To(gomega.BeNil())
	g.Expect(meta.Name).To(gomega.Equal("tst"))
	g.Expect(*meta.Properties.Royalty).To(gomega.Equal("0"))
	g.Expect(*meta.Properties.Symbol).To(gomega.Equal("tst"))
	g.Expect(*meta.Properties.IpRights).To(gomega.Equal("https://artion2.mypinata.cloud/def"))

	data = "{\"name\":\"tst\",\"image\":\"https://artion2.mypinata.cloud/abc\",\"description\":\"aaa\",\"properties\":{\"symbol\":\"tst\",\"address\":\"\",\"royalty\":\"0\",\"IP_Rights\":\"\",\"collection\":\"Collection\"}}"
	meta, err = DecodeJsonMetadata([]byte(data))
	g.Expect(err).To(gomega.BeNil())
	g.Expect(meta.Name).To(gomega.Equal("tst"))
	g.Expect(*meta.Properties.Collection).To(gomega.Equal("Collection"))
}
