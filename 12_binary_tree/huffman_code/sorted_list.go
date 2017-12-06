package huffman_code

import "fmt"

type link struct {
	value *Node
	next  *link
}

func NewLink(node *Node) *link {
	return &link{value: node}
}

func (l *link) PrintLink() {
	fmt.Print(l.value.char, ": ", l.value.count, "; ")
}

type sortedList struct {
	first *link
}

func NewList() sortedList {
	return sortedList{first: nil}
}

func (l *sortedList) IsEmpty() bool {
	return l.first == nil
}

func (l *sortedList) Insert(node *Node) {
	newLink := NewLink(node)
	if l.IsEmpty() || newLink.value.count <= l.first.value.count {
		newLink.next = l.first
		l.first = newLink
		return
	}
	var prev *link
	current := l.first
	for current != nil && current.value.count < newLink.value.count {
		prev = current
		current = current.next
	}
	prev.next = newLink
	if current != nil {
		newLink.next = current
	}
}

func (l *sortedList) Remove() *Node {
	if l.IsEmpty() {
		return nil
	}
	removed := l.first
	l.first = removed.next
	return removed.value
}

func (l *sortedList) PrintList() {
	current := l.first
	for current != nil {
		current.PrintLink()
		current = current.next
	}
	fmt.Println()
}
