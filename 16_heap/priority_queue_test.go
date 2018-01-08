package _6_heap

import "testing"

func TestPriorityQueue(t *testing.T) {
	q := newPriorityQ(10)
	if !q.isEmpty() {
		t.Error()
	}
	q.insert(5)
	q.insert(10)
	q.insert(7)
	q.insert(2)
	q.insert(6)

	if q.peek() != 10 {
		t.Error()
	}
	if q.remove() != 10 {
		t.Error()
	}
	if q.remove() != 7 {
		t.Error()
	}
	if q.remove() != 6 {
		t.Error()
	}
	if q.remove() != 5 {
		t.Error()
	}
	if q.remove() != 2 {
		t.Error()
	}
	if !q.isEmpty() {
		t.Error()
	}
}
