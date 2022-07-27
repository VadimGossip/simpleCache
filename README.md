Go Cache helps you to made simple cache based on map with Set\Get\Delete functions.
Race consdition protected.

Example

```
package main

import (
	"fmt"
	"github.com/VadimGossip/simpleCache"
	"time"
)

func main() {
    cache := simpleCache.NewCache()

    cache.Set("userId", 42, time.Second * 5)
    userId, err := cache.Get("userId")
    if err != nil {
    	fmt.Println(err)
    }
    fmt.Println(userId) // will print 42

    time.Sleep(time.Second * 4)

    userId, err = cache.Get("userId")
    if err != nil {
    	fmt.Println(err)
    }
    fmt.Println(userId) // will print 42 item still alive

    time.Sleep(time.Second * 2)
    userId, err = cache.Get("userId")
    if err != nil {
    	fmt.Println(err) //item not exists
    }
}

```
