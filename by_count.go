package lru

type cacheByCount struct {
	cache
}

func NewByCount(max int64) *cacheByCount {
	return &cacheByCount{_new(max)}
}

func (c *cacheByCount) Add(v interface{}) {
	if i, ok := c.add(v); ok {
		c.threshold++
	}
}

func (c *cacheByCount) Del(v interface{}) {
	if i := c.del(v); i != nil {
		c.threshold--
	}
}

func (c *cacheByCount) Reclaim() (res []interface{}) {
	for i := c.reclaim(); i != nil; i = c.reclaim() {
		res = append(res, i.element.Value)
		c.threshold--
	}

	return
}
