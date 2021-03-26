package algo

import "goks/cache"

type Repository interface {
	Set(topic *cache.Topic) (err error)
	Get(key string) (result *cache.Topic, err error)
	GetOldest() (res *cache.Topic, err error)
	RemoveOldest() (err error)
	Peek(key string) (res *cache.Topic, err error)
	GetKeys() (keys []string, err error)
}
