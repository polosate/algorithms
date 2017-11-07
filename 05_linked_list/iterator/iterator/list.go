package iterator

import "fmt"

type IList interface {
	GetFirst() *Link
	SetFirst(*Link)
	IsEmpty() bool
	GetIterator() IIterator
	DisplayList()
}

type linkList struct {
	first *Link
}

func NewLinkList() IList {
	return &linkList{
		first: nil,
	}
}

func (ll *linkList) GetFirst() *Link {
	return ll.first
}

func (ll *linkList) SetFirst(l *Link) {
	ll.first = l
}

func (ll *linkList) IsEmpty() bool {
	return ll.first == nil
}

func (ll *linkList) GetIterator() IIterator {
	return NewIterator(ll)
}

func (ll *linkList) DisplayList() {
	fmt.Println("List (first --> last)")
	current := ll.first
	for current != nil {
		current.DisplayLink()
		current = current.next
	}
	fmt.Println("")
}
