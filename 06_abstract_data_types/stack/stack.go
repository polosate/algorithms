package stack

import (
	"errors"
	"fmt"

	. "algorithms/06_abstract_data_types/base"
)

type IStack interface {
	Push(dData float32)
	Pop() (float32, error)
	Peek() (float32, error)
	IsEmpty() bool
	IsFull() bool
	DisplayStack()
}

type stack struct {
	ll      ILinkList
	size    int
	maxSize int
}

func NewStack(maxSize int) IStack {
	return &stack{
		ll:      NewLinkList(),
		size:    0,
		maxSize: maxSize,
	}
}

func (s *stack) IsEmpty() bool {
	return s.ll.IsEmpty()
}

func (s *stack) IsFull() bool {
	return s.size == s.maxSize
}

func (s *stack) Push(dData float32) {
	if s.IsFull() {
		return
	} else {
		s.ll.InsertFirst(dData)
		s.size++
	}

}

func (s *stack) Pop() (float32, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	el := s.ll.DeleteFirst()
	s.size--
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
