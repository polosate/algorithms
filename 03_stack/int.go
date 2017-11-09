package stack

import (
	"errors"
	"fmt"
)

type IntStack struct {
	abstractStack
}

func NewIntStack(size int) IntStack {
	return IntStack{
		newAbstractStack(size),
	}
}

func (s *IntStack) Push(el int64) error {
	return s.abstractStack.push(el)
}

func (s *IntStack) Pop() (int64, error) {
	el, err := s.Peek()
	if err != nil {
		return -1, err
	} else {
		s.top--
		return int64(el), nil
	}
}

func (s *IntStack) Peek() (int64, error) {
	el, err := s.abstractStack.peek()
	if err != nil {
		return -1, err
	}
	if intEl, ok := el.(int64); !ok {
		return -1, errors.New(fmt.Sprintf("received incompatibles type: value %v, type %T", el, el))
	} else {
		return int64(intEl), nil
	}
}

func (s *IntStack) IsEmpty() bool {
	return s.abstractStack.isEmpty()
}

func (s *IntStack) IsFull() bool {
	return s.abstractStack.isFull()
}

func (s *IntStack) DisplayStack() {

	for i := 0; i <= s.top; i++ {
		fmt.Print(s.abstractStack.data[i], "; ")
	}
	fmt.Println()
}
