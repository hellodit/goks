package lru

import (
	"container/list"
	"errors"
	"goks/cache"
)

type Repository struct {
	maxSizeItem       uint64
	items             map[string]*list.Element
	itemsPositionList *list.List
}

func (lru *Repository) GetKeys() (keys []string, err error) {
	keys = make([]string, len(lru.items))
	i := 0
	for elem := lru.itemsPositionList.Back(); elem != nil; elem = elem.Prev() {
		keys[i] = elem.Value.(*cache.Topic).Key
		i++
	}
	return
}

func NewLruCache(size uint64) *Repository {
	c := &Repository{
		maxSizeItem:       size,
		items:             make(map[string]*list.Element),
		itemsPositionList: list.New(),
	}
	return c
}

func (lru *Repository) GetOldest() (res *cache.Topic, err error) {
	elem := lru.itemsPositionList.Back()
	if elem != nil {
		res = elem.Value.(*cache.Topic)
		return
	}
	return
}

func (lru *Repository) RemoveOldest() error {
	elem := lru.itemsPositionList.Back()
	if elem != nil {
		lru.removeElement(elem)
	}

	return errors.New("cache missing")
}

func (lru *Repository) removeElement(e *list.Element) {
	lru.itemsPositionList.Remove(e)
	top := e.Value.(*cache.Topic)
	delete(lru.items, top.Key)
}

func (lru *Repository) Set(topic *cache.Topic) error {
	// Check for exiting key items
	if ent, exist := lru.items[topic.Key]; exist {
		lru.itemsPositionList.MoveToFront(ent)
		ent.Value.(*cache.Topic).Value = topic.Value
		return errors.New("can't add the cache, key has been used")
	}

	// push front
	res := lru.itemsPositionList.PushFront(topic)
	lru.items[topic.Key] = res

	// Verify size

	if uint64(lru.itemsPositionList.Len()) > lru.maxSizeItem {
		lru.RemoveOldest()
	}

	return nil
}

func (lru *Repository) Get(key string) (result *cache.Topic, err error) {

	if ent, exist := lru.items[key]; exist {
		result = ent.Value.(*cache.Topic)
		lru.itemsPositionList.MoveToFront(ent)
		return
	}

	return nil, errors.New("cache missed")
}

// Peek returns the key value (or undefined if not found) without updating
// the "recently used"-ness of the key.
func (lru *Repository) Peek(key string) (res *cache.Topic, err error) {
	if elem, ok := lru.items[key]; ok {
		res = elem.Value.(*cache.Topic)
		return
	}
	return nil, errors.New("cache missed")
}
