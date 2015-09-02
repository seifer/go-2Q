package lru

import (
	"container/list"
)

type item struct {
	active  bool
	element *list.Element
}

type cache struct {
	max, cnt int
	index    map[interface{}]*item
	active   *list.List
	inactive *list.List
}

func New(max int) *cache {
	return &cache{
		max:      max,
		index:    make(map[interface{}]*item),
		active:   new(list.List),
		inactive: new(list.List),
	}
}

func (c *cache) Add(v interface{}) {
	if ci, ok := c.index[v]; ok {
		if ci.active {
			c.active.MoveToFront(ci.element)
		} else {
			c.inactive.Remove(ci.element)

			ci.active = true
			ci.element = c.active.PushFront(v)
		}

		return
	}

	c.cnt++
	c.index[v] = &item{false, c.inactive.PushFront(v)}
}

func (c *cache) Reclaim() (res []interface{}) {
	for c.cnt > c.max && c.inactive.Len() > 0 {
		c.cnt--
		res = append(res, c.inactive.Remove(c.inactive.Back()))
	}

	for c.cnt > c.max && c.active.Len() > 0 {
		c.cnt--
		res = append(res, c.active.Remove(c.active.Back()))
	}

	return
}
