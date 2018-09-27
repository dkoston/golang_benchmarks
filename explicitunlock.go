package golang_unlock_benchmark

import (
	"errors"
	"sync"
)

const ENODATA = 61

type Cache struct {
	Locker sync.Mutex
	Cache map[string]int
}

func NewCache() *Cache {
	return &Cache{
		Cache: make(map[string]int),
	}
}

func (s *Cache) SetCacheItem(key string, value int) {
	s.Cache[key] = value
}

func (s *Cache) CacheItem(key string) (value int, err error) {
	value, ok := s.Cache[key]
	if !ok {
		return value, errors.New(string(ENODATA))
	}
	return value, nil
}

func (s *Cache) RemoveCacheItem(key string) (err error) {
	_, ok := s.Cache[key]
	if !ok {
		return errors.New(string(ENODATA))
	}
	delete(s.Cache, key)
	return nil
}

func getTestValues(i int) (string, string) {
	if i % 3 == 0 {
		return "bar", "bar"
	} else if i % 2 == 0 {
		return "foo", "bar"
	}

	return "foo", "foo"
}

func DeferFunction(check1 string, check2 string) (value int, err error){
	cache := NewCache()

	cache.Locker.Lock()
	defer cache.Locker.Unlock()

	cache.SetCacheItem("foo", 1)

	value, err = cache.CacheItem(check1)
	if err != nil {
		return value, err
	}

	err = cache.RemoveCacheItem(check2)
	if err != nil {
		return  value, err
	}

	return  value, nil
}

func ExplicitUnlockFunction(check1 string, check2 string) (value int, err error){
	cache := NewCache()

	cache.Locker.Lock()

	cache.SetCacheItem("foo", 1)

	value, err = cache.CacheItem(check1)
	if err != nil {
		cache.Locker.Unlock()
		return value, err
	}

	err = cache.RemoveCacheItem(check2)
	if err != nil {
		cache.Locker.Unlock()
		return  value, err
	}

	cache.Locker.Unlock()
	return  value, nil
}
