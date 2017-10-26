package personQueue

import (
	baseQ "data_structures_and_algorithms/04_queue"
)

type personQueue struct {
	baseQueue baseQ.IQueue
	queueNum  int
}

func NewQueue(size, num int) personQueue {
	return personQueue{
		baseQueue: baseQ.NewAbstractQueueWoCounter(size),
		queueNum:  num,
	}
}

func (pq *personQueue) AddPerson(person Person) error {
	return pq.baseQueue.Insert(person)
}

func (pq *personQueue) PeekPerson() (Person, error) {
	element, err := pq.baseQueue.Peek()
	if err != nil {
		return Person{}, err
	}
	person, _ := element.(Person)
	return person, nil
}

func (pq *personQueue) RemovePerson() (Person, error) {
	element, err := pq.baseQueue.Remove()
	if err != nil {
		return Person{}, err
	}
	person, _ := element.(Person)
	return person, nil
}

func (pq *personQueue) IsEmpty() bool {
	return pq.baseQueue.IsEmpty()
}

func (pq *personQueue) Size() int {
	return pq.baseQueue.Size()
}
