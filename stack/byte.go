package stack

import (
	"errors"
	"fmt"
)

type byteStack struct {
	abstractStack
}

func NewByteStack(size int) byteStack {
	return byteStack{
		newAbstractStack(size),
	}
}

func (s *byteStack) Push(el byte) error {
	return s.abstractStack.push(el)
}

func (s *byteStack) Pop() (byte, error) {
	el, err := s.Peek()
	if err != nil {
		return 0, err
	} else {
		s.top--
		return el, nil
	}
}

func (s *byteStack) Peek() (byte, error) {
	el, err := s.abstractStack.peek()
	if err != nil {
		return 0, err
	}
	if byteEl, ok := el.(byte); !ok {
		return 0, errors.New(fmt.Sprintf("received incompatibles type: value %v, type %T", el, el))
	} else {
		return byteEl, nil
	}
}

func (s *byteStack) IsEmpty() bool {
	return s.abstractStack.isEmpty()
}

func (s *byteStack) IsFull() bool {
	return s.abstractStack.isFull()
}
