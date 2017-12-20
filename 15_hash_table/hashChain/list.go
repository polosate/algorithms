package hashChain

import (
	. "algorithms/15_hash_table"
	"fmt"
)

type link struct {
	item *DataItem
	next *link
}

func newLink(item *DataItem) *link {
	return &link{
		item: item,
	}
}

type list struct {
	first *link
}

func newList() list {
	return list{}
}

func (l *list) isEmpty() bool {
	return l.first == nil
}

func (l *list) insertFirst(item *DataItem) {
	newLink := newLink(item)
	newLink.next = l.first
	l.first = newLink
}

func (l *list) find(key int64) *DataItem {
	curLink := l.first
	for curLink != nil {
		if curLink.item.GetKey() == key {
			return curLink.item
		}
		curLink = curLink.next
	}
	return nil
}

func (l *list) delete(key int64) *DataItem {
	var prevLink *link
	curLink := l.first
	for curLink != nil && curLink.item.GetKey() != key {
		prevLink = curLink
		curLink = curLink.next
	}
	if curLink == nil {
		return nil
	}
	temp := curLink.item
	if curLink == l.first {
		l.first = curLink.next
	} else {
		prevLink.next = curLink.next
	}
	curLink = nil
	return temp
}

func (l *list) display() {
	cur := l.first
	for cur != nil {
		fmt.Print(cur.item.GetKey(), " ")
		cur = cur.next
	}
	fmt.Println()
}
