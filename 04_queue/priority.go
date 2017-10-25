package queue

import (
	"errors"
)

type priorityQueue struct {
	maxSize int
	nItems  int
	data    []int64
}

func NewPriorityQueue(size int) priorityQueue {
	return priorityQueue{
		maxSize: size,
		nItems:  0,
		data:    make([]int64, size),
	}
}

func (q *priorityQueue) insert(elem int64) error {
	if q.isFull() {
		return errors.New("04_queue is full")
	}
	if q.nItems == 0 {
		q.data[q.nItems] = elem
		q.nItems++
	} else {
		var i int
		for i = q.nItems - 1; i >= 0; i-- {
			if q.data[i] < elem {
				q.data[i+1] = q.data[i]
			} else {
				break
			}
		}
		q.nItems++
		q.data[i+1] = elem
	}
	return nil
}

func (q *priorityQueue) remove() (int64, error) {
	if q.isEmpty() {
		return -1, errors.New("04_queue is empty")
	}
	el := q.data[q.nItems-1]
	q.nItems--
	return el, nil
}

func (q *priorityQueue) peek() (int64, error) {
	if q.isEmpty() {
		return -1, errors.New("04_queue is empty")
	}
	return q.data[q.nItems-1], nil
}

func (q *priorityQueue) isFull() bool {
	return q.size() == q.maxSize
}

func (q *priorityQueue) isEmpty() bool {
	return q.size() == 0
}

func (q *priorityQueue) size() int {
	return q.nItems
}
