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

func (s *Shop) IsEmpty() bool {
	for _, q := range s.queues {
		if !q.IsEmpty() {
			return false
		}
	}
	return true
}

func (s *Shop) Balancer(person Person) {
	sizes := []int{}
	for _, q := range s.queues {
		sizes = append(sizes, q.Size())
		fmt.Println(sizes)
	}
	minI := 0
	minV := sizes[0]
	for i, v := range sizes {
		if v < minV {
			minI = i
		}
	}
	s.queues[minI].AddPerson(person)
}
