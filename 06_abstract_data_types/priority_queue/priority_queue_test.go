package priority_queue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue(5)
	if !q.IsEmpty() {
		t.Error("Expected IsEmpty", true, "actual", false)
	}
	q.Insert(5)
	q.Insert(2)
	q.Insert(1)
	q.Insert(4)
	q.Insert(7)

	el, err := q.Remove()
	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}
	if el != 1 {
		t.Error("Expected el", 1, "actual", el)
	}

	el, err = q.Remove()
	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}
	if el != 2 {
		t.Error("Expected el", 2, "actual", el)
	}

	el, err = q.Remove()
	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}
	if el != 4 {
		t.Error("Expected el", 4, "actual", el)
	}

	el, err = q.Remove()
	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}
	if el != 5 {
		t.Error("Expected el", 5, "actual", el)
	}
	if q.Size() != 1 {
		t.Error("Expected size", 1, "actual", q.Size())
	}

	el, err = q.Peek()
	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}
	if el != 7 {
		t.Error("Expected el", 7, "actual", el)
	}
	if q.Size() != 1 {
		t.Error("Expected size", 1, "actual", q.Size())
	}

	el, err = q.Remove()
	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}
	if el != 7 {
		t.Error("Expected el", 7, "actual", el)
	}

	_, err = q.Remove()
	if err.Error() != "queue is empty" {
		t.Error("Expected err", "queue is empty", "actual", err.Error())
	}
}
