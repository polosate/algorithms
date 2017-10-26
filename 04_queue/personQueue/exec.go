package personQueue

import "fmt"

func main() {
	var queues []personQueue
	for i := 1; i <= 3; i++ {
		queues = append(queues, NewQueue(10, i))
	}

	for i := 0; i < 30; i++ {
		balancer(queues, NewPerson(i, 10))
	}

	for !Done(queues) {
		for i := range queues {
			p, _ := queues[i].PeekPerson()
			if p.delay-1 == 0 {
				queues[i].RemovePerson()
			}
			fmt.Println(queues[i].baseQueue.Display())
		}
	}
}

func Done(queues []personQueue) bool {
	for _, q := range queues {
		if !q.IsEmpty() {
			return false
		}
	}
	return true
}

func balancer(queues []personQueue, person Person) {
	sizes := []int{}
	for _, q := range queues {
		sizes = append(sizes, q.Size())
	}
	minI := 0
	minV := sizes[0]
	for i := 1; i < len(sizes); i++ {
		if sizes[i] < minV {
			minI = i
		}
	}
	queues[minI].AddPerson(person)
}
