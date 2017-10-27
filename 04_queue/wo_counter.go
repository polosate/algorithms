package queue

import (
	"errors"
	"fmt"
)

type IQueue interface {
	Insert(elem interface{}) error
	Remove() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
	IsFull() bool
	Size() int
	Display() string
	Display1() []interface{}
}

type abstractQueueWoCounter struct {
	maxSize int
	front   int
	rear    int
	data    []interface{}
}

func NewAbstractQueueWoCounter(size int) IQueue {
	return &abstractQueueWoCounter{
		maxSize: size + 1,
		front:   0,
		rear:    0,
		data:    make([]interface{}, size+1),
	}
}

func (q *abstractQueueWoCounter) Insert(elem interface{}) error {
	if q.IsFull() {
		return errors.New("04_queue is full")
	}
	if q.rear == q.maxSize {
		q.rear = 0
	}
	q.data[q.rear] = elem
	q.rear++
	return nil
}

func (q *abstractQueueWoCounter) Remove() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("04_queue is empty")
	}
	if q.front == q.maxSize {
		q.front = 0
	}
	el := q.data[q.front]
	q.front++
	return el, nil
}

func (q *abstractQueueWoCounter) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("04_queue is empty")
	}
	if q.front == q.maxSize {
		q.front = 0
	}
	return q.data[q.front], nil
}

func (q *abstractQueueWoCounter) IsFull() bool {
	return q.Size()+1 == q.maxSize
}

func (q *abstractQueueWoCounter) IsEmpty() bool {
	return q.Size() == 0
}

func (q *abstractQueueWoCounter) Size() int {
	if q.rear >= q.front {
		return q.rear - q.front
	} else {
		return q.maxSize - q.front + q.rear
	}
}

func (q *abstractQueueWoCounter) Display1() []interface{} {
	var res []interface{}
	if q.front <= q.rear {
		for i := q.front; i < q.rear; i++ {
			res = append(res, q.data[i])
		}
	} else {
		for i := q.front; i <= q.maxSize-1; i++ {
			res = append(res, q.data[i])
		}
		for i := 0; i < q.rear; i++ {
			res = append(res, q.data[i])
		}
	}
	return res
}

func (q *abstractQueueWoCounter) Display() string {
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
