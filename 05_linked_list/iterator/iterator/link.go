package iterator

import "fmt"

type link struct {
	dData float32
	next  *link
}

func NewLink(dData float32) *link {
	return &link{
		dData: dData,
	}
}

func (l *link) DisplayLink() {
	fmt.Print(l.dData, "; ")
}

func (l *link) GetData() float32 {
	return l.dData
}
