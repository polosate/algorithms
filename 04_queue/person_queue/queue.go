package person_queue

import (
	baseQ "algorithms/04_queue"
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

func (pq *personQueue) AddPerson(person *Person) error {
	return pq.baseQueue.Insert(person)
}

func (pq *personQueue) PeekPerson() (*Person, error) {
	element, err := pq.baseQueue.Peek()
	if err != nil {
		return &Person{}, err
	}
	person, _ := element.(*Person)
	return person, nil
}

func (pq *personQueue) RemovePerson() (*Person, error) {
	element, err := pq.baseQueue.Remove()
	if err != nil {
		return &Person{}, err
	}
	person, _ := element.(Person)
	return &person, nil
}

func (pq *personQueue) IsEmpty() bool {
	return pq.baseQueue.IsEmpty()
}

func (pq *personQueue) Size() int {
	return pq.baseQueue.Size()
}

func (pq *personQueue) PersonsNums() []int {
	var result []int
	response := pq.baseQueue.Display1()
	for _, v := range response {
		p, _ := v.(*Person)
		result = append(result, p.GetNum())
	}
	return result

}
