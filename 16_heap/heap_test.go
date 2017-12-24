package _6_heap

import "testing"

func TestNewHeap(t *testing.T) {
	h := NewHeap(10)
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
