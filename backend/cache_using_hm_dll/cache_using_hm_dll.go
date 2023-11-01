package cacheusinghmdll

import (
	"net/http"
	"strconv"
	"sync"

	"time"

	"github.com/gin-gonic/gin"
)

type DLLNode struct {
	Key, Value string
	Prev, Next *DLLNode
}

type CacheWithHLDLL struct {
	CacheMap   map[string]*DLLNode
	Head, Tail *DLLNode
	MaxSize    int
	mu         sync.Mutex
}

var CacheWithHL_DLL *CacheWithHLDLL

func ListKeysDLL(c *gin.Context) {
	CacheWithHL_DLL.mu.Lock()
	defer CacheWithHL_DLL.mu.Unlock()

	keys := make([]string, 0, len(CacheWithHL_DLL.CacheMap))
	for k := range CacheWithHL_DLL.CacheMap {
		keys = append(keys, k)
	}
	c.JSON(http.StatusOK, keys)
}

func GetFromCacheDLL(c *gin.Context) {
	key := c.Param("key")
	CacheWithHL_DLL.mu.Lock()
	defer CacheWithHL_DLL.mu.Unlock()

	if node, ok := CacheWithHL_DLL.CacheMap[key]; ok {
		// Move the accessed node to the front of the DLL
		CacheWithHL_DLL.moveToFront(node)
		c.JSON(http.StatusOK, node.Value)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Key not found in cache"})
	}
}

func SetInCacheDLL(c *gin.Context) {
	var payload struct {
		Key        string `json:"key"`
		Value      string `json:"value"`
		Expiration string `json:"expiration"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CacheWithHL_DLL.mu.Lock()
	defer CacheWithHL_DLL.mu.Unlock()

	expiration, err := strconv.ParseInt(payload.Expiration, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration value"})
		return
	}

	node := &DLLNode{
		Key:   payload.Key,
		Value: payload.Value,
	}
	CacheWithHL_DLL.addToFront(node)

	go removeAfterExpirationDLL(node, time.Duration(expiration)*time.Second)

	c.JSON(http.StatusOK, gin.H{"message": "Value added to cache"})
}

func removeAfterExpirationDLL(node *DLLNode, duration time.Duration) {
	<-time.After(duration)
	CacheWithHL_DLL.mu.Lock()
	defer CacheWithHL_DLL.mu.Unlock()
	CacheWithHL_DLL.removeNode(node)
}

func (cache *CacheWithHLDLL) addToFront(node *DLLNode) {
	if cache.Head == nil {
		cache.Head = node
		cache.Tail = node
	} else {
		node.Next = cache.Head
		cache.Head.Prev = node
		cache.Head = node
	}
	cache.CacheMap[node.Key] = node
	if len(cache.CacheMap) > cache.MaxSize {
		cache.removeLast()
	}

}

func (cache *CacheWithHLDLL) moveToFront(node *DLLNode) {
	if node != cache.Head {
		if node == cache.Tail {
			cache.Tail = node.Prev
		}
		node.Prev.Next = node.Next
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}
		node.Next = cache.Head
		node.Prev = nil
		cache.Head.Prev = node
		cache.Head = node
	}
}

func (cache *CacheWithHLDLL) removeNode(node *DLLNode) {
	delete(cache.CacheMap, node.Key)
	if node == cache.Head {
		cache.Head = node.Next
		if cache.Head != nil {
			cache.Head.Prev = nil
		}
	} else if node == cache.Tail {
		cache.Tail = node.Prev
		if cache.Tail != nil {
			cache.Tail.Next = nil
		}
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
}

func (cache *CacheWithHLDLL) removeLast() {
	if cache.Tail != nil {
		delete(cache.CacheMap, cache.Tail.Key)
		if cache.Tail == cache.Head {
			cache.Tail = nil
			cache.Head = nil
		} else {
			cache.Tail = cache.Tail.Prev
			cache.Tail.Next = nil
		}
	}
}
