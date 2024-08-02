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

**LRU Eviction Policy:**
1. Use the same DLL maintained for FIFO.
2. Whenever a key is referenced/updated, delete it from the head and insert it at the tail. 
3. The Least Recently Used key is always at the head of the DLL.
4. Deleting and Inserting is O(1) since we are maintaining respective head and tail pointers.   
5. Check [this branch](https://github.com/PaulisMatrix/hermes-kv/tree/LRU) for the implementation.

**KV store with transactions:**

Basic assumptions:
1. `commit()` or `rollback` can be called only once in a transaction. You can specify n number of operations in between enclosed in between `begin()` and `end()`
2. Let's not consider nested transactions for now.

Implementation details:
1. Maintain a `tempKVMap` recording all the db changes. 
2. On `commit()`, iterate over all k,v pairs `O(N)` and call `Set(k, v)` of the underlying main KV store to record the final changes.
3. On `rollback()`, delete the `tempKVMap` altogether. 
4. Use flag `isTxActive` to check which Map to refer whenever any method is called.
5. For `delete()`, maintain a `isTombStone` boolean to mark a kv pair as deleted in localState. On `commit()`, call the actual `delete()` method to delete it from the globalState. On `rollback()`, nothing needs to be done, just clear the localState.

**Concurrent Transactions:**
Snapshot isolation with LWW to resolve the conflicts - 
1. Whenever a tx starts, it gets its own copy of the `globalStateMap`. Multiple txs can modify their own `localStateMap`. This is concurrent safe cause each have their own copy. 
2. On `commit()`, if one or more tx has modified the same value then the latest update will be committed to the global state i.e Last Write Wins. 
3. Each update to a key is timestamped so we will pick the latest timestamp for conflict resolution.
4. This works well but guarantee upto date updates due to clock skews. 

Multiple readers, Single writer - 
1. Concurrent txs. Similar to sqlite. Single writer, multiple readers. Rest of the writers will be blocked until the current active one succeeds. Specify a timeout like [busy_timeout](https://sqlite.org/c3ref/busy_timeout.html) pragma of sqlite, to specify how long the blocked writers should wait.
2. 
  
**Similar works/references**
* https://github.com/patrickmn/go-cache/
* https://www.freecodecamp.org/news/design-a-key-value-store-in-go/
* Implementating different consistency/isolation levels: 
    * Phil's(EDB-Postgres) (blog post)[https://notes.eatonphil.com/2024-05-16-mvcc.html]
    * Jesse's(MongoDB) (blog post)[https://emptysqua.re/blog/pycon-2023-consistency-isolation/]

**Misc:**
```
-- generate and open the cover profile

❯ go test ./... -v -cover -race -coverprofile cover.out
❯ go tool cover -html cover.out -o cover.html
❯ open cover.html
```