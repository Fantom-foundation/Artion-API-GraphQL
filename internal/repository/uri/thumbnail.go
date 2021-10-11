package uri

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"github.com/disintegration/imaging"
)

const thumbnailMaxHeight = 500
const thumbnailMaxWidth = 500

func createThumbnail(input types.Image) (output types.Image, err error) {
	if input.Mimetype == "image/svg+xml" {
		return input, nil // skip SVG thumbnailing
	}

	reader := bytes.NewReader(input.Data)
	var writer bytes.Buffer

	img, err := imaging.Decode(reader, imaging.AutoOrientation(true))
	if err != nil {
		return types.Image{}, err
	}

	small := imaging.Fit(img, thumbnailMaxWidth, thumbnailMaxHeight, imaging.Linear)

	if input.Mimetype == "image/jpeg" {
		err = imaging.Encode(&writer, small, imaging.JPEG, imaging.JPEGQuality(80))
	} else  {
		err = imaging.Encode(&writer, small, imaging.PNG)
		input.Mimetype = "image/png"
	}
	if err != nil {
		return types.Image{}, err
	}
	return types.Image{
		Data: writer.Bytes(),
		Mimetype: input.Mimetype,
	}, nil
}
