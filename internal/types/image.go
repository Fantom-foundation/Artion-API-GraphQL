package types

import (
	"fmt"
	svg "github.com/h2non/go-is-svg"
	"net/http"
	"strings"
)

// Image represents image of NFT downloaded from specified URI
type Image struct {
	Data []byte
	Type ImageType
}

type ImageType int8

const (
	ImageTypeUnknown ImageType = iota
	ImageTypeSvg
	ImageTypeGif
	ImageTypeJpeg
	ImageTypePng
	ImageTypeWebp
	ImageTypeMp4
)

func (i ImageType) Mimetype() string {
	switch i {
	case ImageTypeSvg: return "image/svg+xml"
	case ImageTypeGif: return "image/gif"
	case ImageTypeJpeg: return "image/jpeg"
	case ImageTypePng: return "image/png"
	case ImageTypeWebp: return "image/webp"
	case ImageTypeMp4: return "video/mp4"
	}
	return ""
}

func (i ImageType) Extension() string {
	switch i {
	case ImageTypeSvg: return ".svg"
	case ImageTypeGif: return ".gif"
	case ImageTypeJpeg: return ".jpg"
	case ImageTypePng: return".png"
	case ImageTypeWebp: return ".webp"
	case ImageTypeMp4: return ".mp4"
	}
	return ""
}

func ImageTypeFromMimetype(data []byte) (ImageType, error) {
	mimetype := http.DetectContentType(data)
	switch mimetype {
	case "image/svg+xml": return ImageTypeSvg, nil
	case "image/gif": return ImageTypeGif, nil
	case "image/jpeg": return ImageTypeJpeg, nil
	case "image/png": return ImageTypePng, nil
	case "image/webp": return ImageTypeWebp, nil
	case "video/mp4": return ImageTypeMp4, nil
	}
	if strings.HasPrefix(mimetype, "text/xml") || strings.HasPrefix(mimetype, "text/plain") {
		if svg.Is(data) {
			return ImageTypeSvg, nil
		}
	}
	return ImageTypeUnknown, fmt.Errorf("unrecognized image type %s", mimetype)
}

func ImageTypeFromExtension(uri string) (mimetype ImageType) {
	uri = strings.ToLower(uri)
	if strings.HasSuffix(uri, ".svg") {
		return ImageTypeSvg
	}
	if strings.HasSuffix(uri, ".gif") {
		return ImageTypeGif
	}
	if strings.HasSuffix(uri, ".jpg") || strings.HasSuffix(uri, ".jpeg") {
		return ImageTypeJpeg
	}
	if strings.HasSuffix(uri, ".png") {
		return ImageTypePng
	}
	if strings.HasSuffix(uri, ".webp") {
		return ImageTypeWebp
	}
	if strings.HasSuffix(uri, ".mp4") {
		return ImageTypeMp4
	}
	return ImageTypeUnknown
}
