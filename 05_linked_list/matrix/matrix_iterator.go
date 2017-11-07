package matrix

type matrixIterator struct {
	matrix *matrix

	current  *list
	previous *list
}

func NewMIterator(m *matrix) matrixIterator {
	matrix := matrixIterator{matrix: m}
	return matrix
}

func (mi *matrixIterator) Reset() {
	mi.current = mi.matrix.firstList
	mi.previous = nil
}

func (mi *matrixIterator) InsertAfter(l *list) {

}
