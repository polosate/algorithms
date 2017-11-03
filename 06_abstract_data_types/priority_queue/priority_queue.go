package priority_queue

import (
	. "data_structures_and_algorithms/06_abstract_data_types"
	"errors"
)

type IPriorityQueue interface {
	Insert(dData float32)
	Remove() (float32, error)
	Peek() (float32, error)
	IsEmpty() bool
	Size() int

	DisplayQueue()
}

type priorityQueue struct {
	list ISortedList
}

func NewPriorityQueue() IPriorityQueue {
	return &priorityQueue{
		list: NewSortedList(),
	}
}

func (pq *priorityQueue) Insert(dData float32) {
	pq.list.Insert(dData)
}

func (pq *priorityQueue) Remove() (float32, error) {
	item := pq.list.Remove()
	if item == nil {
		return 0, errors.New("queue is empty")
	}
	return item.GetData(), nil
}

func (pq *priorityQueue) Peek() (float32, error) {
	item := pq.list.Peek()
	if item == nil {
		return 0, errors.New("queue is empty")
	}
	return item.GetData(), nil
}

func (pq *priorityQueue) Size() int {
	return pq.list.Size()
}

func (pq *priorityQueue) IsEmpty() bool {
	return pq.list.IsEmpty()
}

func (pq *priorityQueue) DisplayQueue() {
	pq.list.DisplayList()
}
