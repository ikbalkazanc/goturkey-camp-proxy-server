package main

import "sync"

type ReqularCacheMap struct {
	sync.RWMutex
	internal map[string]Cache
}

func NewReqularCacheMap() *ReqularCacheMap {
	return &ReqularCacheMap{
		internal: make(map[string]Cache),
	}
}

func (rm *ReqularCacheMap) Load(key string) (value Cache, ok bool) {

	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *ReqularCacheMap) Delete(key string) {

	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *ReqularCacheMap) Store(key string, value Cache) {

	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}
