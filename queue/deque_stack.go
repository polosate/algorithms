package queue

import "errors"

type IStack interface {
	Push(elem interface{}) error
	Pop() (interface{}, error)
	//Peek() (interface{}, error)
	IsEmpty() bool
	IsFull() bool
}

type stack struct {
	deque deque
}

func NewStack(size int) IStack {
	return &stack{
		deque: newDeque(size),
	}
}

func (s *stack) Push(elem interface{}) error {
	err := s.deque.insertRight(elem)
	if err != nil {
		return errors.New("stack is full")
	}
	return nil
}

func (s *stack) Pop() (interface{}, error) {
	elem, err := s.deque.removeRight()
	if err != nil {
		return nil, errors.New("stack is empty")
	}
	return elem, nil
}

//func (s stack) Peek() (interface{}, error) {
//
//}

func (s *stack) IsEmpty() bool {
	return s.deque.isEmpty()
}

func (s *stack) IsFull() bool {
	return s.deque.isFull()
}
