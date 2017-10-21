package stack

import (
	"errors"
)

type abstractStack struct {
	size int
	top  int
	data []interface{}
}

func newStack(size int) *abstractStack {
	return &abstractStack{
		size: size,
		top:  -1,
		data: make([]interface{}, size),
	}
}

func (s *abstractStack) push(elem interface{}) error {
	if s.isFull() {
		return errors.New("abstractStack is full")
	} else {
		s.top++
		s.data[s.top] = elem
		return nil
	}
}

func (s *abstractStack) pop() (interface{}, error) {
	if s.isEmpty() {
		return -1, errors.New("abstractStack is empty")
	}
	el := s.data[s.top]
	s.top--
	return el, nil
}

func (s *abstractStack) peek() (interface{}, error) {
	if s.isEmpty() {
		return -1, errors.New("abstractStack is empty")
	}
	return s.data[s.top], nil
}

func (this *abstractStack) isEmpty() bool {
	return this.top == -1
}

func (this *abstractStack) isFull() bool {
	return this.top == this.size-1
}
