package stack

import "errors"

type stack struct {
	maxSize int
	data    []int64
	top     int
}

func newStack(size int) stack {
	return stack{
		maxSize: size,
		data:    make([]int64, size),
		top:     -1,
	}
}

func (this *stack) Push(elem int64) error {
	if this.IsFull() {
		return errors.New("stack is full")
	} else {
		this.top++
		this.data[this.top] = elem
		return nil
	}
}

func (this *stack) Pop() (int64, error) {
	if this.IsEmpty() {
		return -1, errors.New("stack is empty")
	} else {
		el := this.data[this.top]
		this.top--
		return el, nil
	}
}

func (this *stack) Peek() (int64, error) {
	if this.IsEmpty() {
		return -1, errors.New("stack is empty")
	} else {
		return this.data[this.top], nil
	}
}

func (this *stack) IsEmpty() bool {
	return this.top == -1
}

func (this *stack) IsFull() bool {
	return this.top == this.maxSize-1
}
