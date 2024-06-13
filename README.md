**Hermes KV**:

Why **Hermes**?

Cause Hermes is the Greek god of commerce, communication, and the messenger of the gods. This name suggests speed and efficiency in data retrieval and storage. ^^thanks to gpt for the name.

**Basic KV Stores:**

1.  A simple in-memory cache system for storing key-value pairs.
2.  Supported methods:
    * `Set(key string, value interface{}) error` 
    * `Get(key string) (interface{}, error)`
    * `Delete(key string) error`

**FIFO Eviction Policy:**

1.  Extending the basic cache to incorporate the First-In-First-Out (FIFO) eviction policy.
2.  Implement the cache to evict the oldest item first when the cache reaches its predefined capacity.
3.  map for O(1) key access
4.  Doubly linked list for O(1) eviction. Deleting the head which is the first key stored so eligible for eviction.
    Adding a new node is also O(1) since we are maintaining prev pointer in a doubly linked list.

**KV store with transactions:**


**References on In-memory cache:**
  * https://github.com/patrickmn/go-cache/
  * 

**Misc:**
```
-- generate and open the cover profile

❯ go test ./... -v -cover -race -coverprofile cover.out
❯ go tool cover -html cover.out -o cover.html
❯ open cover.html
```