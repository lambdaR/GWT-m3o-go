package main

import (
	"fmt"
	"os"

	"go.m3o.com/cache"
)

// Delete a value from the cache
func main() {
	cacheService := cache.NewCacheService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cacheService.Delete(&cache.DeleteRequest{
		Key: "foo",
	})
	fmt.Println(rsp, err)
}
