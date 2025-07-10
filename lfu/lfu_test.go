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
