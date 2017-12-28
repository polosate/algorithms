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

type iStackWithMin interface {
	iStack
	min() (int64, error)
}

type stackWithMin struct {
	array    []int64
	top      int
	size     int
	minStack iStack
}

func newStackWithMin(size int) iStackWithMin {
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
	min, err := s.min()
	if err != nil || data <= min {
		s.minStack.push(data)
	}
	s.top++
	s.array[s.top] = data
	return nil
}

func (s *stackWithMin) pop() (int64, error) {
	if s.top == -1 {
		return 0, errors.New("stack is empty")
	}
	data := s.array[s.top]
	s.top--
	min, _ := s.min()
	if data == min {
		s.minStack.pop()
	}
	return data, nil
}

func (s *stackWithMin) peek() (int64, error) {
	if s.top == -1 {
		return 0, errors.New("stack is empty")
	}
	return s.array[s.top], nil
}

func (s *stackWithMin) min() (int64, error) {
	return s.minStack.peek()

}
