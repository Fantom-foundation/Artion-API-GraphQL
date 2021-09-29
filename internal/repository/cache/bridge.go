// Package cache provides in-memory shared cache.
package cache

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"github.com/allegro/bigcache"
	"time"
)

// config represents the configuration setup used by the repository
// to establish and maintain required connectivity to external services
// as needed.
var cfg *config.Config

// log represents the logger to be used by the repository.
var log logger.Logger

// MemCache represents in-memory cache.
type MemCache struct {
	cache *bigcache.BigCache
}

// New creates a new in-memory bridge instance.
func New() *MemCache {
	log.Debugf("creating in-memory cache")
	c, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             2048,
		LifeWindow:         cfg.Cache.Eviction,
		CleanWindow:        5 * time.Minute,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       1024,
		Verbose:            false,
		HardMaxCacheSize:   cfg.Cache.MaxSize,
		Logger:             log,
	})
	if err != nil {
		log.Criticalf("can not create cache; %s", err.Error())
		return nil
	}
	return &MemCache{cache: c}
}

// SetConfig sets the repository configuration to be used to establish
// and maintain external repository connections.
func SetConfig(c *config.Config) {
	cfg = c
}

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l.ModuleLogger("cache")
}
