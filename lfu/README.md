# LFU Cache library

## Requirements

### What is LFU?

* LFU (Least Frequently Used) is a cache eviction policy that evicts the least frequently used items first.

### How to use?

* Initialize the object with the capacity of the cache

```go
c := NewLFUCache(10)
```

* Put a key-value pair in the cache

```go
c.Put(1, 1)
```

* Get a value from the cache

```go
c.Get(1)
```

### Functionality

- [x] Initialize the cache object
- [x] Put a value
    * [x] Input: key, value (int, int)
    * Process:
        * [x] If present, update the value
        * [x] If not, add the key-value pair to the cache
        * [x] When the cache is full, evict the least frequently used item
            * [x] If there are multiple items with the same frequency, evict the least recently used item
    * Output: void
- [x] Get a value
    * [x] Input: key (int)
    * [x] Output: 
        * [x]   If exists, returns value (int)
        * [x] If not, returns -1
