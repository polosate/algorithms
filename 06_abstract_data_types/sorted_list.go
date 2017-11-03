package _6_abstract_data_types

import "fmt"

type ISortedList interface {
	IsEmpty() bool
	Insert(dData float32)
	Remove() *Link
	Peek() *Link
	Size() int
	DisplayList()
}

type sortedList struct {
	first *Link
	size  int
}

func NewSortedList() ISortedList {
	return &sortedList{
		first: nil,
		size:  0,
	}
}

func (sl *sortedList) IsEmpty() bool {
	return sl.first == nil
}

func (sl *sortedList) Insert(dData float32) {
	newLink := newLink(dData)
	if sl.IsEmpty() {
		newLink.next = sl.first
		sl.first = newLink
	} else {
		current := sl.first
		var previuos *Link

		for current != nil && current.dData < dData {
			previuos = current
			current = current.next
		}
		if current == nil {
			newLink.next = previuos.next
			previuos.next = newLink
		} else if current == sl.first {
			newLink.next = current
			sl.first = newLink
		} else {
			newLink.next = current
			previuos.next = newLink
		}
	}
	sl.size++
}

func (sl *sortedList) Remove() *Link {
	if sl.IsEmpty() {
		return nil
	}
	temp := sl.first
	sl.first = temp.next
	sl.size--
	return temp
}

func (sl *sortedList) Peek() *Link {
	if sl.IsEmpty() {
		return nil
	}
	return sl.first
}

func (sl *sortedList) Size() int {
	return sl.size
}

func (sl *sortedList) DisplayList() {
	if sl.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	current := sl.first
	for current != nil {
		current.DisplayLink()
		current = current.next
	}
	fmt.Println("\n-------------")
}
