package cache

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	lru "github.com/hashicorp/golang-lru"
)

type CacheWithMutex struct {
	Cache *lru.Cache
	mu    sync.Mutex
}

var CacheAlongMutex *CacheWithMutex

func ListKeys(c *gin.Context) {
	CacheAlongMutex.mu.Lock()
	defer CacheAlongMutex.mu.Unlock()

	keys := CacheAlongMutex.Cache.Keys()
	c.JSON(http.StatusOK, keys)

}

func GetFromCache(c *gin.Context) {
	key := c.Param("key")
	CacheAlongMutex.mu.Lock()
	defer CacheAlongMutex.mu.Unlock()

	if val, ok := CacheAlongMutex.Cache.Get(key); ok {
		c.JSON(http.StatusOK, val)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Key not found in cache"})
	}
}

func SetInCache(c *gin.Context) {
	var payload struct {
		Key        string `json:"key"`
		Value      string `json:"value"`
		Expiration string `json:"expiration"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CacheAlongMutex.mu.Lock()
	defer CacheAlongMutex.mu.Unlock()

	expiration, err := strconv.ParseInt(payload.Expiration, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration value"})
		return
	}

	CacheAlongMutex.Cache.Add(payload.Key, payload.Value)

	go removeAfterExpiration(payload.Key, time.Duration(expiration)*time.Second)

	c.JSON(http.StatusOK, gin.H{"message": "Value added to cache"})

}

func removeAfterExpiration(key string, duration time.Duration) {
	<-time.After(duration)
	CacheAlongMutex.mu.Lock()
	defer CacheAlongMutex.mu.Unlock()
	CacheAlongMutex.Cache.Remove(key)
}
