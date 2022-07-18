Go Cache helps you to made simple cache based on map with Set\Get\Delete functions.
Example
```
package main

import (
"fmt"
"github.com/VadimGossip/simpleCache"
)

func main() {
cache := simpleCache.NewCache()

	cache.Set("userId", 42)
	userId, err := cache.Get("userId")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userId)

	cache.Delete("userId")
	userId, err = cache.Get("userId")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userId)
}
```
