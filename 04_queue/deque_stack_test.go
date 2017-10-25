package queue

import (
	"reflect"
	"testing"
)

func TestDequeStack(t *testing.T) {
	stack := NewStack(5)

	if !stack.IsEmpty() {
		t.Error("expected IsEmpty", true, "have", false)
	}

	el, err := stack.Pop()
	if err.Error() != "03_stack is empty" {
		t.Error("expected err", "03_stack is empty", "have", err.Error())
	}

	stack.Push(22)
	stack.Push(77)

	if stack.IsEmpty() {
		t.Error("expected IsEmpty", false, "have", true)
	}

	if stack.IsFull() {
		t.Error("expected IsFull", false, "have", true)
	}

	stack.Push(33)
	stack.Push(44)
	stack.Push(55)

	err = stack.Push(66)
	if err.Error() != "03_stack is full" {
		t.Error("expected err", "03_stack is full", "have", err.Error())
	}

	if !stack.IsFull() {
		t.Error("expected IsFull", true, "have", false)
	}

	el, err = stack.Pop()
	if err != nil {
		t.Error("expected error", nil, "have", err.Error())
	}
	if !reflect.DeepEqual(el, 55) {
		t.Error("expected element", 55, "have", el)
	}
}
