package buyers

import (
	"sync"
)

const (
	queueSize  = 100
	queueCount = 3
)

func Process() {
	var wg sync.WaitGroup

	shop := NewShop(queueCount)

	for _, q := range shop.queues {
		wg.Add(1)
		go q.Process()
	}

	for i := 0; i < 100; i++ {
		shop.Balancer(NewPerson(i, 10))
	}


	wg.Wait()
}

func initQueue(count int) []personQueue {
	var queues []personQueue
	for i := 0; i < count; i++ {
		queues = append(queues, NewQueue(queueSize, i))
	}
	return queues
}
