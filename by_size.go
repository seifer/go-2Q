package lru

type cacheBySize struct {
	cache
}

func NewBySize(max int64) *cacheBySize {
	return &cacheBySize{_new(max)}
}

func (c *cacheBySize) Add(v interface{}, size int64) {
	if _, ok := c.add(v, size); ok {
		c.threshold += size
	}
}
