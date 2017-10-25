package queue

import (
	"errors"
	"fmt"
)

type abstractQueueWoCounter struct {
	maxSize int
	front   int
	rear    int
	data    []interface{}
}

func newAbstractQueueWoCounter(size int) abstractQueueWoCounter {
	return abstractQueueWoCounter{
		maxSize: size + 1,
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
	return q.data[q.front], nil
}

func (q *abstractQueueWoCounter) isFull() bool {
	return q.size()+1 == q.maxSize
}

func (q *abstractQueueWoCounter) isEmpty() bool {
	return q.size() == 0
}

func (q *abstractQueueWoCounter) size() int {
	if q.rear >= q.front {
		return q.rear - q.front
	} else {
		return q.maxSize - q.front + q.rear
	}
}

func (q *abstractQueueWoCounter) display() string {
	var res string
	if q.front <= q.rear {
		for i := q.front; i < q.rear; i++ {
			res += fmt.Sprint(q.data[i])
		}
	} else {
		for i := q.front; i <= q.maxSize-1; i++ {
			res += fmt.Sprint(q.data[i])
		}
		for i := 0; i < q.rear; i++ {
			res += fmt.Sprint(q.data[i])
		}
	}
	return res
}
