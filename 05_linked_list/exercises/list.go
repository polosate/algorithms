package exercises

import "fmt"

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
	fmt.Println()
}

func (sl *singleList) InsertFirst(value int64) {
	newLink := NewLink(value)
	newLink.next = sl.first
	sl.first = newLink
}

func (sl *singleList) GetIterator() iterator {
	return NewIterator(sl)
}

func (sl *singleList) Reverse() *singleList {
	if sl.IsEmpty() || sl.first.next == nil {
		return sl
	}
	var prev, cur, next *link
	cur = sl.first
	next = cur.next
	for cur != nil {
		cur.next = prev
		if next == nil {
			sl.first = cur
			return sl
		}
		prev = cur
		cur = next
		next = next.next
	}
	return sl
}
