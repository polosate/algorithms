package personQueue

import "fmt"

type Shop struct {
	queues []personQueue
}

func NewShop(queuesCount int) Shop {
	queues := []personQueue{}
	for i := 0; i < queuesCount; i++ {
		queues = append(queues, NewQueue(10, i))
	}
	return Shop{
		queues: queues,
	}
}

func (s *Shop) tick() {
	for i := range s.queues {
		p, _ := s.queues[i].PeekPerson()
		if p.DecrementDelay() == 0 {
			s.queues[i].RemovePerson()
		}
		fmt.Println("queue#", s.queues[i].queueNum, s.queues[i].Nums())
	}
	fmt.Println("----------------------------------------")
}

func (s *Shop) Done() bool {
	for _, q := range s.queues {
		if !q.IsEmpty() {
			return false
		}
	}
	return true
}

func (s *Shop) balancer(person Person) {
	sizes := []int{}
	for _, q := range s.queues {
		sizes = append(sizes, q.Size())
	}
	minI := 0
	minV := sizes[0]
	for i := 1; i < len(sizes); i++ {
		if sizes[i] < minV {
			minI = i
			minV = sizes[i]
		}
	}
	s.queues[minI].AddPerson(&person)
}
