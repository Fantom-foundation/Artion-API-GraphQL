// Package uri provides tools for downloading data from URI - especially
// NFT tokens JSON Metadata and related images using IPFS or HTTP.
package uri

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/types"
	"encoding/base64"
	"errors"
	"fmt"
	ipfsapi "github.com/ipfs/go-ipfs-api"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ipfsRequestTimeout represents the timeout applied to IPFS requests.
const ipfsRequestTimeout = 30 * time.Second

type Downloader struct {
	ipfsShell        *ipfsapi.Shell
	skipHttpGateways bool
	gateway          string
	gatewayBearer    string
}

// New provides new Downloader instance.
func New(cfg *config.Config) *Downloader {
	d := &Downloader{
		ipfsShell:        ipfsapi.NewShell(cfg.Ipfs.Url),
		skipHttpGateways: cfg.Ipfs.SkipHttpGateways,
		gateway:          cfg.Ipfs.Gateway,
		gatewayBearer:    cfg.Ipfs.GatewayBearer,
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
	jsonMeta, err := types.DecodeJsonMetadata(data)
	if err != nil {
		return nil, fmt.Errorf("unable to decode json; %s", err)
	}
	return jsonMeta, nil
}

// GetImage downloads image from given URI and detect its mimetype
func (d *Downloader) GetImage(uri string) (*types.Image, error) {
	data, mimetype, err := d.getFromUri(uri)
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

// unescapeIPFSUri decodes URL escaped address.
func unescapeIPFSUri(uri string) string {
	iUri, err := url.QueryUnescape(uri)
	if err != nil {
		return "/ipfs/" + uri
	}
	return "/ipfs/" + iUri
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
			return unescapeIPFSUri(uri[34:])
		}
		if strings.HasPrefix(uri, "https://ipfs.io/ipfs/") {
			return unescapeIPFSUri(uri[21:])
		}
		if idx := strings.Index(uri, ".mypinata.cloud/ipfs/"); idx > 8 && idx < 30 {
			return unescapeIPFSUri(uri[idx+21:])
		}
	}
	return ""
}

// getFromIpfs downloads the file from IPFS (URI is expected in "/ipfs/{CID}" form).
func (d *Downloader) getFromIpfs(uri string) (data []byte, mimetype string, err error) {
	if d.gateway == "" {
		reader, err := d.ipfsShell.Cat(uri)
		if err != nil {
			return nil, "", err
		}
		out, err := io.ReadAll(reader)
		if err != nil {
			return nil, "", err
		}
		return out, "", reader.Close()
	} else {
		return d.getFromIpfsGateway(uri)
	}
}

// getFromIpfsGateway downloads the file from IPFS HTTP gateway.
func (d *Downloader) getFromIpfsGateway(uri string) (data []byte, mimetype string, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", d.gateway+uri, nil)
	if err != nil {
		return nil, "", err
	}
	if d.gatewayBearer != "" {
		req.Header.Set("Authorization", "Bearer "+d.gatewayBearer)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	if resp.StatusCode != 200 {
		return nil, "", fmt.Errorf("HTTP gateway returned %s", resp.Status)
	}
	reader := resp.Body
	out, err := io.ReadAll(reader)
	if err != nil {
		return nil, "", err
	}

	mimetype = resp.Header.Get("Content-Type")
	return out, mimetype, reader.Close()
}

// getFromHttp downloads the file from HTTP.
func (d *Downloader) getFromHttp(uri string) (data []byte, mimetype string, err error) {
	client := http.Client{
		Timeout: 1 * time.Minute,
	}

	resp, err := client.Get(uri)
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
	parts := strings.Split(uri, ",")
	if len(parts) < 2 {
		return nil, "", errors.New("Invalid data uri - no comma: " + uri)
	}
	mimetype = strings.Split(parts[0][5:], ";")[0]

	out, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, "", err
	}
	return out, mimetype, nil
}
