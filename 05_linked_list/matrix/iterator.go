package matrix

type iterator struct {
	myList *list

	current  *link
	previous *link
}

func NewIterator(l *list) iterator {
	iterator := iterator{
		myList: l,
	}
	iterator.Reset()
	return iterator
}

func (i *iterator) Reset() {
	i.current = i.myList.GetFirst()
	i.previous = nil
}

func (i *iterator) AtEnd() bool {
	if i.myList.IsEmpty() {
		return true
	} else {
		return i.current.next == nil
	}
}

func (i *iterator) GetCurrent() *link {
	return i.current
}

func (i *iterator) NextLink() {
	if i.myList.IsEmpty() {
		return
	} else if i.AtEnd() {
		i.Reset()
	} else {
		i.previous = i.current
		i.current = i.current.next
	}
}

func (i *iterator) InsertAfter(value float32) {
	newLink := NewLink(value)
	if i.myList.IsEmpty() {
		i.myList.SetFirst(newLink)
		i.Reset()
	} else {
		newLink.next = i.current.next
		i.current.next = newLink
		i.NextLink()
	}
}
