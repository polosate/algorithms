package exercises

import "errors"

type iStack interface {
	push(int64) error
	pop() (int64, error)
	peek() (int64, error)
	isFull() bool
	isEmpty() bool
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

func (s *stack) isFull() bool {
	return s.top == s.size-1
}

func (s *stack) isEmpty() bool {
	return s.top == -1
}

func (s *stack) push(data int64) error {
	if s.top == s.size-1 {
		return errors.New("stack is full")
	}
	s.top++
	s.array[s.top] = data
	return nil
}

func (s *stack) pop() (int64, error) {
	if s.top == -1 {
		return 0, errors.New("stack is empty")
	}
	data := s.array[s.top]
	s.top--
	return data, nil
}

func (s *stack) peek() (int64, error) {
	if s.top == -1 {
		return 0, errors.New("stack is empty")
	}
	return s.array[s.top], nil
}
