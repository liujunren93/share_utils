package lru

import (
	"container/list"
	"errors"
	"time"
)

// EvictCallback is used to get a callback when a cache entry is evicted
type EvictCallback func(key interface{}, value interface{})

// LRU implements a non-thread safe fixed size LRU cache
type LRU struct {
	size      int64
	expire    int64
	evictList *list.List
	items     map[interface{}]*list.Element
}

// entry is used to hold a value in the evictList
type entry struct {
	key      interface{}
	value    interface{}
	createAt int64
}

// NewLRU constructs an LRU of the given size
func NewLRU(size, expire int64) (*LRU, error) {
	if size <= 0 && expire <= 0 {
		return nil, errors.New("Must provide a positive size or expire")
	}
	c := &LRU{
		size:      size,
		evictList: list.New(),
		expire:    expire,
		items:     make(map[interface{}]*list.Element),
	}
	return c, nil
}

// Purge is used to completely clear the cache.
func (c *LRU) Purge() {
	c.items = make(map[interface{}]*list.Element)
	c.evictList.Init()
}

// Add adds a value to the cache.  Returns true if an eviction occurred.
func (c *LRU) Add(key, value interface{}) (evicted bool) {
	now := time.Now().Local().Unix()
	// Check for existing item
	if ent, ok := c.items[key]; ok {
		ent.Value.(*entry).createAt = now
		ent.Value.(*entry).value = value
		c.evictList.MoveToFront(ent)

		return false
	}

	// Add new item
	ent := &entry{key, value, now}
	element := c.evictList.PushFront(ent)
	c.items[key] = element
	var evict bool
	if c.size > 0 {
		evict = c.evictList.Len() > int(c.size)
	}
	if c.expire > 0 {
		evict = now-c.evictList.Back().Value.(*entry).createAt >= c.expire
	}
	// Verify size not exceeded
	if evict {
		c.removeOldest()
	}
	return evict
}

// Get looks up a key's value from the cache.
func (c *LRU) Get(key interface{}) (value interface{}, ok bool) {
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		if ent.Value.(*entry) == nil {
			return nil, false
		}
		return ent.Value.(*entry).value, true
	}
	return
}

// Contains checks if a key is in the cache, without updating the recent-ness
// or deleting it for being stale.
func (c *LRU) Contains(key interface{}) (ok bool) {
	_, ok = c.items[key]
	return ok
}

// Peek returns the key value (or undefined if not found) without updating
// the "recently used"-ness of the key.
func (c *LRU) Peek(key interface{}) (value interface{}, ok bool) {
	var ent *list.Element
	if ent, ok = c.items[key]; ok {
		return ent.Value.(*entry).value, true
	}
	return nil, ok
}

// Remove removes the provided key from the cache, returning if the
// key was contained.
func (c *LRU) Remove(key interface{}) (present bool) {
	if ent, ok := c.items[key]; ok {
		c.removeElement(ent)
		return true
	}
	return false
}

// RemoveOldest removes the oldest item from the cache.
func (c *LRU) RemoveOldest() (key interface{}, value interface{}, ok bool) {
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
		kv := ent.Value.(*entry)
		return kv.key, kv.value, true
	}
	return nil, nil, false
}

// GetOldest returns the oldest entry
func (c *LRU) GetOldest() (key interface{}, value interface{}, ok bool) {
	ent := c.evictList.Back()
	if ent != nil {
		kv := ent.Value.(*entry)
		return kv.key, kv.value, true
	}
	return nil, nil, false
}

// Keys returns a slice of the keys in the cache, from oldest to newest.
func (c *LRU) Keys() []interface{} {
	keys := make([]interface{}, len(c.items))
	i := 0
	for ent := c.evictList.Back(); ent != nil; ent = ent.Prev() {
		keys[i] = ent.Value.(*entry).key
		i++
	}
	return keys
}

// Len returns the number of items in the cache.
func (c *LRU) Len() int {
	return c.evictList.Len()
}

// Resize changes the cache size.
func (c *LRU) Resize(size int64) (evicted int) {
	diff := c.Len() - int(size)
	if diff < 0 {
		diff = 0
	}
	for i := 0; i < diff; i++ {
		c.removeOldest()
	}
	c.size = size
	return diff
}

// removeOldest removes the oldest item from the cache.
func (c *LRU) removeOldest() {
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
	}
}

// removeElement is used to remove a given list element from the cache
func (c *LRU) removeElement(e *list.Element) {
	c.evictList.Remove(e)
	kv := e.Value.(*entry)
	delete(c.items, kv.key)
	// if c.onEvict != nil {
	// 	c.onEvict(kv.key, kv.value)
	// }
}
