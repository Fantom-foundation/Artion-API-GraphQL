package filecache

import (
	"artion-api-graphql/internal/logger"
	"errors"
	"os"
)

// log represents the logger to be used by the filecache.
var log logger.Logger

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l.ModuleLogger("cache")
}

const fileCacheDir = "/tmp/artion/"

func PullIpfsFile(cid string) (content []byte) {
	content, err := os.ReadFile(fileCacheDir + cid)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Errorf("unable to read filecache; %s", err)
	}
	return content
}

func PushIpfsFile(cid string, content []byte) (err error) {
	return os.WriteFile(fileCacheDir+cid, content, 0644)
}
