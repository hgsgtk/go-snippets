package hashmap

import "testing"

func TestHashmap(t *testing.T) {
	hashmap := Constructor()
	hashmap.Put(1, 1)
	hashmap.Put(2, 2)
	got := hashmap.Get(1)
	if got != 1 {
		t.Errorf("Get(1) = %d; want 1", got)
	}

	got = hashmap.Get(3)
	if got != -1 {
		t.Errorf("Get(3) = %d; want -1", got)
	}

	hashmap.Put(2, 1)
	got = hashmap.Get(2)
	if got != 1 {
		t.Errorf("Get(2) = %d; want 1", got)
	}

	hashmap.Remove(2)
	got = hashmap.Get(2)
	if got != -1 {
		t.Errorf("Get(2) = %d; want 1", got)
	}
}