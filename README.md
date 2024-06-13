**Implement a Basic Cache:**

1. Design a simple in-memory cache system for storing key-value pairs.
2. Implement **`Set(key string, value interface{})`** and **`Get(key string) interface{}`** methods.

**Implement FIFO Eviction Policy:**

1. Extend the basic cache to incorporate the First-In-First-Out (FIFO) eviction policy.
2. Implement the cache to evict the oldest item first when the cache reaches its predefined capacity.

map for O(1) key access
DLL for O(1) eviction. Deleting the head which is the first key stored so eligible for eviction.
Adding a new node is also O(1) since we are maintaining prev pointer in a doubly linked list.

```
❯ go test ./... -v -cover -race -coverprofile cover.out    
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
capacity breached, deleting head node...
--- PASS: TestKVCapBreach (0.00s)
=== RUN   TestZeroCapKV
--- PASS: TestZeroCapKV (0.00s)
=== RUN   TestKVRacer
--- PASS: TestKVRacer (2.00s)
PASS
coverage: 79.5% of statements
ok  	ravenmail	3.183s	coverage: 79.5% of statements

-- generate and open the cover profile
go tool cover -html cover.out -o cover.html
open cover.html
```

**References on In-memory cache:**
  * https://github.com/patrickmn/go-cache/
  * 