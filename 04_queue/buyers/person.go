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
	return Person{
		num:   num,
		delay: time.Second * time.Duration(rand.Intn(delayRange)),
	}
}

func (p *Person) GetNum() int {
	return p.num
}

func (p *Person) GetDelay() time.Duration {
	return p.delay
}
