package _6_heap

import "testing"

func TestInsert(t *testing.T) {
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

func TestRemove(t *testing.T) {
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

	r, _ := h.remove()
	if r.getKey() != 10 {
		t.Error()
	}
	h.displayHeap()
}

func TestRestoreHeap(t *testing.T) {
	h := NewHeap(10)
	h.insert(2)
	h.insert(8)
	h.toss(7)
	h.toss(6)
	h.toss(3)
	h.toss(10)
	h.restoreHeap()
	r, _ := h.remove()
	if r.getKey() != 10 {
		t.Error()
	}
	r, _ = h.remove()
	if r.getKey() != 8 {
		t.Error()
	}
	r, _ = h.remove()
	if r.getKey() != 7 {
		t.Error()
	}
	r, _ = h.remove()
	if r.getKey() != 6 {
		t.Error()
	}
	r, _ = h.remove()
	if r.getKey() != 3 {
		t.Error()
	}
	r, _ = h.remove()
	if r.getKey() != 2 {
		t.Error()
	}
	_, err := h.remove()
	if err == nil {
		t.Error()
	}
}
