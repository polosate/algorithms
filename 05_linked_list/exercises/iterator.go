package exercises

type iterator struct {
	current  *link
	previous *link
	list     *singleList
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
	i.previous = nil
}

func (i *iterator) GetCurrent() *link {
	return i.current
}

// ->
// -> 1->
// ->1->2->
func (i *iterator) NextLink() {
	i.previous = i.current
	i.current = i.current.next
}
