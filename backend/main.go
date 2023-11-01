// main.go
package main

import (
	"lru-cache/cache"
	cacheusinghmdll "lru-cache/cache_using_hm_dll"
	"lru-cache/router"

	lru "github.com/hashicorp/golang-lru"
)

func main() {

	cacheType := "dll"

	if cacheType == "lru" {
		lruCache, _ := lru.New(1024) // Change the number of keys
		cache.CacheAlongMutex = &cache.CacheWithMutex{Cache: lruCache}
	} else if cacheType == "dll" {
		cacheWithHLDLL := &cacheusinghmdll.CacheWithHLDLL{
			CacheMap: make(map[string]*cacheusinghmdll.DLLNode),
			MaxSize:  1024, // Change the number of keys
		}
		cacheusinghmdll.CacheWithHL_DLL = cacheWithHLDLL
	}

	r := router.SetupRouter(cacheType)
	r.Run(":8080")
}
