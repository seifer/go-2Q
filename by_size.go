package lru

type cacheBySize struct {
	cache
}

func NewBySize(max int64) *cacheBySize {
	return &cacheBySize{_new(max)}
}

func (c *cacheBySize) Add(v interface{}, size int64) {
	if i := c.add(v); i != nil {
		c.threshold++
	}
}

func (c *cacheBySize) Del(v interface{}) {
	if i := c.del(v); i != nil {
		c.threshold--
	}
}

func (c *cacheBySize) Reclaim() (res []interface{}) {
	for i := c.reclaim(); i != nil; i = c.reclaim() {
		res = append(res, i.element.Value)
		c.threshold--
	}

	return
}
