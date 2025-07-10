package lfu_test

import (
	"testing"

	"github.com/hgsgtk/go-snippets/lfu"
)

func TestLFUCache_PutUpdate(t *testing.T) {
	c := lfu.NewLFUCache(2)
	
	// Put initial value
	c.Put(1, 10)
	
	// Update the same key
	c.Put(1, 20)
	
	// Verify the value was updated
	if c.Get(1) != 20 {
		t.Errorf("Expected value 20, got %d", c.Get(1))
	}
}

func TestLFUCache_Capacity(t *testing.T) {
	c := lfu.NewLFUCache(2)
	
	// Add two items
	c.Put(1, 1)
	c.Put(2, 2)
	
	// Add third item - should evict least frequently used
	c.Put(3, 3)
	
	// Key 1 should be evicted (first in, first out for same frequency)
	if c.Get(1) != -1 {
		t.Errorf("Expected key 1 to be evicted, but got value %d", c.Get(1))
	}
	
	// Keys 2 and 3 should still be there
	if c.Get(2) != 2 {
		t.Errorf("Expected value 2, got %d", c.Get(2))
	}
	if c.Get(3) != 3 {
		t.Errorf("Expected value 3, got %d", c.Get(3))
	}
}

func TestLFUCache_LFUEviction(t *testing.T) {
	c := lfu.NewLFUCache(2)
	
	// Add two items
	c.Put(1, 1)
	c.Put(2, 2)
	
	// Access key 2 to increase its frequency
	c.Get(2)
	
	// Add third item - should evict key 1 (lower frequency)
	c.Put(3, 3)
	
	// Key 1 should be evicted (lower frequency)
	if c.Get(1) != -1 {
		t.Errorf("Expected key 1 to be evicted, but got value %d", c.Get(1))
	}
	
	// Keys 2 and 3 should still be there
	if c.Get(2) != 2 {
		t.Errorf("Expected value 2, got %d", c.Get(2))
	}
	if c.Get(3) != 3 {
		t.Errorf("Expected value 3, got %d", c.Get(3))
	}
}

func TestLFUCache_LRUTieBreak(t *testing.T) {
	c := lfu.NewLFUCache(2)
	
	// Add two items with same frequency
	c.Put(1, 1)
	c.Put(2, 2)
	
	// Both have frequency 1, so adding a third should evict the first one (LRU)
	c.Put(3, 3)
	
	// Key 1 should be evicted (first in, first out)
	if c.Get(1) != -1 {
		t.Errorf("Expected key 1 to be evicted, but got value %d", c.Get(1))
	}
	
	// Keys 2 and 3 should still be there
	if c.Get(2) != 2 {
		t.Errorf("Expected value 2, got %d", c.Get(2))
	}
	if c.Get(3) != 3 {
		t.Errorf("Expected value 3, got %d", c.Get(3))
	}
}	

func TestLFUCache_Sequence(t *testing.T) {
	c := lfu.NewLFUCache(2)
	
	// Test sequence: ["LFUCache","put","put","get","put","get","get","put","get","get","get"]
	// [[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]
	
	// put(1,1)
	c.Put(1, 1)
	
	// put(2,2)
	c.Put(2, 2)
	
	// get(1) - should return 1
	if val := c.Get(1); val != 1 {
		t.Errorf("Expected Get(1) to return 1, got %d", val)
	}
	
	// put(3,3) - should evict key 2 (key 1 has higher frequency now)
	c.Put(3, 3)
	
	// get(2) - should return -1 (evicted)
	if val := c.Get(2); val != -1 {
		t.Errorf("Expected Get(2) to return -1 (evicted), got %d", val)
	}
	
	// get(3) - should return 3
	if val := c.Get(3); val != 3 {
		t.Errorf("Expected Get(3) to return 3, got %d", val)
	}
	
	// put(4,4) - should evict key 1 (key 3 has higher frequency now)
	c.Put(4, 4)
	
	// get(1) - should return -1 (evicted)
	if val := c.Get(1); val != -1 {
		t.Errorf("Expected Get(1) to return -1 (evicted), got %d", val)
	}
	
	// get(3) - should return 3
	if val := c.Get(3); val != 3 {
		t.Errorf("Expected Get(3) to return 3, got %d", val)
	}
	
	// get(4) - should return 4
	if val := c.Get(4); val != 4 {
		t.Errorf("Expected Get(4) to return 4, got %d", val)
	}
}

func BenchmarkLFUCache_Get(b *testing.B) {
	c := lfu.NewLFUCache(1000)
	
	// Pre-populate cache
	for i := 0; i < 1000; i++ {
		c.Put(i, i*2)
	}
	
	// Access keys multiple times to create different frequencies
	for i := 0; i < 1000; i++ {
		c.Get(i % 100) // Some keys will be accessed more frequently
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Get(i % 1000)
	}
}

func BenchmarkLFUCache_Put(b *testing.B) {
	c := lfu.NewLFUCache(1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Put(i%1000, i)
	}
}

func BenchmarkLFUCache_Mixed(b *testing.B) {
	c := lfu.NewLFUCache(1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			c.Put(i%1000, i)
		} else {
			c.Get(i % 1000)
		}
	}
}	
