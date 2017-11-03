package priority_queue

import (
	"errors"

	. "algorithms/06_abstract_data_types/base"
)

type IPriorityQueue interface {
	Insert(dData float32)
	Remove() (float32, error)
	Peek() (float32, error)
	IsEmpty() bool
	IsFull() bool
	Size() int

	DisplayQueue()
}

type priorityQueue struct {
	list    ISortedList
	maxSize int
	size    int
}

func NewPriorityQueue(maxSize int) IPriorityQueue {
	return &priorityQueue{
		list:    NewSortedList(),
		maxSize: maxSize,
		size:    0,
	}
}

func (pq *priorityQueue) Insert(dData float32) {
	if pq.size == pq.maxSize {
		return
	} else {
		pq.list.Insert(dData)
		pq.size++
	}
}

func (pq *priorityQueue) Remove() (float32, error) {
	item := pq.list.Remove()
	if item == nil {
		return 0, errors.New("queue is empty")
	}
	pq.size--
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
	return pq.size
}

func (pq *priorityQueue) IsEmpty() bool {
	return pq.list.IsEmpty()
}

func (pq *priorityQueue) IsFull() bool {
	return pq.size == pq.maxSize
}

func (pq *priorityQueue) DisplayQueue() {
	pq.list.DisplayList()
}
