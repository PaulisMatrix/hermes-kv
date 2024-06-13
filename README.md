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

**Current tests:**
```
❯ go test ./... -v -cover -race -coverprofile cover.out    
?   	hermeskv/examples	[no test files]
=== RUN   TestDLLSet
--- PASS: TestDLLSet (0.00s)
=== RUN   TestDLLDeleteHead
--- PASS: TestDLLDeleteHead (0.00s)
=== RUN   TestDLLDeleteTail
--- PASS: TestDLLDeleteTail (0.00s)
=== RUN   TestDLLDeleteNode
--- PASS: TestDLLDeleteNode (0.00s)
=== RUN   TestSetKV
--- PASS: TestSetKV (0.00s)
=== RUN   TestGetKV
--- PASS: TestGetKV (0.00s)
=== RUN   TestDeleteKV
--- PASS: TestDeleteKV (0.00s)
=== RUN   TestKVCapBreach
--- PASS: TestKVCapBreach (0.00s)
=== RUN   TestZeroCapKV
--- PASS: TestZeroCapKV (0.00s)
=== RUN   TestKVSetRacer
--- PASS: TestKVSetRacer (2.00s)
=== RUN   TestKVDeleteRacer
--- PASS: TestKVDeleteRacer (2.00s)
PASS
coverage: 90.3% of statements
ok  	hermeskv	5.640s	coverage: 90.3% of statements

-- generate and open the cover profile
go tool cover -html cover.out -o cover.html
open cover.html
```