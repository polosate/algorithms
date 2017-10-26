package personQueue

import (
	"math/rand"
)

type Person struct {
	num   int
	delay int
}

func NewPerson(num int, delayRange int) Person {
	d := rand.Intn(delayRange)
	return Person{
		num:   num,
		delay: d,
	}
}

func (p *Person) GetNum() int {
	return p.num
}

func (p *Person) GetDelay() int {
	return p.delay
}
