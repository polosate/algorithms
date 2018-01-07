package _6_heap

import "testing"

func TestAscInsert(t *testing.T) {
	h := NewAscHeap(10)
	h.insert(2)
	h.insert(8)
	h.insert(10)
	h.insert(7)
	h.insert(6)
	h.insert(3)
	h.insert(5)
	h.insert(4)
	h.insert(9)

	h.displayHeap()
}

func TestAscRemove(t *testing.T) {
	h := NewAscHeap(10)
	h.insert(2)
	h.insert(8)
	h.insert(10)
	h.insert(7)
	h.insert(6)
	h.insert(3)
	h.insert(5)
	h.insert(4)
	h.insert(9)

	h.displayHeap()

	r, _ := h.remove()
	if r.getKey() != 2 {
		t.Error()
	}
	h.displayHeap()
}
