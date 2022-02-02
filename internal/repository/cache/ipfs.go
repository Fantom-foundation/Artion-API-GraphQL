package cache

func (c *MemCache) PullIpfsFile(cid string) (content []byte) {
	content, err := c.cache.Get("ipfs-cid-" + cid)
	if err != nil {
		return nil
	}
	return content
}

func (c *MemCache) PushIpfsFile(cid string, content []byte) (err error) {
	err = c.cache.Set("ipfs-cid-"+cid, content)
	if err != nil {
		log.Errorf("can not store ipfs file in cache; %s", err.Error())
	}
	return
}
