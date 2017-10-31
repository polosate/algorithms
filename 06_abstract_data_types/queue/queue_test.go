package queue

import (
	"testing"
)

func TestLinkLIstQueue(t *testing.T) {
	q := NewQueue()
	if !q.IsEmpty() {
		t.Error("Expected IsEmpty", true, "actual", false)
	}

	q.Insert(1.1)
	q.Insert(2.2)
	q.Insert(3.3)
	q.Insert(4.4)
	q.DisplayQueue()

	el, err := q.Remove()

	if err != nil {
		t.Error("Expected err", nil, "actual", err.Error())
	}
	if el != 1.1 {
		t.Error("Expected el", 1.1, "actual", el)
	}

	q.DisplayQueue()

}
