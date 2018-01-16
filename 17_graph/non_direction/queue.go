package non_direction

import (
	abstractQ "algorithms/04_queue"
)

type queue struct {
	abstractQ.IQueue
}

func newQueue(size int) queue {
	return queue{abstractQ.NewAbstractQueueWoCounter(size)}
}

func (q *queue) isEmpty() bool {
	return q.IsEmpty()
}

func (q *queue) add(v int) {
	q.Insert(v)
}

func (q *queue) remove() int {
	i, _ := q.Remove()
	element, _ := i.(int)
	return element
}

func (q *queue) peek() int {
	i, _ := q.Peek()
	element, _ := i.(int)
	return element
}
