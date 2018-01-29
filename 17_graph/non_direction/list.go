package non_direction

type link struct {
	vertex *vertex
	index  int
	next   *link
}

func newLink(vertex *vertex, i int) *link {
	return &link{vertex: vertex, index: i}
}

type list struct {
	first *link
}

func newList() list {
	return list{}
}

func (l *list) add(v *vertex, i int) {
	newLink := newLink(v, i)
	if l.first == nil {
		l.first = newLink
	} else {
		temp := l.first
		newLink.next = temp
		l.first = newLink
	}
}

func (l *list) find() int {
	current := l.first
	for current != nil {
		if !current.vertex.wasVisited {
			return current.index
		}
		current = current.next
	}
	return -1
}
