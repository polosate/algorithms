package buyers

type Shop struct {
	queues []personQueue
}

func NewShop(count int) Shop {
	queues := make([]personQueue, count)
	for i := 0; i < count; i++ {
		queues = append(queues, NewQueue(queueSize, i))
	}
	return Shop{
		queues: queues,
	}
}

func (s *Shop) IsEmpty() bool {
	res := 0
	for _, q := range s.queues {
		if q.Size() == 0 {
			res++
		}
	}
	return res == len(s.queues)
}

func (s *Shop) Balancer(person Person) {
	sizes := make([]int, len(s.queues))
	for _, q := range s.queues {
		sizes = append(sizes, q.Size())
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
