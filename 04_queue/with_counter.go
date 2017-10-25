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

func NewAbstractQueue(size int) abstractQueue {
	return abstractQueue{
		maxSize: size,
		front:   0,
		nItems:  0,
		rear:    0,
		data:    make([]interface{}, size),
	}
}

func (q *abstractQueue) Insert(elem interface{}) error {
	if q.IsFull() {
		return errors.New("04_queue is full")
	}
	if q.rear == q.maxSize {
		q.rear = 0
	}
	q.data[q.rear] = elem
	q.rear++
	q.nItems++
	return nil
}

func (q *abstractQueue) Remove() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("04_queue is empty")
	}
	if q.front == q.maxSize {
		q.front = 0
	}
	el := q.data[q.front]
	q.front++
	q.nItems--
	return el, nil
}

func (q *abstractQueue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("04_queue is empty")
	}
	return q.data[q.front], nil
}

func (q *abstractQueue) IsFull() bool {
	return q.nItems == q.maxSize
}

func (q *abstractQueue) IsEmpty() bool {
	return q.nItems == 0
}

func (q *abstractQueue) Size() int {
	return q.nItems
}
