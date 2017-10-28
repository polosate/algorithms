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
