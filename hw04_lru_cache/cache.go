package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mutex    sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item, ok := c.items[key]

	if ok {
		c.queue.MoveToFront(item)
		item.Value = cacheItem{key: key, value: value}
		return true
	}

	i := cacheItem{key: key, value: value}
	c.queue.PushFront(i)
	c.items[key] = c.queue.Front()

	if len(c.items) > c.capacity {
		out := c.queue.Back()
		c.queue.Remove(out)
		delete(c.items, out.Value.(cacheItem).key)
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}
	c.queue.MoveToFront(item)
	return item.Value.(cacheItem).value, true
}

func (c *lruCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.queue = NewList()

	for key := range c.items {
		delete(c.items, key)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
