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

func NewClient(options ...*cache.Option) (c cache.CacheRepository) {
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

func NewCacheOptions() (op *cache.Option) {
	return &cache.Option{}
}

func NewCacheRepository(option cache.Option) algo.Repository {
	var cacheRepo algo.Repository
	cacheRepo = lru.NewLruCache(option.MaxSizeItem)
	return cacheRepo
}

func mergeCacheOptions(options ...*cache.Option) (opts *cache.Option) {
	opts = new(cache.Option)
	// Check given option
	for _, op := options {
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
	c.RUnlock()
	if err != nil {
		return
	}
	return topic.Value, nil
}

func (c *Cache) Delete(key string) (err error) {
	panic("implement me")
}

func (c *Cache) GetKeys() (keys []string, err error) {
	c.RLock()
	keys, err = c.repo.GetKeys()
	if err != nil {
		return
	}
	c.RUnlock()
	return
}

func (c *Cache) ClearCache() (err error) {
	// Acquire a write lock before modifying the cache
	c.Lock()
	defer c.Unlock()

	// Remove all elements from the itemsPositionList and clear the items map in the Repository struct
	c.repo = NewCacheRepository(cache.Option{MaxSizeItem: cache.DefaultCapacity})

	// Return an error if any issues occur during the clearing process
	return nil
}

func (c *Cache) PeekByKey(key string) (val interface{}, err error) {
	c.RLock()
	topic, err := c.repo.Peek(key)
	if err != nil {
		return nil, err
	}
	c.RUnlock()
	return topic, nil
}
