package stack

import "testing"

func TestIntStack(t *testing.T) {
	stack := NewStack()
	if !stack.IsEmpty() {
		t.Error("expected IsEmpty", true, "have", false)
	}
	el, err := stack.Pop()
	if err.Error() != "stack is empty" {
		t.Error("expected err", "stack is empty", "have", err.Error())
	}

	stack.Push(22)
	stack.Push(77)
	stack.Push(33)
	stack.Push(44)
	stack.Push(55)
	stack.DisplayStack()

	if stack.IsEmpty() {
		t.Error("expected IsEmpty", false, "have", true)
	}

	el, err = stack.Pop()
	if err != nil {
		t.Error("expected error", nil, "have", err.Error())
	}
	if el != float32(55) {
		t.Error("expected element", 55, "have", el)
	}
	stack.DisplayStack()

	el, _ = stack.Peek()
	if el != float32(44) {
		t.Error("expected element", 44, "have", el)
	}

	el, _ = stack.Peek()
	if el != float32(44) {
		t.Error("expected element", 44, "have", el)
	}
	stack.DisplayStack()

	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.DisplayStack()
	el, _ = stack.Pop()
	if el != float32(22) {
		t.Error("expected element", 22, "have", el)
	}
	stack.DisplayStack()
}
