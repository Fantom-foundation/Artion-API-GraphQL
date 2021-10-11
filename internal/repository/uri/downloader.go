// Package uri provides tools for downloading data from URI - especially
// NFT tokens JSON Metadata and related images using IPFS or HTTP.
package uri

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/types"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	ipfsapi "github.com/ipfs/go-ipfs-api"
	"io"
	"net/http"
	"strings"
	"time"
)

// ipfsRequestTimeout represents the timeout applied to IPFS requests.
const ipfsRequestTimeout = 5 * time.Second

type Downloader struct {
	ipfsShell *ipfsapi.Shell
	skipHttpGateways bool
}

// New provides new Downloader instance.
func New(cfg *config.Config) *Downloader {
	d := &Downloader{
		ipfsShell: ipfsapi.NewShell(cfg.Ipfs.Url),
		skipHttpGateways: cfg.Ipfs.SkipHttpGateways,
	}
	d.ipfsShell.SetTimeout(ipfsRequestTimeout)
	return d
}

// GetJsonMetadata download and parse NFT token Metadata JSON from given URI
func (d *Downloader) GetJsonMetadata(uri string) (*types.JsonMetadata, error) {
	data, _, err := d.getFromUri(uri)
	if err != nil {
		return nil, fmt.Errorf("unable to download json; %s", err)
	}
	jsonMeta, err := d.decodeJson(data)
	if err != nil {
		return nil, fmt.Errorf("unable to decode json; %s", err)
	}
	return jsonMeta, nil
}

// GetImage downloads image from given URI and detect its mimetype
func (d *Downloader) GetImage(uri string) (image *types.Image, err error) {
	data, mimetype, err := d.getFromUri(uri)
	if err != nil {
		return nil, fmt.Errorf("unable to download image; %s", err)
	}
	if err == nil && mimetype == "" {
		mimetype = d.mimetypeFromImageUri(uri)
	}
	out := types.Image{
		Data:     data,
		Mimetype: mimetype,
	}
	return &out, nil
}

// GetImageThumbnail downloads image from given URI scale it to thumbnail size
func (d *Downloader) GetImageThumbnail(uri string) (thumbnail *types.Image, err error) {
	image, err := d.GetImage(uri)
	if err != nil || image == nil {
		return nil, err
	}
	thumb, err := createThumbnail(*image)
	return &thumb, err
}

// getFromUri resolves the URI and download file from the URI using appropriate protocol
func (d *Downloader) getFromUri(uri string) (data []byte, mimetype string, err error) {
	if strings.HasPrefix(uri, "data:") {
		return d.getFromDataUri(uri)
	}
	if ipfsUri := d.getIpfsUri(uri); ipfsUri != "" {
		return d.getFromIpfs(ipfsUri)
	}
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		return d.getFromHttp(uri)
	}
	return nil, "", errors.New("Unexpected URI scheme for " + uri)
}

// getIpfsUri try to obtain IPFS URI from the URI - returns empty string for non-ipfs uri
// This function is responsible for IPFS URI detection, unification and for conversion
// of known IPFS HTTP gateways URI into IPFS URI.
func (d *Downloader) getIpfsUri(uri string) string {
	if strings.HasPrefix(uri, "/ipfs/") {
		return uri
	}
	if strings.HasPrefix(uri, "ipfs://") {
		return "/ipfs/" + uri[7:]
	}
	if d.skipHttpGateways {
		if strings.HasPrefix(uri, "https://gateway.pinata.cloud/ipfs/") {
			return "/ipfs/" + uri[34:]
		}
		if strings.HasPrefix(uri, "https://ipfs.io/ipfs/") {
			return "/ipfs/" + uri[21:]
		}
		if idx := strings.Index(uri, ".mypinata.cloud/ipfs/"); idx > 8 && idx < 30 {
			return "/ipfs/" + uri[idx+21:]
		}
	}
	return ""
}

// getFromIpfs downloads the file from IPFS (URI is expected in "/ipfs/{CID}" form).
func (d *Downloader) getFromIpfs(uri string) (data []byte, mimetype string, err error) {
	reader, err := d.ipfsShell.Cat(uri)
	if err != nil {
		return nil, "", err
	}
	out, err := io.ReadAll(reader)
	if err != nil {
		return nil, "", err
	}
	return out, "", reader.Close()
}

// getFromHttp downloads the file from HTTP.
func (d *Downloader) getFromHttp(uri string) (data []byte, mimetype string, err error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, "", err
	}
	if resp.StatusCode != 200 {
		return nil, "", fmt.Errorf("HTTP server returned %s", resp.Status)
	}
	reader := resp.Body
	out, err := io.ReadAll(reader)
	if err != nil {
		return nil, "", err
	}

	mimetype = resp.Header.Get("Content-Type")
	return out, mimetype, reader.Close()
}

// getFromDataUri obtains the file encoded in "data:" URI.
func (d *Downloader) getFromDataUri(uri string) (data []byte, mimetype string, err error) {
	splitted := strings.Split(uri, ",")
	if len(splitted) < 2 {
		return nil, "", errors.New("Invalid data uri - no comma: " + uri)
	}
	mimetype = strings.Split(splitted[0][5:], ";")[0]

	out, err := base64.StdEncoding.DecodeString(splitted[1])
	if err != nil {
		return nil, "", err
	}
	return out, mimetype, nil
}

// decodeJson parses the NFT token Metadata JSON.
func (d *Downloader) decodeJson(data []byte) (*types.JsonMetadata, error) {
	var out types.JsonMetadata
	err := json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// mimetypeFromImageUri tries to guess the image mime-type from extension in URI.
// To be used when the protocol (like IPFS) does not provide mime-type info.
func (d *Downloader) mimetypeFromImageUri(uri string) (mimetype string) {
	uri = strings.ToLower(uri)
	if strings.HasSuffix(uri, ".svg") {
		return "image/svg+xml"
	}
	if strings.HasSuffix(uri, ".gif") {
		return "image/gif"
	}
	if strings.HasSuffix(uri, ".jpg") || strings.HasSuffix(uri, ".jpeg") {
		return "image/jpeg"
	}
	if strings.HasSuffix(uri, ".png") {
		return "image/png"
	}
	if strings.HasSuffix(uri, ".webp") {
		return "image/webp"
	}
	return ""
}
