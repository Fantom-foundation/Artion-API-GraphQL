package handlers

import (
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"net/http"
	"strings"
)

// TokenImage builds a HTTP handler function for Token images.
func TokenImage(log logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("Panic in TokenImage handler; %s", r)
				w.WriteHeader(500)
				w.Write([]byte("Request handling failed"))
			}
		}()

		uri, err := resolveUri(req.URL.Path)
		if err != nil {
			log.Errorf("token image request handling failed; %s", err)
			w.WriteHeader(500)
			_,_ = w.Write([]byte("Request handling failed: " + err.Error()))
			return
		}

		image, err := repository.R().GetTokenImage(uri)
		if err != nil {
			log.Errorf("unable to get token image; %s", err)
			w.WriteHeader(500)
			_,_ = w.Write([]byte("Obtaining token image failed: " + err.Error()))
		}
		if image == nil {
			w.WriteHeader(404)
			_,_ = w.Write([]byte("The token has no image"))
			return
		}

		w.Header().Add("Content-Type", image.Mimetype)
		w.WriteHeader(200)
		_, err = w.Write(image.Data)
		if err != nil {
			log.Errorf("writing token image response failed; %s", err)
		}
	})
}

func resolveUri(path string) (imageUri string, err error) {
	pathParts := strings.Split(path, "/")
	if len(pathParts) != 4 {
		return "", errors.New("invalid amount of slash delimiters in URL")
	}
	tokenAddress := common.HexToAddress(pathParts[2])
	tokenId, err := hexutil.DecodeBig(pathParts[3])
	if err != nil {
		return "", fmt.Errorf("unable to hex-decode tokenId; %s", err)
	}

	tok, err := repository.R().GetToken(tokenAddress, hexutil.Big(*tokenId))
	if err != nil {
		return "", fmt.Errorf("unable to get token in db; %s", err)
	}

	jsonMetadata, err := repository.R().GetTokenJsonMetadata(tok.Uri)
	if err != nil {
		return "", fmt.Errorf("unable to get token json metadata; %s", err)
	}
	if jsonMetadata.Image == nil || *jsonMetadata.Image == "" {
		return "", fmt.Errorf("token has no image; %s", err)
	}
	return *jsonMetadata.Image, nil
}
