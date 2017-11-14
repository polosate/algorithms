package exercises

type singleList struct {
	first *link
}

func NewSingleList() *singleList {
	return &singleList{
		first: nil,
	}
}

func (sl *singleList) IsEmpty() bool {
	return sl.first == nil
}

func (sl *singleList) DisplayList() {
	current := sl.first
	for current != nil {
		current.DisplayLink()
		current = current.next
	}
}

func (sl *singleList) InsertFirst(value int64) {
	newLink := NewLink(value)
	newLink.next = sl.first
	sl.first = newLink
}
