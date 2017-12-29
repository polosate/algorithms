package exercises

import "github.com/pkg/errors"

type queue struct {
	size     int
	curSize  int
	newStack iStack
	oldStack iStack
}

func newQueue(size int) queue {
	return queue{
		size:     size,
		curSize:  0,
		newStack: newStack(size),
		oldStack: newStack(size),
	}
}

func (q *queue) isFull() bool {
	return q.curSize == q.size
}

func (q *queue) isEmpty() bool {
	return q.curSize == 0
}

func (q *queue) add(data int64) error {
	if err := q.newStack.push(data); err != nil {
		return err
	}
	q.curSize++
	return nil
}

func (q *queue) remove() (int64, error) {
	if q.isEmpty() {
		return 0, errors.New("queue is empty")
	}
	q.shift()
	data, _ := q.oldStack.pop()
	q.curSize--
	return data, nil
}

func (q *queue) shift() {
	if q.oldStack.isEmpty() {
		for !q.newStack.isEmpty() {
			data, _ := q.newStack.pop()
			q.oldStack.push(data)
		}
	}
}
