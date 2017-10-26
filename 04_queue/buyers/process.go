package buyers

import (
	"time"
)

const (
	queueSize  = 100
	queueCount = 3
)

func Exec() {

	shop := NewShop(queueCount)

	for _, q := range shop.queues {
		go q.Process()
	}

	time.Sleep(time.Second * 1)

	for i := 0; i < 10; i++ {
		shop.Balancer(NewPerson(i, 10))
	}

	for !shop.IsEmpty() {
		continue
	}
}
