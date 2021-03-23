package goks

import (
	"goks/algo"
	"goks/algo/lru"
	"goks/cache"
	"sync"
)

// Cache ...
type Cache struct {
	sync.RWMutex
	repo algo.Repository
}



func NewClient(options ...*cache.Option)(c cache.CacheRepository){
	option := mergeCacheOptions(options...)
	if option.MaxSizeItem == 0 {
		// if not set use default size
		option.MaxSizeItem = cache.DefaultCapacity
	}

	c = &Cache{
		repo: NewCacheRepository(*option),
	}
	return c
}

func NewCacheOptions() (op *cache.Option){
	return &cache.Option{}
}

func NewCacheRepository(option cache.Option) algo.Repository{
	var cacheRepo algo.Repository
	cacheRepo = lru.NewLruCache(option.MaxSizeItem)
	return cacheRepo
}

func mergeCacheOptions(options ...*cache.Option) (opts *cache.Option)  {
	opts = new(cache.Option)
	// Check given option
	for _, op := range options{
		if op.MaxSizeItem != 0 {
			opts.MaxSizeItem = op.MaxSizeItem
		}
	}
	return
}

func (c *Cache) Set(key string, value interface{}) error {
	topic := &cache.Topic{
		Key:   key,
		Value: value,
	}
	c.Lock()
	err := c.repo.Set(topic)
	if err != nil {
		return err
	}
	c.Unlock()

	return nil
}

func (c *Cache) Get(key string) (val interface{}, err error) {
	c.RLock()
	topic, err := c.repo.Get(key)
	c.RLock()
	if err != nil {
		return
	}
	return topic.Value, nil
}

func (c *Cache) Delete(key string) (err error) {
	panic("implement me")
}

func (c *Cache) GetKeys() (keys []string, err error) {
	panic("implement me")
}

func (c *Cache) ClearCache() (err error) {
	panic("implement me")
}

func (c *Cache) PeekCache() (err error) {
	panic("implement me")
}
