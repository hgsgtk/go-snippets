package singlylinkedlist

import (
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2)
	got := linkedList.Get(1)
	if got != 2 {
		t.Errorf("Get(1) = %d; want 2", got)
	}
	linkedList.DeleteAtIndex(1)
	got = linkedList.Get(1)
	if got != 3 {
		t.Errorf("Get(1) = %d; want 3", got)
	}
}

func TestComplexLinkedListOperations(t *testing.T) {
	linkedList := Constructor()
		
	linkedList.AddAtHead(7)
	linkedList.AddAtHead(2)
	linkedList.AddAtHead(1)
	linkedList.AddAtIndex(3, 0)
	linkedList.DeleteAtIndex(2)
	linkedList.AddAtHead(6)
	linkedList.AddAtTail(4)
	
	// Test get(4) - should return value at index 4
	got := linkedList.Get(4)
	if got != 4 {
		t.Errorf("Get(4) = %d; want 4", got)
	}
	
	linkedList.AddAtHead(4)
	linkedList.AddAtIndex(5, 0)
	linkedList.AddAtHead(6)
		
	// Verify final state: [6, 4, 6, 1, 2, 0, 0, 4]
	expected := []int{6, 4, 6, 1, 2, 0, 0, 4}
	for i, want := range expected {
		got = linkedList.Get(i)
		if got != want {
			t.Errorf("Get(%d) = %d; want %d", i, got, want)
		}
	}
}

