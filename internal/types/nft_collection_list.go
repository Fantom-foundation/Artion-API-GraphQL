package types

type NFTCollectionList struct {
	// List keeps the actual NFTCollection.
	Collection []*NFTCollection

	// TotalCount indicates total number of results.
	TotalCount int64

	// HasPrev indicates there are some results before this results page.
	HasPrev bool

	// HasNext indicates there are some results after this results page.
	HasNext bool
}

func (c *NFTCollectionList) Reverse() {
	// anything to swap at all?
	if c.Collection == nil || len(c.Collection) < 2 {
		return
	}

	// swap elements
	for i, j := 0, len(c.Collection)-1; i < j; i, j = i+1, j-1 {
		c.Collection[i], c.Collection[j] = c.Collection[j], c.Collection[i]
	}

	// swap next/previous page flag
	c.HasNext, c.HasPrev = c.HasPrev, c.HasNext
}
