package list

import "fmt"

type ILinkList interface {
	IsEmpty() bool
	InsertFirst(dData float32)
	InsertLast(dData float32)
	DeleteFirst() *Link
	PeekFirst() *Link
	DisplayList()
}

type linkList struct {
	first *Link
	last  *Link
}

func NewLinkList() ILinkList {
	return &linkList{
		first: nil,
		last:  nil,
	}
}

func (ll *linkList) IsEmpty() bool {
	return ll.first == nil
}

func (ll *linkList) InsertFirst(dData float32) {
	l := newLink(dData)
	if ll.IsEmpty() {
		ll.last = l
	}
	l.next = ll.first
	ll.first = l
}

func (ll *linkList) InsertLast(dData float32) {
	l := newLink(dData)
	if ll.IsEmpty() {
		ll.first = l
	} else {
		ll.last.next = l
	}
	ll.last = l
}

func (ll *linkList) DeleteFirst() *Link {
	if ll.IsEmpty() {
		return nil
	}
	temp := ll.first
	ll.first = temp.next
	return temp
}

func (ll *linkList) PeekFirst() *Link {
	return ll.first
}

func (ll *linkList) DisplayList() {
	current := ll.first
	for current != nil {
		current.DisplayLink()
		current = current.next
	}
	fmt.Println("\n-------")
}
