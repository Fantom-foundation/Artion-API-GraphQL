package pinner

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

// log represents the logger to be used by the repository.
var log logger.Logger

// Pinner allows to pin files to Pinata IPFS service
type Pinner struct {
	pinataBearer string
}

// New provides new Pinner instance.
func New(cfg *config.Config) *Pinner {
	if cfg.Ipfs.GatewayBearer == "" {
		panic("unable to init Pinner - Pinata bearer token not configured")
	}
	return &Pinner{
		pinataBearer: cfg.Ipfs.GatewayBearer,
	}
}

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l.ModuleLogger("pinner")
}

// PinFile uploads the file to Pinata
// based on https://github.com/wabarc/ipfs-pinner/blob/v1.0.1/pkg/pinata/pinata.go
func (p Pinner) PinFile(filename string, content []byte) (cid string, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("file", filename)
		if err != nil {
			return
		}

		if _, err = part.Write(content); err != nil {
			return
		}
	}()

	req, err := http.NewRequest(http.MethodPost, "https://api.pinata.cloud/pinning/pinFileToIPFS", r)
	req.Header.Add("Content-Type", m.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+p.pinataBearer)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	log.Infof("pinning file " + filename)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(data, &dat); err != nil {
		return "", err
	}

	if errStr, hasErr := dat["error"].(string); hasErr {
		log.Errorf("pinata error: %s", errStr)
		return "", errors.New(errStr)
	}
	if hash, ok := dat["IpfsHash"].(string); ok {
		log.Infof("file pinned as " + hash)
		return hash, nil
	}
	log.Errorf("pinata returned no IpfsHash - response: %s", data)
	return "", errors.New("pinata returned no IpfsHash")
}
