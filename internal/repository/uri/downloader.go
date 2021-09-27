package uri

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/types"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/log"
	ipfsapi "github.com/ipfs/go-ipfs-api"
	"io"
	"net/http"
	"strings"
)

type downloader struct {
	ipfsShell *ipfsapi.Shell
}

func New(cfg *config.Config) Downloader {
	return &downloader{
		ipfsShell: ipfsapi.NewShell(cfg.Ipfs.Url),
	}
}

func (d *downloader) GetJsonMetadata(uri string) (*types.JsonMetadata, error) {
	data, _, err := d.getFromUri(uri)
	if err != nil {
		return nil, err
	}
	return d.decodeJson(data)
}

func (d *downloader) GetImage(uri string) (image *types.Image, err error) {
	data, mimetype, err := d.getFromUri(uri)
	if err == nil && mimetype == "" {
		mimetype = d.mimetypeFromImageUri(uri)
	}
	out := types.Image{
		Data: data,
		Mimetype: mimetype,
	}
	return &out, nil
}

func (d *downloader) getFromUri(uri string) (data []byte, mimetype string, err error) {
	log.Debug("Loading url "+ uri)
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
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		return d.getFromHttp(uri)
	}
	return nil, "", errors.New("Unexpected URI scheme for " + uri)
}

func (d *downloader) getFromIpfs(uri string) (data []byte, mimetype string, err error) {
	reader,err := d.ipfsShell.Cat(uri)
	if err != nil {
		return nil, "", err
	}
	out, err := io.ReadAll(reader)
	if err != nil {
		return nil, "", err
	}
	return out, "", reader.Close()
}

func (d *downloader) getFromHttp(uri string) (data []byte, mimetype string, err error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, "", err
	}
	reader := resp.Body
	out, err := io.ReadAll(reader)
	if err != nil {
		return nil, "", err
	}

	mimetype = resp.Header.Get("Content-Type")
	return out, mimetype, reader.Close()
}

func (d *downloader) getFromDataUri(uri string) (data []byte, mimetype string, err error) {
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

func (d *downloader) decodeJson(data []byte) (*types.JsonMetadata, error) {
	var out types.JsonMetadata
	err := json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (d *downloader) mimetypeFromImageUri(uri string) (mimetype string) {
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
