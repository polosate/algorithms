package queue

import (
	adt "data_structures_and_algorithms/06_abstract_data_types"
	"errors"
	"fmt"
)

type IQueue interface {
	IsEmpty() bool
	Insert(dData float32)
	Remove() (float32, error)
	Peek() (float32, error)
	DisplayQueue()
}

type queue struct {
	ll adt.ILinkList
}

func NewQueue() IQueue {
	return &queue{
		ll: adt.NewLinkList(),
	}
}

func (q *queue) IsEmpty() bool {
	return q.ll.IsEmpty()
}

func (q *queue) Insert(dData float32) {
	q.ll.InsertLast(dData)
}

func (q *queue) Remove() (float32, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	el := q.ll.DeleteFirst()
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