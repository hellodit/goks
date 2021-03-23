package main

import (
	"fmt"
	"goks"
	"log"
)

func main(){
	client := goks.NewClient(goks.NewCacheOptions().SetMaxSizeItem(3))
	err := client.Set("key-1","value brot")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Set("key-2","value brot")
	if err != nil {
		log.Fatal(err)
	}
	val, err := client.Get("key-1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}
