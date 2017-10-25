package queue

import (
	"errors"
)

type deque struct {
	maxSize int
	front   int
	rear    int
	data    []interface{}
}

func newDeque(size int) deque {
	return deque{
		maxSize: size + 1,
		front:   0,
		rear:    0,
		data:    make([]interface{}, size+1),
	}
}

func (d *deque) insertLeft(el interface{}) error {
	if d.isFull() {
		return errors.New("deque is full")
	}
	if d.front == 0 {
		d.front = d.maxSize
	}
	d.front--
	d.data[d.front] = el
	return nil
}

func (d *deque) insertRight(el interface{}) error {
	if d.isFull() {
		return errors.New("deque is full")
	}
	if d.rear == d.maxSize {
		d.rear = 0
	}
	d.data[d.rear] = el
	d.rear++
	return nil
}

func (d *deque) removeLeft() (interface{}, error) {
	if d.isEmpty() {
		return nil, errors.New("deque is empty")
	}
	if d.front == d.maxSize {
		d.front = 0
	}
	el := d.data[d.front]
	d.data[d.front] = nil
	d.front++
	return el, nil
}

func (d *deque) removeRight() (interface{}, error) {
	if d.isEmpty() {
		return nil, errors.New("deque is empty")
	}

	if d.rear == 0 {
		d.rear = d.maxSize
	}
	d.rear--
	el := d.data[d.rear]
	d.data[d.rear] = nil
	return el, nil
}

func (d *deque) isFull() bool {
	return d.size() == d.maxSize-1

}

func (d *deque) isEmpty() bool {
	return d.size() == 0
}

func (d *deque) size() int {
	if d.front <= d.rear {
		return d.rear - d.front
	} else {
		return d.maxSize - d.front + d.rear
	}
}
