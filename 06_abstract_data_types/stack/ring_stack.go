package stack

import (
	"algorithms/05_linked_list/ring"
	"errors"
	"fmt"
)

type ringStack struct {
	ring    ring.IRing
	size    int
	maxSize int
}

func NewRingStack(maxSize int) IStack {
	return &ringStack{
		ring:    ring.NewRing(),
		size:    0,
		maxSize: maxSize,
	}
}

func (s *ringStack) IsEmpty() bool {
	return s.ring.IsEmpty()
}

func (s *ringStack) IsFull() bool {
	return s.size == s.maxSize
}

func (s *ringStack) Push(dData float32) {
	if s.IsFull() {
		return
	} else {
		s.ring.Insert(dData)
		s.size++
	}
}

func (s *ringStack) Pop() (float32, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	el := s.ring.Remove()
	s.size--
	return el.GetValue(), nil
}

func (s *ringStack) Peek() (float32, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	el := s.ring.Peek()
	return el.GetValue(), nil
}

func (s *ringStack) DisplayStack() {
	fmt.Println("Stack (top --> bottom)")
	s.ring.DisplayRing()
}
