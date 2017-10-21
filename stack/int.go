package stack

import (
	"errors"
	"fmt"
)

type intStack struct {
	*abstractStack
}

func NewIntStack(size int) *intStack {
	return &intStack{
		newAbstractStack(size),
	}
}

func (s *intStack) Push(el int64) error {
	return s.abstractStack.push(el)
}

func (s *intStack) Pop() (int64, error) {
	el, err := s.Peek()
	if err != nil {
		return -1, err
	} else {
		s.top--
		return int64(el), nil
	}
}

func (s *intStack) Peek() (int64, error) {
	el, err := s.abstractStack.peek()
	if err != nil {
		return -1, err
	}
	if intEl, ok := el.(int64); !ok {
		return -1, errors.New(fmt.Sprintf("received incompatibles type: value %v, type^ %T", intEl, intEl))
	} else {
		return int64(intEl), nil
	}
}

func (s *intStack) IsEmpty() bool {
	return s.abstractStack.isEmpty()
}

func (s *intStack) IsFull() bool {
	return s.abstractStack.isFull()
}
