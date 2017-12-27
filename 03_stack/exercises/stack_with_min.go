package exercises

import "errors"

type iStack interface {
	push(int64) error
	pop() (int64, error)
	peek() (int64, error)
}

type stack struct {
	array []int64
	top   int
	size  int
}

func newStack(size int) iStack {
	return &stack{
		array: make([]int64, size),
		top:   -1,
		size:  size,
	}
}

func (s *stack) push(data int64) error {
	return nil
}

func (s *stack) pop() (int64, error) {
	return 0, nil
}

func (s *stack) peek() (int64, error) {
	return 0, nil
}

type stackWithMin struct {
	array    []int64
	top      int
	size     int
	minStack iStack
}

func newStackWithMin(size int) iStack {
	return &stackWithMin{
		array:    make([]int64, size),
		top:      -1,
		size:     size,
		minStack: newStack(size),
	}
}

func (s *stackWithMin) push(data int64) error {
	if s.top == s.size-1 {
		return errors.New("stack is full")
	}

	if s.top == -1 {
		s.array[s.top] = data
		s.minStack.push(data)
	} else {

	}

	return nil
}

func (s *stackWithMin) pop() (int64, error) {
	return 0, nil
}

func (s *stackWithMin) peek() (int64, error) {
	return 0, nil
}

func (s *stackWithMin) min() int64 {
	return 0
}
