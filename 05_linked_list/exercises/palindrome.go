package exercises

import "github.com/pkg/errors"

type stack struct {
	array []int64
	top   int
}

func newStack() *stack {
	return &stack{
		array: make([]int64, 256),
		top:   -1,
	}
}

func (s *stack) push(val int64) {
	s.top++
	s.array[s.top] = val
}

func (s *stack) pop() (int64, error) {
	if s.top == -1 {
		return -1, errors.New("stack is empty")
	}
	val := s.array[s.top]
	s.top--
	return val, nil
}

func isPalindrome(l *singleList) bool {
	temp := newStack()
	slowIter := l.GetIterator()
	fastIter := l.GetIterator()

	for fastIter.GetCurrent() != nil && fastIter.GetCurrent().next != nil {
		temp.push(slowIter.GetCurrent().value)
		slowIter.NextLink()
		fastIter.NextLink()
		fastIter.NextLink()
	}

	if fastIter.GetCurrent() != nil {
		slowIter.NextLink()
	}

	for slowIter.GetCurrent() != nil {
		expected, err := temp.pop()
		if err != nil {
			return false
		}
		if slowIter.GetCurrent().value != expected {
			return false
		}
		slowIter.NextLink()
	}
	return true
}
