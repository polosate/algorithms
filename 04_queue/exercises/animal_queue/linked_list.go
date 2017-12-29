package exercises

type ILinkedList interface {
	addLast(a IAnimal, order int)
	removeFirst() IAnimal
	peekFirstOrder() int
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
