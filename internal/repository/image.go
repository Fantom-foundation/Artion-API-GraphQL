package repository

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

const thumbnailMaxHeight = 500
const thumbnailMaxWidth = 500

// createThumbnail resize the Image
func createThumbnail(input types.Image) (output types.Image, err error) {
	if input.Type == types.ImageTypeSvg || len(input.Data) == 0 {
		return input, nil // skip thumbnailing of SVG and empty files
	}
	if input.Type == types.ImageTypeMp4 {
		input, err = createVideoThumbnail(input)
		if err != nil {
			return types.Image{}, err
		}
	}
	return createImageThumbnail(input)
}

func createImageThumbnail(input types.Image) (output types.Image, err error) {
	reader := bytes.NewReader(input.Data)
	var writer bytes.Buffer

	img, err := imaging.Decode(reader, imaging.AutoOrientation(true))
	if err != nil {
		return types.Image{}, fmt.Errorf("decoding failed (type %d); %s", input.Type, err)
	}

	small := imaging.Fit(img, thumbnailMaxWidth, thumbnailMaxHeight, imaging.Linear)

	if input.Type == types.ImageTypeJpeg {
		err = imaging.Encode(&writer, small, imaging.JPEG, imaging.JPEGQuality(80))
	} else {
		err = imaging.Encode(&writer, small, imaging.PNG) // also for GIFs
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

func createVideoThumbnail(input types.Image) (output types.Image, err error) {
	inputReader := bytes.NewReader(input.Data)
	writer := bytes.NewBuffer(nil)
	frameNum := 3
	err = ffmpeg.
		Input("pipe:").
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithInput(inputReader).
		WithOutput(writer).
		Run()
	if err != nil {
		return types.Image{}, err
	}
	return types.Image{
		Data: writer.Bytes(),
		Type: types.ImageTypeJpeg,
	}, nil
}
