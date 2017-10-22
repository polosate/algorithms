package queue

import (
	"reflect"
	"testing"
)

func TestAbstractWoCounter(t *testing.T) {
	q := newAbstractQueueWoCounter(5)
	if !q.isEmpty() {
		t.Error("Expected isEmpty", true, "actual", false)
	}

	q.insert(1)
	q.insert(2)
	q.insert(3)
	q.insert(4)

	s := q.size()
	if s != 4 {
		t.Error("Expected size", 4, "actual", s)
	}

	el, _ := q.remove()
	s = q.size()
	if s != 3 {
		t.Error("Expected size", 3, "actual", s)
	}
	if !reflect.DeepEqual(el, 1) {
		t.Error("Expected el", 1, "actual", el)
	}

	q.insert(5)
	s = q.size()
	if s != 4 {
		t.Error("Expected size", 4, "actual", s)
	}
	q.insert(6)
	s = q.size()
	if s != 5 {
		t.Error("Expected size", 5, "actual", s)
	}

	err := q.insert(10)
	if err.Error() != "queue is full" {
		t.Error("Expected err", "queue is full", "actual", err.Error())
	}

	q.remove()
	s = q.size()
	if s != 4 {
		t.Error("Expected size", 4, "actual", s)
	}

	q.insert(8)
	s = q.size()
	if s != 5 {
		t.Error("Expected size", 5, "actual", s)
	}

	q.remove()
	s = q.size()
	if s != 4 {
		t.Error("Expected size", 4, "actual", s)
	}
	q.remove()
	s = q.size()
	if s != 3 {
		t.Error("Expected size", 3, "actual", s)
	}
	q.remove()
	s = q.size()
	if s != 2 {
		t.Error("Expected size", 2, "actual", s)
	}
	q.remove()
	s = q.size()
	if s != 1 {
		t.Error("Expected size", 1, "actual", s)
	}

	q.remove()
	s = q.size()
	if s != 0 {
		t.Error("Expected size", 0, "actual", s)
	}

	el, err = q.remove()
	if err.Error() != "queue is empty" {
		t.Error("Expected err", "queue is empty", "actual", err.Error())
	}
}
