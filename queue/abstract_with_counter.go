package queue

import (
	"errors"
)

type abstractQueue struct {
	maxSize int
	front   int
	rear    int
	nItems  int
	data    []interface{}
}

func newAbstractQueue(size int) *abstractQueue {
	return &abstractQueue{
		maxSize: size,
		front:   0,
		nItems:  0,
		rear:    0,
		data:    make([]interface{}, size),
	}
}

func (q *abstractQueue) insert(elem interface{}) error {
	if q.isFull() {
		return errors.New("queue is full")
	}
	if q.rear == q.maxSize {
		q.rear = 0
	}
	q.data[q.rear] = elem
	q.rear++
	q.nItems++
	return nil
}

func (q *abstractQueue) remove() (interface{}, error) {
	if q.isEmpty() {
		return nil, errors.New("queue is empty")
	}
	if q.front == q.maxSize {
		q.front = 0
	}
	el := q.data[q.front]
	q.front++
	q.nItems--
	return el, nil
}

func (q *abstractQueue) peek() (interface{}, error) {
	if q.isEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.data[q.front], nil
}

func (q *abstractQueue) isFull() bool {
	return q.nItems == q.maxSize
}

func (q *abstractQueue) isEmpty() bool {
	return q.nItems == 0
}

func (q *abstractQueue) size() int {
	return q.nItems
}
