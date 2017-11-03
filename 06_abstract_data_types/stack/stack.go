package stack

import (
	"errors"
	"fmt"

	. "data_structures_and_algorithms/06_abstract_data_types/list"
)

type IStack interface {
	Push(dData float32)
	Pop() (float32, error)
	Peek() (float32, error)
	IsEmpty() bool
	DisplayStack()
}

type stack struct {
	ll ILinkList
}

func NewStack() IStack {
	return &stack{
		ll: NewLinkList(),
	}
}

func (s *stack) IsEmpty() bool {
	return s.ll.IsEmpty()

}

func (s *stack) Push(dData float32) {
	s.ll.InsertFirst(dData)
}

func (s *stack) Pop() (float32, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	el := s.ll.DeleteFirst()
	return el.GetData(), nil
}

func (s *stack) Peek() (float32, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	el := s.ll.PeekFirst()
	return el.GetData(), nil
}

func (s *stack) DisplayStack() {
	fmt.Println("Stack (top --> bottom)")
	s.ll.DisplayList()
}
