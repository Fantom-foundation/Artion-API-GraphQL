// Package filecache stores files uploaded into IPFS, so they are available
// before the local IPFS node is able to download them.
package filecache

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"errors"
	"fmt"
	"os"
)

// log represents the logger to be used by the filecache.
var log logger.Logger

// FileCache represents cache in a filesystem.
type FileCache struct {
	dir string
}

func New(cfg *config.Config) *FileCache {
	fc := FileCache{
		dir: cfg.Ipfs.FileCacheDir,
	}
	err := os.MkdirAll(fc.dir, 0700)
	if err != nil {
		panic(fmt.Errorf("unable to create filecache directory \"%s\"; %s", fc.dir, err))
	}
	return &fc
}

func (fc *FileCache) PullIpfsFile(cid string) (content []byte) {
	content, err := os.ReadFile(fc.dir+cid)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Errorf("unable to read filecache; %s", err)
	}
	return content
}

func (fc *FileCache) PushIpfsFile(cid string, content []byte) (err error) {
	return os.WriteFile(fc.dir+cid, content, 0644)
}

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l.ModuleLogger("filecache")
}
