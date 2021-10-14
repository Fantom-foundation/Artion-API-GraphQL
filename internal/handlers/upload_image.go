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
	"net/http"
)

type uploadProcessor func(identity common.Address, image types.Image) (string, error)

// UploadImageHandler builds a HTTP handler function for images (tokens, user avatars) upload.
func UploadImageHandler(log logger.Logger, process uploadProcessor) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("Panic in UploadImageHandler handler; %s", r)
				w.WriteHeader(500)
				w.Write([]byte("Request handling failed"))
			}
		}()

		statusCode, response := processImageUpload(req, process, log)
		w.WriteHeader(statusCode)
		w.Write([]byte(response))
	})
}

func processImageUpload(req *http.Request, process uploadProcessor, log logger.Logger) (statusCode int, response string) {
	identity, err := auth.GetIdentityOrErr(req.Context())
	if err != nil {
		return 401, "Unauthorized"
	}

	err = req.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		return 500, "Unable to parse multipart/form-data"
	}

	file, _, err := req.FormFile("file")
	defer file.Close()
	if err != nil {
		return 500, "Unable to parse multipart/form-data file \"file\""
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return 500, "Unable to read uploaded file"
	}
	content := buf.Bytes()

	contentType := http.DetectContentType(content)
	log.Infof("Content type from DetectContentType: %s", contentType)
	imgType := types.ImageTypeFromMimetype(contentType)
	if imgType == types.ImageTypeUnknown {
		return 500, "Unrecognized image type " + contentType
	}

	image := types.Image{
		Data: content,
		Type: imgType,
	}

	response, err = process(*identity, image)
	if err != nil {
		log.Errorf("upload failed; %s", err)
		return 500, err.Error()
	}
	log.Info("Uploaded OK")
	return 200, response
}

func StoreUserAvatar(identity common.Address, image types.Image) (string, error) {
	err := repository.R().UploadUserAvatar(identity, image)
	if err != nil {
		return "", fmt.Errorf("user avatar upload failed; %s", err)
	}
	return "OK", nil
}

func StoreUserBanner(identity common.Address, image types.Image) (string, error) {
	err := repository.R().UploadUserBanner(identity, image)
	if err != nil {
		return "", fmt.Errorf("user banner upload failed; %s", err)
	}
	return "OK", nil
}
