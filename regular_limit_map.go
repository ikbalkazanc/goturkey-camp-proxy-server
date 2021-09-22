package main

import "sync"

type RegularLimitMap struct {
	sync.RWMutex
	internal map[string]Limit
}

func NewRegularLimitMap() *RegularLimitMap {
	return &RegularLimitMap{
		internal: make(map[string]Limit),
	}
}

func (rm *RegularLimitMap) Load(key string) (value Limit, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *RegularLimitMap) Increase(key string) (value Limit, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	result.count++
	rm.internal[key] = result
	rm.RUnlock()
	return result, ok
}

func (rm *RegularLimitMap) Delete(key string) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *RegularLimitMap) Store(key string, value Limit) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}
