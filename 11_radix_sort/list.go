package _1_radix_sort

import "fmt"

type link struct {
	value int64
	next  *link
}

func NewLink(value int64) *link {
	return &link{value: value}
}

func (l *link) GetValue() int64 {
	return l.value
}

func (l *link) DisplayLink() {
	fmt.Print(l.value, "; ")
}

type list struct {
	first *link
	last  *link
}

func NewList() list {
	return list{
		first: nil,
		last:  nil,
	}
}

func (ll *list) isEmpty() bool {
	return ll.first == nil
}

func (ll *list) InsertLast(value int64) {
	newLink := NewLink(value)
	if ll.isEmpty() {
		ll.first = newLink
	} else {
		ll.last.next = newLink
	}
	ll.last = newLink
}

func (ll *list) RemoveFirst() int64 {
	temp := ll.first
	if ll.first == ll.last {
		ll.last = nil
	}
	ll.first = temp.next
	return temp.value
}
