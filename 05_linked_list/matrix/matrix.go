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
	newMatrix.setFirstList()
	currentList := newMatrix.firstList
	for i:=0; i <x-2; i++ {
		list := NewList()
		iter := list.GetIterator()
		for j := 0; j < y; j++ {
			iter.InsertAfter(0)
		}
		currentList.nextList = list
		currentList = list
	}
	return &newMatrix
}

func (m *matrix) IsEmpty() bool {
	return m.firstList == nil
}

func (m *matrix) GetFirstList() *list {
	return m.firstList
}

func (m *matrix) setFirstList() {
	firstList := NewList()
	iter := firstList.GetIterator()
	for j := 0; j < m.y; j++ {
		iter.InsertAfter(0)
	}
}

func (m *matrix) DisplayMatrix() {

}
