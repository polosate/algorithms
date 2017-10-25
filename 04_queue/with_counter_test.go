package queue

import (
	"reflect"
	"testing"
)

func TestAbstractWithCounter(t *testing.T) {
	q := NewAbstractQueue(5)

	if !q.IsEmpty() {
		t.Error("Expected IsEmpty", true, "actual", false)
	}
	q.Insert(1)
	q.Insert(2)
	q.Insert(3)
	q.Insert(4)

	s := q.Size()
	if s != 4 {
		t.Error("Expected Size", 4, "actual", s)
	}

	el, err := q.Remove()

	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}

	if !reflect.DeepEqual(el, 1) {
		t.Error("Expected el", 1, "actual", el)
	}

	q.Insert(5)
	q.Insert(6)
	err = q.Insert(7)
	if err.Error() != "04_queue is full" {
		t.Error("Expected err", "04_queue is full", "actual", err.Error())
	}

	el, err = q.Remove()
	if !reflect.DeepEqual(el, 2) {
		t.Error("Expected el", 2, "actual", el)
	}
	el, err = q.Peek()
	if !reflect.DeepEqual(el, 3) {
		t.Error("Expected el", 3, "actual", el)
	}
	el, err = q.Remove()
	if !reflect.DeepEqual(el, 3) {
		t.Error("Expected el", 3, "actual", el)
	}
	el, err = q.Remove()
	if !reflect.DeepEqual(el, 4) {
		t.Error("Expected el", 4, "actual", el)
	}
	el, err = q.Remove()
	if !reflect.DeepEqual(el, 5) {
		t.Error("Expected el", 5, "actual", el)
	}
	el, err = q.Remove()
	if !reflect.DeepEqual(el, 6) {
		t.Error("Expected el", 6, "actual", el)
	}

	_, err = q.Remove()
	if err.Error() != "04_queue is empty" {
		t.Error("Expected err", "04_queue is empty", "actual", err.Error())
	}
}
