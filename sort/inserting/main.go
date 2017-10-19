package inserting

type array struct {
	data []int64
	len  int
}

func newArray(size int) array {
	return array{make([]int64, size), 0}
}

func (this *array) Insert(elem int64) bool {
	if this.len >= cap(this.data) {
		return false
	}
	this.data[this.len] = elem
	this.len++
	return true
}

// 4 6 2 1 8 4 2

func (this *array) InsertingSort() {

}

func (this *array) GetData() []int64 {
	return this.data[:this.len]
}

func (this *array) GetLen() int {
	return this.len
}
