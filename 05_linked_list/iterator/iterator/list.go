package iterator

import "fmt"

type IList interface {
	GetFirst() *link
	SetFirst(*link)
	IsEmpty() bool
	GetIterator() IIterator
	DisplayList()
}

type linkList struct {
	first *link
}

func NewLinkList() IList {
	return &linkList{
		first: nil,
	}
}

func (ll *linkList) GetFirst() *link {
	return ll.first
}

func (ll *linkList) SetFirst(l *link) {
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
