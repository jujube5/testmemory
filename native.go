package testmemory

import (
    "sync"
)

type Cache struct {
    sync.RWMutex
    Items map[string]CacheItems
}

type CacheItems struct {
    Item    any
}

var c Cache

func init() {
	items := make(map[string]CacheItems)
    c = Cache{Items: items}
}

func Set(key string, item any) {
    c.Lock()
    defer c.Unlock()
    c.Items[key] = CacheItems{Item: item}
}

func Get(key string) (any) {
    c.RLock()
    defer c.RUnlock()
    item := c.Items[key]
	return item.Item
}

func Delete(key string) {
    c.Lock()
    defer c.Unlock()
    delete(c.Items, key)
}
