package main

import (
	"fmt"
	"goks"
	"log"
)

func main() {
	client := goks.NewClient(goks.NewCacheOptions().SetMaxSizeItem(5))
	//insert cache key value
	for i := 1; i <= 5; i++ {
		err := client.Set(fmt.Sprintf("key-%d", i), fmt.Sprintf("value-%d", i))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Get Cache key
	keys, err := client.GetKeys()
	if err != nil {
		log.Fatal(err)
	}

	// Peek Cache by key
	for _, key := range keys {
		res, _ := client.PeekByKey(key)
		fmt.Println("Cache value list =>", res)
	}

	//Get cache by key
	val, err := client.Get("key-1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value with given key =>", val)

	// Insert new exceed capacity
	err = client.Set("key-6", "value-6")
	if err != nil {
		log.Fatal(err)
	}

	// Get new Cache key
	newKeys, err := client.GetKeys()
	if err != nil {
		log.Fatal(err)
	}

	// Peek Cache by key
	for _, key := range newKeys {
		res, _ := client.PeekByKey(key)
		fmt.Println("Cache value list =>", res)

	}
}
