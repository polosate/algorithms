package single

import "fmt"

type FirstLastList struct {
	first *Link
	last  *Link
}

func NewFirstLastList() FirstLastList {
	return FirstLastList{
		first: nil,
		last:  nil,
	}
}

func (ll *FirstLastList) IsEmpty() bool {
	return ll.first == nil
}

func (ll *FirstLastList) InsertFirst(iData int, dData float32) {
	l := NewLink(iData, dData)
	if ll.IsEmpty() {
		ll.last = l
	}
	l.next = ll.first
	ll.first = l
}

func (ll *FirstLastList) InsertLast(iData int, dData float32) {
	l := NewLink(iData, dData)
	if ll.IsEmpty() {
		ll.first = l
	} else {
		ll.last.next = l
	}
	ll.last = l
}

func (ll *FirstLastList) DeleteFirst() *Link {
	if ll.IsEmpty() {
		return nil
	}
	temp := ll.first
	ll.first = temp.next
	return temp
}

func (ll *FirstLastList) DisplayList() {
	current := ll.first
	for current != nil {
		current.DisplayLink()
		current = current.next
	}
	fmt.Println("\n--------")
}
