package exercises

type iterator struct {
	current *link
	prev    *link
	list    *singleList
}

func NewIterator(list *singleList) iterator {
	i := iterator{
		list: list,
	}
	i.Reset()
	return i
}

func (i *iterator) Reset() {
	i.current = i.list.first
	i.prev = nil
}
