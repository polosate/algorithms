package single

import (
	"errors"
	"fmt"
)

type sortedList struct {
	first *Link
}

func NewSortedList() sortedList {
	return sortedList{
		first: nil,
	}
}

func (sl *sortedList) IsEmpty() bool {
	return sl.first == nil
}

func (sl *sortedList) Insert(iData int, dData float32) {
	l := NewLink(iData, dData)

	if sl.IsEmpty() {
		l.next = sl.first
		sl.first = l
		return
	}
	current := sl.first
	var previous *Link

	for current != nil && current.iData < l.iData {
		previous = current
		current = current.next
	}
	if previous == nil {
		sl.first = l
	} else {
		previous.next = l
	}
	l.next = current
}

func (sl *sortedList) Remove() (*Link, error) {
	if sl.IsEmpty() {
		return nil, errors.New("base is empty")
	}
	temp := sl.first
	sl.first = temp.next
	return temp, nil
}

func (sl *sortedList) DisplayList() {
	current := sl.first
	for current != nil {
		current.DisplayLink()
		current = current.next
	}
	fmt.Println("Sorted base (first --> last)")
}
