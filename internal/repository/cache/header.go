// Package cache provides in-memory shared cache.
package cache

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/core/types"
	"strings"
)

// headerCacheKey generates a string key for the block header.
func headerCacheKey(id uint64) string {
	var sb strings.Builder
	sb.WriteString("header_")
	sb.WriteString(hexutil.Uint64(id).String())
	return sb.String()
}

// PullHeader tries to pull the given block head from the in-memory cache, if possible.
func (c *MemCache) PullHeader(id uint64) *eth.Header {
	data, err := c.cache.Get(headerCacheKey(id))
	if err != nil {
		return nil
	}

	var blk eth.Header
	if err := json.Unmarshal(data, &blk); err != nil {
		return nil
	}
	return &blk
}

// PushHeader tries to store the given block header to in-memory cache.
func (c *MemCache) PushHeader(hdr *eth.Header) {
	if nil == hdr {
		return
	}

	data, err := json.Marshal(hdr)
	if err != nil {
		log.Errorf("can not encode header; %s", err.Error())
		return
	}

	err = c.cache.Set(headerCacheKey(hdr.Number.Uint64()), data)
	if err != nil {
		log.Errorf("can not store header in cache; %s", err.Error())
	}
}
