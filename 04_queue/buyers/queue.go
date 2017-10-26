package buyers

import (
	"fmt"
	"time"

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

func (pq *personQueue) RemovePerson() {
	element, err := pq.baseQueue.Peek()
	if err != nil {
		fmt.Printf("Queue %d is empty!\n", pq.queueNum)
	}
	person, _ := element.(Person)
	time.Sleep(time.Second * person.GetDelay())
	pq.baseQueue.Remove()
}

func (pq *personQueue) Process() {
	for {
		if pq.baseQueue.IsEmpty() {
			continue
		}
		pq.RemovePerson()
	}
}

func (pq *personQueue) Size() int {
	return pq.baseQueue.Size()
}
