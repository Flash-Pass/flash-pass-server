package cache

import (
	"errors"
	cmap "github.com/orcaman/concurrent-map/v2"
)

var (
	KeyNotFoundError = errors.New("key not found")
)

type Cache[V any] struct {
	m *cmap.ConcurrentMap[string, V]
}

func New[V any]() *Cache[V] {
	m := cmap.New[V]()
	return &Cache[V]{
		m: &m,
	}
}

func (c *Cache[V]) Set(key string, value V) {
	c.m.Set(key, value)
}

func (c *Cache[V]) Get(key string) (V, bool) {
	v, ok := c.m.Get(key)
	if ok {
		return v, ok
	} else {
		var zero V
		return zero, ok
	}
}

func (c *Cache[V]) Update(
	key string, value V,
	f func(exist bool, valueInMap V, newValue V) V,
) V {
	return c.m.Upsert(key, value, f)
}

func (c *Cache[V]) Remove(key string) {
	c.m.Remove(key)
}

func (c *Cache[V]) Iter(f func(key string, v V)) {
	c.m.IterCb(f)
}

func (c *Cache[V]) Count() int {
	return c.m.Count()
}

func (c *Cache[V]) Clear() {
	keys := make([]string, 0, c.m.Count())
	c.Iter(func(key string, v V) {
		keys = append(keys, key)
	})
	for _, key := range keys {
		c.m.Remove(key)
	}
}
