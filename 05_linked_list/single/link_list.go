package single

import "fmt"

type ISingleLinkList interface {
	Delete(key int) *Link
	DeleteFirst() *Link
	DisplayList()
	Find(key int) *Link
	InsertAfter(key, iData int, dData float32)
	InsertFirst(iData int, dData float32)
	IsEmpty() bool
}

type LinkList struct {
	first *Link
}

func NewLinkList() ISingleLinkList {
	return &LinkList{
		first: nil,
	}
}

func (ll LinkList) IsEmpty() bool {
	return ll.first == nil
}

func (ll *LinkList) InsertFirst(iData int, dData float32) {
	l := NewLink(iData, dData)
	l.next = ll.first
	ll.first = l
}

func (ll *LinkList) DeleteFirst() *Link {
	temp := ll.first
	ll.first = ll.first.next
	return temp
}

func (ll *LinkList) DisplayList() {
	current := ll.first
	for current != nil {
		current.DisplayLink()
		current = current.next
	}
	fmt.Println("--------")
}

func (ll *LinkList) Find(key int) *Link {
	current := ll.first
	for current != nil {
		if current.iData == key {
			return current
		}
		current = current.next
	}
	return current
}

func (ll *LinkList) Delete(key int) *Link {
	previous := ll.first
	current := ll.first

	for current != nil {
		if current.iData == key {
			if current == ll.first {
				ll.first = ll.first.next
			} else {
				previous.next = current.next
			}
			return current
		}
		previous = current
		current = current.next
	}
	return nil
}

func (ll *LinkList) InsertAfter(key, iData int, dData float32) {
}
