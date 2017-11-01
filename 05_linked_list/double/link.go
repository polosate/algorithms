package double

import "fmt"

type ILink interface {
	DisplayLink()
	GetData() float32
}

type link struct {
	dData    float32
	previous *link
	next     *link
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
