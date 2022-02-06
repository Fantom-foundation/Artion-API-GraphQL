package handlers

import (
	"artion-api-graphql/internal/auth"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

type uploadProcessor func(identity common.Address, image types.Image, req *http.Request) (string, error)

// UploadImageHandler builds a HTTP handler function for images (tokens, user avatars) upload.
func UploadImageHandler(log logger.Logger, process uploadProcessor) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("Panic in UploadImageHandler handler; %v", r)
				w.WriteHeader(500)

				_, err := fmt.Fprintf(w, "Request handling failed, could not upload; %v", r)
				if err != nil {
					log.Errorf("could not write HTTP content; %s", err.Error())
				}
			}
		}()

		statusCode, response := processImageUpload(req, process, log)

		w.WriteHeader(statusCode)
		_, err := w.Write([]byte(response))
		if err != nil {
			log.Errorf("could not write HTTP response; %s", err.Error())
		}
	})
}

// processImageUpload handles image upload request and provides appropriate output response.
func processImageUpload(req *http.Request, process uploadProcessor, log logger.Logger) (statusCode int, response string) {
	identity, err := auth.GetIdentityOrErr(req.Context())
	if err != nil {
		return 401, "Unauthorized"
	}

	err = req.ParseMultipartForm(cfg.Server.MaxParserMemory)
	if err != nil {
		return 500, "Unable to parse multipart/form-data"
	}

	file, _, err := req.FormFile("file")
	if err != nil {
		return 500, "Unable to parse multipart/form-data file \"file\""
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Errorf("could not close form file; %s", err.Error())
		}
	}(file)

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return 500, "Unable to read uploaded file"
	}
	content := buf.Bytes()

	imgType, err := types.ImageTypeFromMimetype(content)
	if err != nil {
		return 500, err.Error()
	}

	image := types.Image{
		Data: content,
		Type: imgType,
	}

	response, err = process(*identity, image, req)
	if err != nil {
		log.Errorf("upload failed; %s", err)
		return 500, err.Error()
	}
	log.Info("Uploaded OK")
	return 200, response
}

// StoreUserAvatar sends an avatar image to backed repository.
func StoreUserAvatar(identity common.Address, image types.Image, _ *http.Request) (string, error) {
	err := repository.R().UploadUserAvatar(identity, image)
	if err != nil {
		return "", fmt.Errorf("user avatar upload failed; %s", err)
	}
	return "OK", nil
}

// StoreUserBanner sends an user banner to the backend repository.
func StoreUserBanner(identity common.Address, image types.Image, _ *http.Request) (string, error) {
	err := repository.R().UploadUserBanner(identity, image)
	if err != nil {
		return "", fmt.Errorf("user banner upload failed; %s", err)
	}
	return "OK", nil
}

// StoreToken validates and finalizes NFT metadata structure and sends it for uploading.
func StoreToken(identity common.Address, image types.Image, req *http.Request) (string, error) {
	metadataJson := req.FormValue("metadata")
	if metadataJson == "" {
		return "", fmt.Errorf("no token metadata sent")
	}
	metadata, err := types.DecodeJsonMetadata([]byte(metadataJson))
	if err != nil {
		return "", fmt.Errorf("failed to parse json metadata; %s", err)
	}
	if metadata.Name == "" {
		return "", fmt.Errorf("token name not defined in json metadata")
	}

	// override author address
	metadata.Properties.Address = identity.String()
	metadata.Properties.Recipient = identity.String()

	// override createdAt
	metadata.Properties.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	uri, err := repository.R().UploadTokenData(*metadata, image)
	if err != nil {
		return "", fmt.Errorf("token data upload failed; %s", err)
	}
	return uri, nil
}

// StoreCollection validates NFT collection metadata and sends the validated data to backend repository.
func StoreCollection(identity common.Address, image types.Image, req *http.Request) (string, error) {
	applicationJson := req.FormValue("data")
	if applicationJson == "" {
		return "", fmt.Errorf("no collection registration application sent")
	}

	app, err := types.DecodeCollectionApplication([]byte(applicationJson))
	if err != nil {
		return "", fmt.Errorf("failed to parse collection registration application json; %s", err)
	}

	err = repository.R().CanRegisterCollection(&app.Contract, &identity)
	if err != nil {
		return "", err
	}

	err = repository.R().UploadCollectionApplication(*app, image, identity)
	if err != nil {
		return "", fmt.Errorf("collection upload failed; %s", err)
	}

	return "OK", nil
}
