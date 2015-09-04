package lru

type cacheByCount struct {
	cache
}

func NewByCount(max int64) *cacheByCount {
	return &cacheByCount{_new(max)}
}

func (c *cacheByCount) Add(v interface{}) {
	if _, ok := c.add(v, 1); ok {
		c.threshold++
	}
}
