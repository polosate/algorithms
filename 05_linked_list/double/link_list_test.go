package double

import "testing"

func TestInsertFirst(t *testing.T) {
	ll := NewDoubleLinkedList()
	ll.InsertFirst(1)
	ll.InsertFirst(2)
	ll.DisplayForward()
	ll.DisplayBackward()

}

func TestInsertLast(t *testing.T) {
	ll := NewDoubleLinkedList()
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.DisplayForward()
	ll.DisplayBackward()
}

func TestInsertAfter(t *testing.T) {
	ll := NewDoubleLinkedList()
	ll.InsertLast(3.3)
	ll.InsertLast(1.1)
	ll.InsertLast(4.4)
	ll.InsertAfter(3.3, 2.2)
	ll.DisplayForward()
}
