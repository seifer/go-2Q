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

func (c *cache) add(v interface{}) (*item, bool) {
	if i, ok := c.index[v]; ok {
		if i.active {
			c.active.MoveToFront(i.element)
		} else {
			c.inactive.Remove(i.element)

			i.active = true
			i.element = c.active.PushFront(v)
		}

		return i, false
	}

	c.index[v] = &item{0, false, c.inactive.PushFront(v)}

	return c.index[v], true
}

func (c *cache) del(v interface{}) *item {
	if i, ok := c.index[v]; ok {
		if i.active {
			c.active.Remove(i.element)
		} else {
			c.inactive.Remove(i.element)
		}

		delete(c.index, v)

		return i
	}

	return nil
}

func (c *cache) reclaim() (i *item) {
	if c.threshold <= c.max {
		return
	}

	if c.inactive.Len() > 0 {
		v := c.inactive.Remove(c.inactive.Back())
		i = c.index[v]
		delete(c.index, v)
		return
	}

	if c.active.Len() > 0 {
		v := c.active.Remove(c.active.Back())
		i = c.index[v]
		delete(c.index, v)
		return
	}

	return nil
}
