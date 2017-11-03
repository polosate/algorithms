package queue

import (
	"errors"
	"fmt"

	. "algorithms/06_abstract_data_types/base"
)

type IQueue interface {
	IsEmpty() bool
	IsFull() bool
	Insert(dData float32)
	Remove() (float32, error)
	Peek() (float32, error)
	DisplayQueue()
}

type queue struct {
	ll      ILinkList
	size    int
	maxSize int
}

func NewQueue(maxSize int) IQueue {
	return &queue{
		ll:      NewLinkList(),
		size:    0,
		maxSize: maxSize,
	}
}

func (q *queue) IsEmpty() bool {
	return q.ll.IsEmpty()
}

func (q *queue) IsFull() bool {
	return q.size == q.maxSize
}

func (q *queue) Insert(dData float32) {
	if q.IsFull() {
		return
	} else {
		q.ll.InsertLast(dData)
		q.size++
	}
}

func (q *queue) Remove() (float32, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	el := q.ll.DeleteFirst()
	q.size--
	return el.GetData(), nil
}

func (q *queue) Peek() (float32, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	el := q.ll.PeekFirst()
	return el.GetData(), nil
}

func (q *queue) DisplayQueue() {
	fmt.Println("Queue (front --> rear)")
	q.ll.DisplayList()
}
