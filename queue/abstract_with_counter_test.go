package queue

import (
	"reflect"
	"testing"
)

func TestAbstractWithCounter(t *testing.T) {
	q := newAbstractQueue(5)

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

	el, err := q.remove()

	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}

	if !reflect.DeepEqual(el, 1) {
		t.Error("Expected el", 1, "actual", el)
	}

	q.insert(5)
	q.insert(6)
	err = q.insert(7)
	if err.Error() != "queue is full" {
		t.Error("Expected err", "queue is full", "actual", err.Error())
	}

	el, err = q.remove()
	if !reflect.DeepEqual(el, 2) {
		t.Error("Expected el", 2, "actual", el)
	}
	el, err = q.peek()
	if !reflect.DeepEqual(el, 3) {
		t.Error("Expected el", 3, "actual", el)
	}
	el, err = q.remove()
	if !reflect.DeepEqual(el, 3) {
		t.Error("Expected el", 3, "actual", el)
	}
	el, err = q.remove()
	if !reflect.DeepEqual(el, 4) {
		t.Error("Expected el", 4, "actual", el)
	}
	el, err = q.remove()
	if !reflect.DeepEqual(el, 5) {
		t.Error("Expected el", 5, "actual", el)
	}
	el, err = q.remove()
	if !reflect.DeepEqual(el, 6) {
		t.Error("Expected el", 6, "actual", el)
	}

	_, err = q.remove()
	if err.Error() != "queue is empty" {
		t.Error("Expected err", "queue is empty", "actual", err.Error())
	}
}
