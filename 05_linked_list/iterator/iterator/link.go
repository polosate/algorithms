package iterator

import "fmt"

type Link struct {
	dData float32
	next  *Link
}

func NewLink(dData float32) *Link {
	return &Link{
		dData: dData,
	}
}

func (l *Link) DisplayLink() {
	fmt.Print(l.dData, "; ")
}

func (l *Link) GetData() float32 {
	return l.dData
}
