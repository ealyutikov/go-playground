package lru_cache

import (
	"container/list"
)

type LRUCache struct {
	storage  map[int]*list.Element
	list     *list.List
	capacity int
}

type CacheValue struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{make(map[int]*list.Element, capacity), list.New(), capacity}
}

func (cache *LRUCache) Get(key int) int {
	if elem, ok := cache.storage[key]; ok {
		cache.list.MoveToFront(elem)
		return elem.Value.(CacheValue).value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	//update
	if elem, ok := cache.storage[key]; ok {
		elem.Value = CacheValue{key, value}
		cache.list.MoveToFront(elem)
		return
	}

	//full
	if cache.list.Len() == cache.capacity {
		last := cache.list.Back()
		cache.list.Remove(last)
		delete(cache.storage, last.Value.(CacheValue).key)
	}

	//put
	cache.storage[key] = cache.list.PushFront(CacheValue{key, value})
}
