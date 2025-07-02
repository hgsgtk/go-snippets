package bloomfilter

import (
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)

	items := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon"}

	for _, item := range items {
		bf.Add(item)
	}

	for _, item := range items {
		if !bf.Contains(item) {
			t.Errorf("Item %s should be in the bloom filter", item)
		}
	}

	if bf.Contains("no-exist") {
		t.Errorf("Item %s should not be in the bloom filter", "no-exist")
	}
}