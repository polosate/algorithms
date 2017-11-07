package matrix

type list struct {
	first    *link
	nextList *list
}

func NewList() *list {
	return &list{
		first:    nil,
		nextList: nil,
	}
}

func (l *list) GetNextList() *list {
	return l.nextList
}

func (l *list) SetNextList(newList *list) {
	l.nextList = newList
}

func (l *list) GetFirst() *link {
	return l.first
}

func (l *list) SetFirst(newLink *link) {
	l.first = newLink
}

func (l *list) GetIterator() iterator {
	return NewIterator(l)
}

func (l *list) IsEmpty() bool {
	return l.first == nil
}
