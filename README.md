
# LRU Cache Implementation (React+Golang)

This project provides two implementations of an LRU (Least Recently Used) cache in Go. You can choose between the LRU cache and the DLL-based cache by specifying a keyword in the `main` function.


## Tech Stack

 **Client:** React JS

**Server:** Golang



## Table of Contents
- [Overview](#overview)
- [Cache Implementations](#cache-implementations)
- [Getting Started](#getting-started)
- [API Routes](#api-routes)
- [Switching Between Caches](#switching-between-caches)
- [Customization](#customization)
## Overview

Both cache implementations share the following features:

- Cache Operations: Get, Set, and List keys.
- Thread-Safe Access: Synchronized access using mutexes.
- Expiration Handling: Automatic removal of items after a specified duration.
- Selectable Cache Type: Choose between LRU and DLL-based caches.
- Dynamic Cache Size: Set the maximum size of the cache.
- HTTP Routes: API routes for cache management.
- Access Order: Maintain access order for the DLL-based cache, while LRU cache follows LRU order.
- Concurrency Control: Handle concurrent access gracefully.
- Automatic Item Removal: Remove expired items to free up cache space.
- CORS Configuration: Control access using Cross-Origin Resource Sharing (CORS).
- Modularity and Extensibility: Add custom features as needed.
## Cache Implementations

### LRU Cache
- Uses the LRU algorithm to manage cached items.
- Implemented with the `github.com/hashicorp/golang-lru` library.
- Provides routes for cache management.

### DLL-based Cache
- Uses a hashmap and a doubly linked list to manage cached items.
- Provides routes for cache management.
- Maintains access order in the doubly linked list.
## Getting Started


Clone the project

```bash
  git clone https://github.com/vikas2606/LRU-Cache-Implementation-React-Golang.git
```

### Implement Backend
From root folder
```bash
  cd backend
```

Install dependencies

```bash
  go mod tidy

```


Start the server

```bash
  go run main.go
```


### Implement Frontend

From root folder
```bash
  cd frontend
```

Install dependencies

```bash
  npm install

```


Start the Client

```bash
  npm start
```
## Related

Here are some related projects

[Awesome README](https://github.com/matiassingers/awesome-readme)


## API Routes

The following API routes are available for cache management:

- __GET /get__ : List all keys in the cache.
- __GET /get/:key__ : Retrieve a value from the cache based on the provided key.
- __POST /set__: Add a key-value pair to the cache with an optional expiration time.
## Switching Between Caches

You can select between the LRU cache and the DLL-based cache by changing a keyword in the `main.go` file:
```bash
// Change this keyword to select the cache implementation
cacheType := "lru" // Change this to "dll" to use the DLL-based cache

```
## Customization
You can customize the cache size, CORS settings, and other aspects of the application by modifying the appropriate code sections in the `main.go` and `router.go` files.