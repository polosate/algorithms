package iterator

type IIterator interface {
	Reset()
	AtEnd() bool
	NextLink()
	GetCurrent() *link
	InsertAfter(dData float32)
	InsertBefore(dData float32)
	DeleteCurrent() *link
}

type iterator struct {
	list     IList
	current  *link
	previous *link
}

func NewIterator(list IList) IIterator {
	iterator := iterator{
		list: list,
	}
	iterator.Reset()
	return &iterator
}

func (i *iterator) Reset() {
	i.current = i.list.GetFirst()
	i.previous = nil
}

func (i *iterator) AtEnd() bool {
	return i.current.next == nil
}

func (i *iterator) NextLink() {
	i.previous = i.current
	i.current = i.current.next
}

func (i *iterator) GetCurrent() *link {
	return i.current
}

func (i *iterator) InsertAfter(dData float32) {
	link := NewLink(dData)
	if i.list.IsEmpty() {
		i.list.SetFirst(link)
		i.current = link
	} else {
		link.next = i.current.next
		i.current.next = link
		i.NextLink()
	}
}

func (i *iterator) InsertBefore(dData float32) {
	link := NewLink(dData)
	if i.previous == nil {
		link.next = i.list.GetFirst()
		i.list.SetFirst(link)
		i.Reset()
	} else {
		i.previous.next = link
		link.next = i.current
		i.current = link
	}
}

func (i *iterator) DeleteCurrent() *link {
	temp := i.current
	if i.previous == nil {
		i.list.SetFirst(i.current.next)
		i.Reset()
	} else {
		i.previous.next = i.current.next
		if i.AtEnd() {
			i.Reset()
		} else {
			i.current = i.current.next
		}
	}
	return temp
}
