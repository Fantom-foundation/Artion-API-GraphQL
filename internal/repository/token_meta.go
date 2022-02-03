package repository

import (
	"artion-api-graphql/internal/types"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (p *Proxy) GetTokenJsonMetadata(uri string) (*types.JsonMetadata, error) {
	var key strings.Builder
	key.WriteString("GetTokenJsonMetadata")
	key.WriteString(uri)

	jsonMetadata, err, _ := p.callGroup.Do(key.String(), func() (interface{}, error) {
		data, _, err := p.getFileFromUri(uri)
		if err != nil {
			return nil, fmt.Errorf("unable to download json; %s", err)
		}
		jsonMeta, err := types.DecodeJsonMetadata(data)
		if err != nil {
			return nil, fmt.Errorf("unable to decode json; %s", err)
		}
		return jsonMeta, nil
	})
	if err != nil {
		return nil, err
	}
	return jsonMetadata.(*types.JsonMetadata), err
}

// GetImage downloads an image expected on the given URI.
func (p *Proxy) GetImage(imgUri string) (*types.Image, error) {
	var key strings.Builder
	key.WriteString("GetImage")
	key.WriteString(imgUri)

	data, err, _ := p.callGroup.Do(key.String(), func() (interface{}, error) {
		return p.getImageFromUri(imgUri)
	})
	if err != nil {
		log.Errorf("image can not be loaded from %s; %s", imgUri, err.Error())
		return nil, err
	}
	if data == nil {
		log.Errorf("image not found at %s", imgUri)
		return nil, fmt.Errorf("image not found at given URI")
	}
	return data.(*types.Image), err
}

// GetImageThumbnail generates a thumbnail for an image expected to be downloadable from the given URI.
func (p *Proxy) GetImageThumbnail(imgUri string) (*types.Image, error) {
	var key strings.Builder
	key.WriteString("GetImageThumbnail")
	key.WriteString(imgUri)

	data, err, _ := p.callGroup.Do(key.String(), func() (interface{}, error) {
		image, err := p.GetImage(imgUri)
		if err != nil {
			return nil, fmt.Errorf("image loading failed for %s; %s", imgUri, err)
		}
		if nil == image {
			return nil, fmt.Errorf("image %s not found", imgUri)
		}

		log.Infof("loaded %s of type %s", imgUri, image.Type.Mimetype())
		thumb, err := createThumbnail(*image)
		if err != nil {
			return nil, fmt.Errorf("thumbnail creation failed; %s", err)
		}
		return &thumb, nil
	})
	if data == nil {
		return nil, err
	}
	return data.(*types.Image), err
}

func (p *Proxy) UploadTokenData(metadata types.JsonMetadata, image types.Image) (uri string, err error) {
	cid, err := p.pinFile("token-image", image.Data)
	if err != nil {
		return "", fmt.Errorf("uploading token image failed; %s", err)
	}
	imageUri := "https://artion.mypinata.cloud/ipfs/" + cid
	metadata.Image = &imageUri

	data, err := json.Marshal(metadata)
	if err != nil {
		return "", fmt.Errorf("marshaling json meta failed; %s", err)
	}

	cid, err = p.pinFile("token-meta", data)
	if err != nil {
		return "", fmt.Errorf("uploading token meta failed; %s", err)
	}
	return "https://artion.mypinata.cloud/ipfs/" + cid, nil
}

func (p *Proxy) pinFile(filename string, content []byte) (cid string, err error) {
	cid, err = p.pinner.PinFile(filename, content)
	if err == nil {
		err = p.filecache.PushIpfsFile(cid, content)
	}
	return
}

// getImageFromUri downloads image from given URI and detect its mimetype
func (p *Proxy) getImageFromUri(uri string) (*types.Image, error) {
	data, mimetype, err := p.getFileFromUri(uri)
	if err != nil {
		return nil, fmt.Errorf("unable to download image; %s", err)
	}

	if mimetype == "" {
		mimetype = http.DetectContentType(data)
	}

	imgType := types.ImageTypeFromMimetype(mimetype)
	if imgType == types.ImageTypeUnknown {
		imgType = types.ImageTypeFromExtension(uri)
	}

	out := types.Image{
		Data: data,
		Type: imgType,
	}
	return &out, nil
}

func getCidFromIpfsUri(uri string) string {
	cid := uri[6:] // skip /ipfs/
	slashIdx := strings.Index(cid, "/")
	if slashIdx != -1 {
		cid = cid[0:slashIdx]
	}
	return cid
}

// getFileFromUri resolves the URI and download file from the URI using appropriate protocol
func (p *Proxy) getFileFromUri(uri string) (data []byte, mimetype string, err error) {
	if strings.HasPrefix(uri, "data:") {
		return p.uri.GetFromDataUri(uri)
	}

	if ipfsUri := p.uri.GetIpfsUri(uri); ipfsUri != "" {
		cachedContent := p.filecache.PullIpfsFile(getCidFromIpfsUri(ipfsUri))
		if cachedContent != nil {
			return cachedContent, "", nil
		}
		return p.uri.GetFromIpfs(ipfsUri)
	}

	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		return p.uri.GetFromHttp(uri)
	}

	return nil, "", errors.New("Unexpected URI scheme for " + uri)
}
