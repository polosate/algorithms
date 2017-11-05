package ring

import "fmt"

type ILink interface {
	GetValue() float32
	GetNext() *link
	DisplayLink()
}

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

func (l *link) GetNext() *link {
	return l.next
}

func (l *link) DisplayLink() {
	fmt.Print(l.value, "; ")
}
