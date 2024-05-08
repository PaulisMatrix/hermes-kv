**Implement a Basic Cache:**

1. Design a simple in-memory cache system for storing key-value pairs.
2. Implement **`Set(key string, value interface{})`** and **`Get(key string) interface{}`** methods.

**Implement FIFO Eviction Policy:**

1. Extend the basic cache to incorporate the First-In-First-Out (FIFO) eviction policy.
2. Implement the cache to evict the oldest item first when the cache reaches its predefined capacity.

map for O(1) key access
DLL for O(1) eviction. Deleting the head which is the first key stored so eligible for eviction.

