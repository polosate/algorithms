package matrix

type matrix struct {
	firstList *list
	x         int
	y         int
}

func NewMatrix(x, y int) *matrix {
	newMatrix := matrix{
		x: x,
		y: y,
	}
	// create first list
	firstList := newMatrix.emptyList()
	newMatrix.firstList = firstList

	// create other lists
	currentList := newMatrix.firstList
	for i := 0; i < x-2; i++ {
		list := newMatrix.emptyList()
		currentList.nextList = list
		currentList = list
	}
	return &newMatrix
}

func (m *matrix) emptyList() *list {
	firstList := NewList()
	iter := firstList.GetIterator()
	for j := 0; j < m.y; j++ {
		iter.InsertAfter(0)
	}
	return firstList
}

func (m *matrix) SetValue(x, y int, value float32) {
	currentList := m.firstList
	for i := 0; i < x; i++ {
		currentList = currentList.nextList
	}

	currentItem := currentList.first
	for i := 0; i < y; i++ {
		currentItem = currentItem.next
	}

	currentItem.SetValue(5.5)
}

func (m *matrix) DisplayMatrix() {
	currentList := m.firstList
	for currentList != nil {
		currentList.DisplayList()
		currentList = currentList.nextList
	}
}
