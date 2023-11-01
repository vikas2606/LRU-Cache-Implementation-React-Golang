package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	lru "github.com/hashicorp/golang-lru"
)

type CacheWithMutex struct {
	cache *lru.Cache
	mu    sync.Mutex
}

var cacheWithMutex *CacheWithMutex

func main() {

	lruCache, _ := lru.New(5)
	cacheWithMutex = &CacheWithMutex{cache: lruCache}

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	r.GET("/get/:key", getFromCache)
	r.POST("/set", setInCache)

	r.Run(":8080")
}

func getFromCache(c *gin.Context) {
	key := c.Param("key")
	cacheWithMutex.mu.Lock()
	defer cacheWithMutex.mu.Unlock()

	if val, ok := cacheWithMutex.cache.Get(key); ok {
		c.JSON(http.StatusOK, val)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Key not found in cache"})
	}
}

func setInCache(c *gin.Context) {
	var payload struct {
		Key        string `json:"key"`
		Value      string `json:"value"`
		Expiration string `json:"expiration"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cacheWithMutex.mu.Lock()
	defer cacheWithMutex.mu.Unlock()

	expiration, err := strconv.ParseInt(payload.Expiration, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration value"})
		return
	}

	cacheWithMutex.cache.Add(payload.Key, payload.Value)

	go removeAfterExpiration(payload.Key, time.Duration(expiration)*time.Second)

	c.JSON(http.StatusOK, gin.H{"message": "Value added to cache"})

}

func removeAfterExpiration(key string, duration time.Duration) {
	<-time.After(duration)
	cacheWithMutex.mu.Lock()
	defer cacheWithMutex.mu.Unlock()
	cacheWithMutex.cache.Remove(key)
}
