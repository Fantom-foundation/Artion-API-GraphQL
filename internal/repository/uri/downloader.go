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
const ipfsRequestTimeout = 10 * time.Second

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

// getFromUri resolves the URI and download file from the URI using appropriate protocol
func (d *Downloader) getFromUri(uri string) (data []byte, mimetype string, err error) {
	if strings.HasPrefix(uri, "data:") {
		return d.getFromDataUri(uri)
	}
	if strings.HasPrefix(uri, "/ipfs/") {
		return d.getFromIpfs(uri)
	}
	if strings.HasPrefix(uri, "ipfs://") {
		uri = "/ipfs/" + uri[7:]
		return d.getFromIpfs(uri)
	}
	if d.skipHttpGateways {
		if strings.HasPrefix(uri, "https://gateway.pinata.cloud/ipfs/") {
			uri = "/ipfs/" + uri[34:]
			return d.getFromIpfs(uri)
		}
		if strings.HasPrefix(uri, "https://ipfs.io/ipfs/") {
			uri = "/ipfs/" + uri[21:]
			return d.getFromIpfs(uri)
		}
	}
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		return d.getFromHttp(uri)
	}
	return nil, "", errors.New("Unexpected URI scheme for " + uri)
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
