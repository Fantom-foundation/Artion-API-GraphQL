package repository

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"github.com/disintegration/imaging"
)

const thumbnailMaxHeight = 500
const thumbnailMaxWidth = 500

// createThumbnail resize the Image
func createThumbnail(input types.Image) (output types.Image, err error) {
	if input.Type == types.ImageTypeSvg {
		return input, nil // skip SVG thumbnailing
	}

	reader := bytes.NewReader(input.Data)
	var writer bytes.Buffer

	img, err := imaging.Decode(reader, imaging.AutoOrientation(true))
	if err != nil {
		return types.Image{}, err
	}

	small := imaging.Fit(img, thumbnailMaxWidth, thumbnailMaxHeight, imaging.Linear)

	if input.Type == types.ImageTypeJpeg {
		err = imaging.Encode(&writer, small, imaging.JPEG, imaging.JPEGQuality(80))
	} else  {
		err = imaging.Encode(&writer, small, imaging.PNG)
		input.Type = types.ImageTypePng
	}
	if err != nil {
		return types.Image{}, err
	}
	return types.Image{
		Data: writer.Bytes(),
		Type: input.Type,
	}, nil
}
