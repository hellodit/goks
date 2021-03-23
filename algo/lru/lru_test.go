package lru_test

import (
	"fmt"
	"goks/algo/lru"
	"goks/cache"
	"testing"
)

func TestSet(t *testing.T)  {
	repo := lru.NewLruCache(5)
	topic := &cache.Topic{
		Key:   "key-1",
		Value: "topic with key-1",
	}

	err := repo.Set(topic)
	if err != nil {
		t.Fatalf("expected %v, actual %v", nil, err)
	}

	// Check if the item is exists
	item, err := repo.Peek("key-1")

	if err != nil {
		t.Fatalf("expected %v, actual %v", nil, err)
	}

	if item == nil {
		t.Fatalf("expected %v, actual %v", "Hello World", err)
	}
}

func TestSetMultiple(t *testing.T) {
	repo := lru.NewLruCache(5)

	for i := 1; i <= 5; i++ {
		topic := &cache.Topic{
			Key:   fmt.Sprintf("key-%d", i),
			Value: "topic with key-1",
		}
		err := repo.Set(topic)
		if err != nil {
			t.Fatalf("expected %v, actual %v", nil, err)
		}
	}

	for i := 2; i <= 5; i++ {
		item, err := repo.Peek(fmt.Sprintf("key-%d", i))
		if err != nil {
			t.Fatalf("expected %v, actual %v", nil, err)
		}

		if item == nil {
			t.Fatalf("expected %v, actual %v", i, err)
		}
	}

}