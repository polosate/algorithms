package single

type LinkList struct {
	first *Link
}

func NewLinkList() LinkList {
	return LinkList{
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
}
