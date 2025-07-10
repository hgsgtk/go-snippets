package lfu

import "container/list"

type LFUCache struct {
	capacity int
	cache    map[int]int        // key -> value
	freq     map[int]int        // key -> frequency count
	freqMap  map[int]*list.List // frequency -> list of keys (LRU order)
	keyNodes map[int]*list.Element // key -> list element for O(1) removal
	minFreq  int                // minimum frequency in the cache
	nonEmptyFreqs map[int]bool  // track non-empty frequency levels
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		cache:    make(map[int]int),
		freq:     make(map[int]int),
		freqMap:  make(map[int]*list.List),
		keyNodes: make(map[int]*list.Element),
		minFreq:  0,
		nonEmptyFreqs: make(map[int]bool),
	}
}

func (c *LFUCache) Put(key int, value int) {
	// If key already exists, update value and increment frequency
	if _, exists := c.cache[key]; exists {
		c.cache[key] = value
		c.incrementFrequency(key)
		return
	}

	// If cache is at capacity, evict least frequently used item
	if len(c.cache) >= c.capacity {
		c.evictLFU()
	}

	// Add new key-value pair
	c.cache[key] = value
	c.freq[key] = 1
	
	// Add to frequency map
	if c.freqMap[1] == nil {
		c.freqMap[1] = list.New()
	}
	element := c.freqMap[1].PushBack(key)
	c.keyNodes[key] = element
	c.nonEmptyFreqs[1] = true
	
	// Update minimum frequency
	c.minFreq = 1
}

func (c *LFUCache) Get(key int) int {
	// Check if key exists in cache
	if value, exists := c.cache[key]; exists {
		// Increment frequency when accessed
		c.incrementFrequency(key)
		return value
	}
	
	// Key doesn't exist
	return -1
}

// incrementFrequency increases the frequency of a key and updates the frequency maps
func (c *LFUCache) incrementFrequency(key int) {
	freq := c.freq[key]
	
	// Remove from current frequency list using O(1) removal
	if freqList := c.freqMap[freq]; freqList != nil {
		if element, exists := c.keyNodes[key]; exists {
			freqList.Remove(element)
			delete(c.keyNodes, key)
		}
		
		// If this frequency list is now empty, remove it
		if freqList.Len() == 0 {
			delete(c.freqMap, freq)
			delete(c.nonEmptyFreqs, freq)
			// If this was the minimum frequency, update minFreq
			if freq == c.minFreq {
				c.updateMinFreq()
			}
		}
	}
	
	// Increment frequency
	c.freq[key]++
	newFreq := c.freq[key]
	
	// Add to new frequency list
	if c.freqMap[newFreq] == nil {
		c.freqMap[newFreq] = list.New()
	}
	element := c.freqMap[newFreq].PushBack(key)
	c.keyNodes[key] = element
	c.nonEmptyFreqs[newFreq] = true
	
	// Update minFreq if needed (after adding the new frequency)
	if c.minFreq == 0 || newFreq < c.minFreq {
		c.minFreq = newFreq
	}
}

// updateMinFreq finds the next minimum frequency in O(1) average time
func (c *LFUCache) updateMinFreq() {
	// Find the next minimum frequency from non-empty frequencies
	c.minFreq = 0
	for freq := range c.nonEmptyFreqs {
		if c.minFreq == 0 || freq < c.minFreq {
			c.minFreq = freq
		}
	}
}

// evictLFU removes the least frequently used item from the cache
func (c *LFUCache) evictLFU() {
	if c.freqMap[c.minFreq] == nil || c.freqMap[c.minFreq].Len() == 0 {
		return
	}
	
	// Remove the least recently used item from the minimum frequency list
	lruItem := c.freqMap[c.minFreq].Front()
	keyToEvict := lruItem.Value.(int)
	c.freqMap[c.minFreq].Remove(lruItem)
	delete(c.keyNodes, keyToEvict)
	
	// Remove from cache and frequency maps
	delete(c.cache, keyToEvict)
	delete(c.freq, keyToEvict)
	
	// If this frequency list is now empty, remove it and update minFreq
	if c.freqMap[c.minFreq].Len() == 0 {
		delete(c.freqMap, c.minFreq)
		delete(c.nonEmptyFreqs, c.minFreq)
		// Find the next minimum frequency
		c.updateMinFreq()
	}
}

// Debug methods for testing
func (c *LFUCache) GetCacheSize() int {
	return len(c.cache)
}

func (c *LFUCache) GetCapacity() int {
	return c.capacity
}

func (c *LFUCache) GetValue(key int) (int, bool) {
	val, exists := c.cache[key]
	return val, exists
}

func (c *LFUCache) GetFrequency(key int) (int, bool) {
	freq, exists := c.freq[key]
	return freq, exists
}

func (c *LFUCache) GetMinFreq() int {
	return c.minFreq
}
