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

func TestDisplay(t *testing.T) {
	q := newAbstractQueueWoCounter(5)
	res := q.display()
	if res != "" {
		t.Error("Expected str", "", "actual", res)
	}

	q.insert(1)
	res = q.display()
	if res != "1" {
		t.Error("Expected str", "1", "actual", res)
	}

	q.insert(2)
	q.insert(3)
	q.insert(4)
	q.insert(5)
	res = q.display()
	if res != "12345" {
		t.Error("Expected str", "12345", "actual", res)
	}

	q.remove()
	res = q.display()
	if res != "2345" {
		t.Error("Expected str", "2345", "actual", res)
	}

	q.insert(6)
	res = q.display()
	if res != "23456" {
		t.Error("Expected str", "23456", "actual", res)
	}

	q.remove()
	q.remove()
	q.remove()
	q.remove()
	res = q.display()
	if res != "6" {
		t.Error("Expected str", "6", "actual", res)
	}

	q.remove()
	res = q.display()
	if res != "" {
		t.Error("Expected str", "", "actual", res)
	}

}
