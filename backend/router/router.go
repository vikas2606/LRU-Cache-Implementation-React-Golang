package router

import (
	"lru-cache/cache"
	cacheusinghmdll "lru-cache/cache_using_hm_dll"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cacheType string) *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	if cacheType == "lru" {
		r.GET("/get", cache.ListKeys)
		r.GET("/get/:key", cache.GetFromCache)
		r.POST("/set", cache.SetInCache)
	} else if cacheType == "dll" {
		// Define routes using the DLL-based cache
		r.GET("/get", cacheusinghmdll.ListKeysDLL)
		r.GET("/get/:key", cacheusinghmdll.GetFromCacheDLL)
		r.POST("/set", cacheusinghmdll.SetInCacheDLL)
	}
	return r
}
