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

// TokenMetadataHandler builds a HTTP handler function for Token JSON metadata.
func TokenMetadataHandler(log logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("Panic in TokenMetadataHandler; %s", r)

				w.WriteHeader(500)
				_, _ = w.Write([]byte("Request handling failed, check server log for details."))
			}
		}()

		data, err := handleTokenMetadataRequest(req.URL.Path)
		if err != nil {
			log.Errorf("failed to handle token metadata request [%s]; %s", req.URL.Path, err)
			_, _ = w.Write([]byte(fmt.Sprintf("Request handling failed: %s", err)))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		_, err = w.Write(data)
		if err != nil {
			log.Errorf("writing json response failed; %s", err)
		}
	})
}

func handleTokenMetadataRequest(uri string) ([]byte, error) {
	pathParts := strings.Split(uri, "/")
	if len(pathParts) != 5 {
		return nil, errors.New("invalid amount of slash delimiters in URL")
	}
	tokenAddress := common.HexToAddress(pathParts[3])
	tokenId, err := hexutil.DecodeBig(pathParts[4])
	if err != nil {
		return nil, fmt.Errorf("unable to hex-decode tokenId; %s", err)
	}

	tok, err := repository.R().Token(&tokenAddress, (*hexutil.Big)(tokenId))
	if tok == nil {
		return nil, fmt.Errorf("unable to get token in db; %s", err)
	}

	return repository.R().GetRawTokenJsonMetadata(tok.Uri)
}

