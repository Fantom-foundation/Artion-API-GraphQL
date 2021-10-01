// Package blkcache implements circular cache for the latest block headers.
package blkcache

import eth "github.com/ethereum/go-ethereum/core/types"

// Cache represents a circular cache for the latest block headers
type Cache struct {
	list  []*eth.Header
	cap   int
	start int
	end   int
}

// NewCache creates a new circular cache.
func NewCache(capacity int) *Cache {
	return &Cache{
		cap:   capacity,
		list:  make([]*eth.Header, capacity, capacity),
		start: 0,
		end:   0,
	}
}

// Add new header to the cache.
func (c *Cache) Add(hdr *eth.Header) {
	c.list[c.end] = hdr

	// advance start
	adv(&c.end, c.cap)

	// check end
	if c.end == c.start {
		adv(&c.start, c.cap)
	}
}

// Pull pulls all the block headers from the cache
// resetting the cache to empty state.
func (c *Cache) Pull() []*eth.Header {
	size := c.end - c.start
	if size < 0 {
		size = c.cap
	}

	list := make([]*eth.Header, size)
	for i := 0; i < size; i++ {
		ix := c.start + i
		if ix >= c.cap {
			ix -= c.cap
		}

		list[i] = c.list[ix]
	}

	c.start = 0
	c.end = 0
	return list
}

// adv advances the given value respecting cap limit.
func adv(val *int, cap int) {
	*val++
	if *val >= cap {
		*val = 0
	}
}
