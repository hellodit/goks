package cache

const (
	DefaultCapacity = 50
)

// Topic represent cache key:value
type Topic struct {
	Key   string
	Value interface{}
}

// Option represent cache option
type Option struct {
	MaxSizeItem uint64
}

// SetMaxSizeItem max memory item
func (opt *Option) SetMaxSizeItem(size uint64) *Option {
	opt.MaxSizeItem = size
	return opt
}

// Cache represent the public API that will available used by user
type CacheRepository interface {
	Set(key string, value interface{}) error
	Get(key string) (val interface{}, err error)
	Delete(key string) (err error)
	ClearCache() (err error)
	PeekCache() (err error)
}
