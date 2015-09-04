package lru

import (
	"container/list"
)

type item struct {
	value   int64
	active  bool
	element *list.Element
}

type cache struct {
	active   *list.List
	inactive *list.List

	threshold, max int64

	index map[interface{}]*item
}

func _new(max int64) cache {
	return cache{
		max:      max,
		index:    make(map[interface{}]*item),
		active:   list.New(),
		inactive: list.New(),
	}
}

func (c *cache) add(v interface{}, iv int64) (*item, bool) {
	if i, ok := c.index[v]; ok {
		if i.active {
			c.active.MoveToFront(i.element)
		} else {
			c.inactive.Remove(i.element)

			i.active = true
			i.element = c.active.PushFront(v)
		}

		i.value = iv

		return i, false
	}

	c.index[v] = &item{
		iv,
		false,
		c.inactive.PushFront(v),
	}

	return c.index[v], true
}

func (c *cache) Del(v interface{}) {
	if i, ok := c.index[v]; ok {
		if i.active {
			c.active.Remove(i.element)
		} else {
			c.inactive.Remove(i.element)
		}

		delete(c.index, v)
		c.threshold -= i.value

		return
	}

	return
}

func (c *cache) Reclaim() (res []interface{}) {
	if c.threshold <= c.max {
		return
	}

	for c.threshold > c.max && c.inactive.Len() > 0 {
		v := c.inactive.Remove(c.inactive.Back())
		c.threshold -= c.index[v].value
		res = append(res, v)
		delete(c.index, v)
	}

	for c.threshold > c.max && c.active.Len() > 0 {
		v := c.active.Remove(c.active.Back())
		c.threshold -= c.index[v].value
		res = append(res, v)
		delete(c.index, v)
	}

	return
}
