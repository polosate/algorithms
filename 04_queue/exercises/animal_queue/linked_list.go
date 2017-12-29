package animal_queue

// TODO: implement linked list

type ILinkedList interface {
	addLast(a IAnimal, order int)
	removeFirst() IAnimal
	peekFirst() ILink
}

type ILink interface {
	getAnimal() IAnimal
	getOrder() int
	isOlderThan(l ILink) bool

	getNext() ILink
}

type link struct {
	ILink
}

func (l *link) isOlderThan(otherLink ILink) bool {
	return l.getOrder() < otherLink.getOrder()
}
