package types

import "strings"

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
)

func (i ImageType) Mimetype() string {
	switch i {
	case ImageTypeSvg: return "image/svg+xml"
	case ImageTypeGif: return "image/gif"
	case ImageTypeJpeg: return "image/jpeg"
	case ImageTypePng: return "image/png"
	case ImageTypeWebp: return "image/webp"
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
	}
	return ""
}

func ImageTypeFromMimetype(mimetype string) ImageType {
	switch mimetype {
	case "image/svg+xml": return ImageTypeSvg
	case "image/gif": return ImageTypeGif
	case "image/jpeg": return ImageTypeJpeg
	case "image/png": return ImageTypePng
	case "image/webp": return ImageTypeWebp
	}
	return ImageTypeUnknown
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
	return ImageTypeUnknown
}
