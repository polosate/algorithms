package queue

import (
	"errors"
)

type abstractQueueWoCounter struct {
	maxSize int
	front   int
	rear    int
	data    []interface{}
}

func newAbstractQueueWoCounter(size int) *abstractQueueWoCounter {
	return &abstractQueueWoCounter{
		maxSize: size,
		front:   0,
		rear:    0,
		data:    make([]interface{}, size+1),
	}
}

func (q *abstractQueueWoCounter) insert(elem interface{}) error {
	if q.isFull() {
		return errors.New("queue is full")
	}
	if q.rear == q.maxSize {
		q.rear = 0
	}
	q.data[q.rear] = elem
	q.rear++
	return nil
}

// 1 2 3 4 0 0
func (q *abstractQueueWoCounter) remove() (interface{}, error) {
	if q.isEmpty() {
		return nil, errors.New("queue is empty")
	}
	if q.front == q.maxSize {
		q.front = 0
	}
	el := q.data[q.front]
	q.front++
	return el, nil
}

func (q *abstractQueueWoCounter) peek() (interface{}, error) {
	if q.isEmpty() {
		return nil, errors.New("queue is empty")
	}
	if q.front == q.maxSize {
		q.front = 0
	}
	return q.data[q.front], nil
}

func (q *abstractQueueWoCounter) isFull() bool {
	return false
}

func (q *abstractQueueWoCounter) isEmpty() bool {
	return false
}

func (q *abstractQueueWoCounter) size() int {
	return -1
}
