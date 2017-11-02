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

func (this *iterator) Reset() {
	this.current = this.list.GetFirst()
	this.previous = nil
}

func (this *iterator) AtEnd() bool {
	return this.current.next == nil
}

func (this *iterator) NextLink() {
	this.previous = this.current
	this.current = this.current.next
}

func (this *iterator) GetCurrent() *link {
	return this.current
}

func (this *iterator) InsertAfter(dData float32) {
	link := NewLink(dData)
	if this.list.IsEmpty() {
		this.list.SetFirst(link)
		this.current = link
	} else {
		link.next = this.current.next
		this.current.next = link
		this.NextLink()
	}
}

func (this *iterator) InsertBefore(dData float32) {
	link := NewLink(dData)
	if this.previous == nil {
		link.next = this.list.GetFirst()
		this.list.SetFirst(link)
		this.Reset()
	} else {
		this.previous.next = link
		link.next = this.current
		this.current = link
	}
}

func (this *iterator) DeleteCurrent() *link {
	temp := this.current
	if this.previous == nil {
		this.list.SetFirst(this.current.next)
		this.Reset()
	} else {
		this.previous.next = this.current.next
		if this.AtEnd() {
			this.Reset()
		} else {
			this.current = this.current.next
		}
	}
	return temp
}
