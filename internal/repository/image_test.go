package repository

import (
	"artion-api-graphql/internal/repository/uri"
	"artion-api-graphql/internal/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestVideoThumbnail(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	downloader := &uri.Downloader{}
	image, err := downloader.GetImage("https://artion.mypinata.cloud/ipfs/QmePhQPfwwCWzqSTpxa2CQFCWLbDwj2PATdL6AFYRw7nFc")
	g.Expect(image.Type).To(gomega.Equal(types.ImageTypeMp4))
	thumb, err := createThumbnail(*image)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(len(thumb.Data) > 10).To(gomega.BeTrue())
	g.Expect(thumb.Type).To(gomega.Equal(types.ImageTypeJpeg))

	//err = os.WriteFile("/tmp/baku-thumb.jpg", thumb.Data, 0744)
	//g.Expect(err).To(gomega.BeNil())
}

func TestJpgThumbnail(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	downloader := &uri.Downloader{}
	image, err := downloader.GetImage("https://artion.mypinata.cloud/ipfs/QmbcqNsWpQuE56xVKQ226rDM29KM7UegmueqMeXPpq5qwo/140.png")
	g.Expect(image.Type).To(gomega.Equal(types.ImageTypePng))
	thumb, err := createThumbnail(*image)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(len(thumb.Data) > 10).To(gomega.BeTrue())
	g.Expect(thumb.Type).To(gomega.Equal(types.ImageTypePng))

	//err = os.WriteFile("/tmp/140.png", thumb.Data, 0744)
	//g.Expect(err).To(gomega.BeNil())
}
