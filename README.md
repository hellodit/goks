## Goks
Go in-memory cache algorithm currently full support LRU Algorithm  

#### Installation
`commingsoon`

#### Usage

```go
import (
    "fmt"
    "goks"
    "log"
)

// create cache client to communicate with module with max item 5 
client := goks.NewClient(goks.NewCacheOptions().SetMaxSizeItem(5))

// insert key:value pair to cache
err := client.Set("key-1","value-1")
if err != nil {
    log.Fatal(err)
}

// get cache value by key 
val, err := client.Get("key-1")
if err != nil {
    log.Fatal(err)
}

fmt.Println("Value with given key =>", val)
```
For further usage, please refer the `example` folder

### Note
- This Package only for research purpose only 
- Don't use for production purpose 

### How to contribute?
Just create issue then we discuss 