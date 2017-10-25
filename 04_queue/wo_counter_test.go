package queue

import (
	"reflect"
	"testing"
)

func TestAbstractWoCounter(t *testing.T) {
	q := NewAbstractQueueWoCounter(5)
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

	el, _ := q.Remove()
	s = q.Size()
	if s != 3 {
		t.Error("Expected Size", 3, "actual", s)
	}
	if !reflect.DeepEqual(el, 1) {
		t.Error("Expected el", 1, "actual", el)
	}

	q.Insert(5)
	s = q.Size()
	if s != 4 {
		t.Error("Expected Size", 4, "actual", s)
	}
	q.Insert(6)
	s = q.Size()
	if s != 5 {
		t.Error("Expected Size", 5, "actual", s)
	}

	err := q.Insert(10)
	if err.Error() != "04_queue is full" {
		t.Error("Expected err", "04_queue is full", "actual", err.Error())
	}

	q.Remove()
	s = q.Size()
	if s != 4 {
		t.Error("Expected Size", 4, "actual", s)
	}

	q.Insert(8)
	s = q.Size()
	if s != 5 {
		t.Error("Expected Size", 5, "actual", s)
	}

	q.Remove()
	s = q.Size()
	if s != 4 {
		t.Error("Expected Size", 4, "actual", s)
	}
	q.Remove()
	s = q.Size()
	if s != 3 {
		t.Error("Expected Size", 3, "actual", s)
	}
	q.Remove()
	s = q.Size()
	if s != 2 {
		t.Error("Expected Size", 2, "actual", s)
	}
	q.Remove()
	s = q.Size()
	if s != 1 {
		t.Error("Expected Size", 1, "actual", s)
	}

	q.Remove()
	s = q.Size()
	if s != 0 {
		t.Error("Expected Size", 0, "actual", s)
	}

	el, err = q.Remove()
	if err.Error() != "04_queue is empty" {
		t.Error("Expected err", "04_queue is empty", "actual", err.Error())
	}
}

func TestDisplay(t *testing.T) {
	q := NewAbstractQueueWoCounter(5)
	res := q.Display()
	if res != "" {
		t.Error("Expected str", "", "actual", res)
	}

	q.Insert(1)
	res = q.Display()
	if res != "1" {
		t.Error("Expected str", "1", "actual", res)
	}

	q.Insert(2)
	q.Insert(3)
	q.Insert(4)
	q.Insert(5)
	res = q.Display()
	if res != "12345" {
		t.Error("Expected str", "12345", "actual", res)
	}

	q.Remove()
	res = q.Display()
	if res != "2345" {
		t.Error("Expected str", "2345", "actual", res)
	}

	q.Insert(6)
	res = q.Display()
	if res != "23456" {
		t.Error("Expected str", "23456", "actual", res)
	}

	q.Remove()
	q.Remove()
	q.Remove()
	q.Remove()
	res = q.Display()
	if res != "6" {
		t.Error("Expected str", "6", "actual", res)
	}

	q.Remove()
	res = q.Display()
	if res != "" {
		t.Error("Expected str", "", "actual", res)
	}
}
