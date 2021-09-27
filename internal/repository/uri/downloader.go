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

func (ud *downloader) GetJsonMetadata(uri string) (*types.JsonMetadata, error) {
	data, err := ud.getFromUri(uri)
	if err != nil {
		return nil, err
	}
	return ud.decodeJson(data)
}

func (ud *downloader) GetImage(uri string) ([]byte, error) {
	return ud.getFromUri(uri)
}

func (ud *downloader) getFromUri(uri string) ([]byte, error) {
	log.Debug("Loading url "+ uri)
	if strings.HasPrefix(uri, "data:") {
		return ud.getFromDataUri(uri)
	}
	if strings.HasPrefix(uri, "/ipfs/") {
		return ud.getFromIpfs(uri)
	}
	if strings.HasPrefix(uri, "ipfs://") {
		uri = "/ipfs/" + uri[7:]
		return ud.getFromIpfs(uri)
	}
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		return ud.getFromHttp(uri)
	}
	return nil, errors.New("Unexpected URI scheme for " + uri)
}

func (ud *downloader) getFromIpfs(uri string) ([]byte, error) {
	reader,err := ud.ipfsShell.Cat(uri)
	if err != nil {
		return nil, err
	}
	out, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return out, reader.Close()
}

func (ud *downloader) getFromHttp(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	reader := resp.Body
	out, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return out, reader.Close()
}

func (ud *downloader) getFromDataUri(uri string) ([]byte, error) {
	splitted := strings.Split(uri, ",")
	if len(splitted) < 2 {
		return nil, errors.New("Invalid data uri - no comma: " + uri)
	}
	out, err := base64.StdEncoding.DecodeString(splitted[1])
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ud *downloader) decodeJson(data []byte) (*types.JsonMetadata, error) {
	var out types.JsonMetadata
	err := json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
