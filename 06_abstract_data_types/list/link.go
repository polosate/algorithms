package list

import "fmt"

type ILink interface {
	DisplayLink()
	GetData() float32
	SetNext(*Link)
	GetNext() *Link
}

type ISingleLink interface {
	ILink
}

type IDoubleLink interface {
	ILink
	SetPrevious(*Link)
	GetPrevious() *Link
}

type Link struct {
	dData float32
	next  *Link
}

type SingleLink struct {
	Link
}

type DoubleLink struct {
	Link
	previuous *Link
}

func newLink(dData float32) *Link {
	return &Link{
		dData: dData,
	}
}

func newSingleLink(dData float32) ISingleLink {
	return &SingleLink{
		*newLink(dData),
	}
}

func newDoubleLink(dData float32) IDoubleLink {
	return &DoubleLink{
		Link:*newLink(dData),
	}
}

func (l *Link) DisplayLink() {
	fmt.Print(l.dData, "; ")
}

func (l *Link) GetData() float32 {
	return l.dData
}

func (l *Link) SetNext(newLink *Link) {

}

func (l *Link) GetNext() *Link {
	return nil
}

func (l *DoubleLink) SetPrevious(newLink *Link) {

}

func (l *DoubleLink) GetPrevious() *Link {
	return nil
}
