package buyers

import (
	"math/rand"
	"time"
)

type Person struct {
	num   int
	delay time.Duration
}

func NewPerson(num int, delayRange int) Person {
	d := time.Duration(time.Millisecond * time.Duration(rand.Intn(delayRange)))

	return Person{
		num:   num,
		delay: d / 1000000,
	}
}

func (p *Person) GetNum() int {
	return p.num
}

func (p *Person) GetDelay() time.Duration {
	return p.delay
}
