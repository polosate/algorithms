package stack

import (
	"errors"
	"fmt"
)

type DoubleStack struct {
	abstractStack
}

func NewDoubleStack(size int) DoubleStack {
	return DoubleStack{
		newAbstractStack(size),
	}
}

func (s *DoubleStack) Push(el float64) error {
	return s.abstractStack.push(el)
}

func (s *DoubleStack) Pop() (float64, error) {
	el, err := s.Peek()
	if err != nil {
		return -1, err
	} else {
		s.top--
		return float64(el), nil
	}
}

func (s *DoubleStack) Peek() (float64, error) {
	el, err := s.abstractStack.peek()
	if err != nil {
		return -1, err
	}
	if intEl, ok := el.(float64); !ok {
		return -1, errors.New(fmt.Sprintf("received incompatibles type: value %v, type %T", el, el))
	} else {
		return float64(intEl), nil
	}
}

func (s *DoubleStack) IsEmpty() bool {
	return s.abstractStack.isEmpty()
}

func (s *DoubleStack) IsFull() bool {
	return s.abstractStack.isFull()
}
