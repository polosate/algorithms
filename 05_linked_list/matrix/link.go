package matrix

import "fmt"

type link struct {
	value float32
	next  *link
}

func NewLink(value float32) *link {
	return &link{
		value: value,
	}
}

func (l *link) GetValue() float32 {
	return l.value
}

func (l *link) SetValue(value float32) {
	l.value = value
}

func (l *link) GetNext() *link {
	return l.next
}

func (l *link) SetNext(newLInk *link) {
	l.next = newLInk
}

func (l *link) DisplayLink() {
	fmt.Print(l.value, "; ")
}
