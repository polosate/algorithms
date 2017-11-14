package exercises

import "fmt"

type link struct {
	value int64
	next  *link
}

func NewLink(value int64) *link {
	return &link{
		value: value,
	}
}

func (l link) DisplayLink() {
	fmt.Print(l.value, "; ")
}
