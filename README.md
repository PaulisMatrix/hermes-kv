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

* Basic assumptions:
  1. Only one transaction can be issued at a time. Assume no concurrent transactions.
  2. `commit()` or `rollback` can be called only once in a transaction. You can specify n number of operations in between enclosed in between `begin()` and `end()`
  3. Let's not consider nested transactions for now.

* Implementation details:
  1. Maintain a `tempKVMap` recording all the db changes. 
  2. On `commit()`, iterate over all k,v pairs `O(N)` and call `Set(k, v)` of the underlying main KV store to record the final changes.
  3. On `rollback()`, delete the `tempKVMap` altogether. 
  4. Use flag `isTxActive` to check which Map to refer whenever any method is called.
  5. For `delete()`, maintain a `isTombStone` boolean to mark a kv pair as deleted in localState. On `commit()`, call the actual `delete()` method to delete it from the globalState. On `rollback()`, nothing needs to be done, just clear the localState.

* References:
  1. https://www.freecodecamp.org/news/design-a-key-value-store-in-go/
  2. 
  
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